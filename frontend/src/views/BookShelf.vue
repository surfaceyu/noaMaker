<script lang="ts" setup>
import { computed, onBeforeUnmount, ref, shallowRef } from 'vue';
import { type Ref } from 'vue';
import { models } from '@wailsjs/go/models';
import { Search } from '@element-plus/icons-vue'
import { ElMessage } from 'element-plus'
import { GetSites } from '@wailsjs/go/backend/BookSitesHandler';
import { SearchBook, SearchChapter, SearchContent } from '@wailsjs/go/backend/BookHandler'
import RichEditor from '@components/RichEditor.vue';
// import { Editor, Toolbar } from '@wangeditor/editor-for-vue'

const input3 = ref('')
const select = ref(0)
const selectBook = ref<models.Book>()
const selectChapter = ref<models.Chapter>()
const tableData: Ref<Array<models.BookSite>> = ref([]);
const books = ref<Array<models.Book>>([]);
const chapters = ref<Array<models.Chapter>>([]);
const content = ref<models.Content>();
const editContent = ref('')

const initSitesData = async () => {
    tableData.value = await GetSites()
}

const bookSites = computed(() => {
    return tableData.value
})

initSitesData()

/**
 * 搜索书籍
 */
async function handleSearch(name: string, index: number) {
    ElMessage.info('书本搜索中......')
    // 清空现有的
    books.value = []
    books.value = await SearchBook(name, index) || []
    ElMessage.success('书本搜索完成！')
}

async function handleSelectBook(book: models.Book) {
    ElMessage.info('章节搜索中......')
    chapters.value = []
    chapters.value = await SearchChapter(book, bookSites.value[select.value])
    ElMessage.success('章节搜索完成！')
}

async function handleSelectChapter(chapter: models.Chapter) {
    ElMessage.info('正文搜索中......')
    content.value = await SearchContent(chapter, bookSites.value[select.value])
    editContent.value = `${chapter.chapterName}${content.value.content}`
    ElMessage.success('正文搜索完成！')
}
</script>

<template>
    <h5 class="main-content-title">书架管理</h5>
    <div class="mbt-20">
        <el-input v-model="input3" @keyup.enter="handleSearch(input3, select)" placeholder="输入书名进行搜索" class="input-with-select">
            <template #append>
                <el-button :icon="Search" @click="handleSearch(input3, select)" />
            </template>
            <template #prepend>
                <el-select v-model="select" filterable placeholder="选择书源" style="width: 115px">
                    <el-option v-for="(item, index) in bookSites" :label="item.name" :value="index" />
                </el-select>
            </template>
        </el-input>
    </div>
    <el-row :gutter="20" class="mbt-20">
        <el-col :span="12" :offset="0">
            <!-- 选择器选择选项后触发 handleSelectBook 方法-->
            <el-select v-model="selectBook" value-key="bookName" placeholder="选择书本" @change="handleSelectBook">
                <el-option v-for="item in books" :label="item.bookName" :value="item" :key="item.bookName">
                    <span style="float: left">{{ item.bookName }}</span>
                    <span style="
                  float: right;
                  color: var(--el-text-color-secondary);
                  font-size: 13px;
                ">{{ item.author }}</span>
                </el-option>
            </el-select>
        </el-col>
        <el-col :span="12" :offset="0">
            <el-select v-model="selectChapter" filterable value-key="chapterName" placeholder="选择章节"
                @change="handleSelectChapter">
                <el-option v-for="item in chapters" :label="item.chapterName" :value="item" :key="item.chapterName">
                    {{ item.chapterName }}
                </el-option>
            </el-select>
        </el-col>
    </el-row>
    <RichEditor :data="editContent" :name="selectBook?.bookName" :chapter="selectChapter?.chapterName"/>
</template>

<style scoped>
.main-content-title {
    text-align: left;
    padding-left: 20px;
    font-size: large;
    color: #e7ac08;
}

.input-with-select .el-input-group__prepend {
    background-color: var(--el-fill-color-blank);
}

.mbt-20 {
    margin-bottom: 20px;
    margin-top: 20px;
}
</style>
