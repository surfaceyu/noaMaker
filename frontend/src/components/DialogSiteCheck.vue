<script lang="ts" setup>
import { computed, ref } from 'vue';
import { ElMessage } from 'element-plus'
import { models } from '@wailsjs/go/models';
import { SearchBook, SearchChapter, SearchContent } from '@wailsjs/go/backend/BookHandler'
import { EditSites } from '@wailsjs/go/backend/BookSitesHandler'

const props = defineProps({
    data: {
        type: Object as () => models.BookSite,
        required: true,
    },
    index: {
        type: Number,
        required: true,
    },
})

const isShow = ref(false)
const books = ref<Array<models.Book>>([]);
const chapters = ref<Array<models.Chapter>>([]);
const content = ref<models.Content>()

defineExpose({
    isShow
})
/**
 * 搜索书籍
 */
async function handleSearch() {
    // 检测搜索功能 默认搜索关于剑仙的小说
    books.value = await SearchBook("剑仙", props.index)
}
/**
 * 搜索章节列表
 */
async function handleChapter() {
    let book = books.value[0];
    if (!book) {
        await handleSearch();
        book = books.value[0];
        if (!book) {
            ElMessage.error("没有搜索到书籍，请先搜索书本！")
            return;
        }
    }
    // 搜索book的章节列表
    const data = await SearchChapter(book, props.data)
    if (!data) {
        ElMessage.error("没有搜索到章节，请检查搜索规则！")
        return;
    } else {
        chapters.value = data
    }
}
/**
 * 搜索章节内容
 */
async function handleContent() {
    let ch = chapters.value[0];
    if (!ch) {
        await handleChapter();
        ch = chapters.value[0];
        if (!ch) {
            ElMessage.error("没有搜索到书籍章节，请检查章节或书本规则！")
            return;
        }
    }
    const data = await SearchContent(chapters.value[0], props.data)
    if (!data) {
        ElMessage.error("没有搜索到小说内容，请检查搜索规则！")
        return;
    } else {
        content.value = data
    }
}

const randTagSize = () => {
    const sizes = ["small", "default", "large"];
    // 返回随机的size
    return sizes[Math.floor(Math.random() * sizes.length)];
}

const randTagColor = () => {
    const colors = ["", "success", "info", "warning", "danger"];
    // 返回随机的color
    return colors[Math.floor(Math.random() * colors.length)];
}

/**
 * 步骤条对应的处理方法
 */
const stepHandles = [handleSearch, handleChapter, handleContent]

/**
 *  检查步骤
 */
const checkStep = () => {
    // 如果存在对应的处理函数，则执行该函数
    if (stepHandles[props.data.checkStep]) {
        stepHandles[props.data.checkStep]()
    }
}
const checkOk = () => {
    if (props.data.checkStep < 3) {
        props.data.checkStep++;
    }
}
const stepReset = () => {
    props.data.checkStep = 0
    resetTags();
}
/**
 * 清除Tag显示数据
 */
const resetTags = () => {
    books.value = []
}

const stepDisable = computed(() => {
    return props.data.checkStep >= 3
})

const active = computed(() => {
    EditSites(props.data)
    return props.data.checkStep
})

const stepData = computed(() => {
    let result: Array<models.Book | models.Chapter | string> = []
    switch (props.data.checkStep) {
        case 0:
            result = books.value.filter((_, i) => i < 3)
            break;
        case 1:
            result = chapters.value.filter((_, i) => i < 3)
            break;
        case 2:
            // 字符串截取100位 result
            if (content.value) {
                result.push(`${content.value.content.substring(0, 30)}......`);
            }
            break;
        default:
            break;
    }
    return result
})
</script>

<template>
    <el-dialog v-model="isShow" title="检测书源">
        <el-steps :active="active" finish-status="success" simple style="margin-top: 20px">
            <el-step title="书籍搜索" />
            <el-step title="章节列表" />
            <el-step title="章节内容" />
        </el-steps>
        <el-button class="mt-12" type="danger" disabled>操作</el-button>
        <el-button class="mt-12" :disabled="stepDisable" @click="checkStep">检测</el-button>
        <el-button class="mt-12" :disabled="stepDisable" @click="checkOk">检测完成</el-button>
        <el-button class="mt-12" :disabled="stepDisable" @click="stepReset">重置</el-button>
        <el-row class="el-row-books">
            <el-col :span="24">
                <el-tag v-for="item in stepData" effect="dark" :type="randTagColor()" :size="randTagSize()">
                    {{ item }}
                </el-tag>
            </el-col>
        </el-row>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="isShow = false">关闭</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<style scoped>
.el-row {
    margin-bottom: 20px;
}

.el-row:last-child {
    margin-bottom: 0;
}

.el-col {
    border-radius: 4px;
}

.el-tag {
    margin-right: 10px;
    margin-top: 2px;
}

.dialog-footer button:first-child {
    margin-right: 10px;
}

.mt-12 {
    margin-top: 12px;
}
</style>
