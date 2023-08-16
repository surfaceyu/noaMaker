  
<script setup lang="ts">
import _ from 'lodash'
import { onUpdated, ref } from 'vue';
import { models } from '@wailsjs/go/models';
import { Search } from '@element-plus/icons-vue'
import { SearchBookByRule } from '@backend/BookHandler'

const form = defineProps({
    data: {
        type: Object as () => models.BookSiteUri,
        required: true
    },
})

const book = ref<models.Book>()

/**
 * 搜索书籍
 */
async function handleSearch() {
    // 检测搜索功能 默认搜索关于剑仙的小说
    const books = await SearchBookByRule("剑仙", form.data)
    if (books && books.length > 0) {
        book.value = books[0]
    } else {
        book.value = undefined
    }
}

function getProperty(obj: any, key: string) {
    return obj[key];
}

function attrIsSuccess(attr: string): boolean {
    return book.value ? getProperty(book.value, attr) : false
}

onUpdated(_.debounce(() => {
    handleSearch()
}, 300))
</script>

<template>
    <el-text> <el-icon>
            <Collection />
        </el-icon>书本搜索规则</el-text>
    <el-form-item prop="searchUrl.uri" label="搜索规则">
        <el-input v-model="form.data.uri">
            <template #append>
                <el-button :icon="Search" @click="handleSearch" />
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="searchUrl.bookList" label="列表规则">
        <el-input v-model="form.data.bookList">
            <template #append>
                <el-icon class="el-icon-success" v-if="attrIsSuccess('bookName')">
                    <SuccessFilled />
                </el-icon>
                <el-icon class="el-icon-failed" v-else>
                    <CircleCloseFilled />
                </el-icon>
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="searchUrl.author" label="作者规则">
        <el-input v-model="form.data.author">
            <template #append>
                <el-tooltip effect="dark" :content="book?.author" placement="top-start">
                    <el-icon class="el-icon-success" v-if="attrIsSuccess('author')">
                        <SuccessFilled />
                    </el-icon>
                    <el-icon class="el-icon-failed" v-else>
                        <CircleCloseFilled />
                    </el-icon>
                </el-tooltip>
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="searchUrl.bookName" label="书名规则">
        <el-input v-model="form.data.bookName">
            <template #append>
                <el-tooltip effect="dark" :content="book?.bookName" placement="top-start">
                    <el-icon class="el-icon-success" v-if="attrIsSuccess('bookName')">
                        <SuccessFilled />
                    </el-icon>
                    <el-icon class="el-icon-failed" v-else>
                        <CircleCloseFilled />
                    </el-icon>
                </el-tooltip>
            </template></el-input>
    </el-form-item>
    <el-form-item prop="searchUrl.bookId" label="书本ID规则">
        <el-input v-model="form.data.bookId">
            <template #append>
                <el-tooltip effect="dark" :content="book?.bookId" placement="top-start">
                    <el-icon class="el-icon-success" v-if="attrIsSuccess('bookId')">
                        <SuccessFilled />
                    </el-icon>
                    <el-icon class="el-icon-failed" v-else>
                        <CircleCloseFilled />
                    </el-icon>
                </el-tooltip>
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="searchUrl.bookUri" label="章节列表规则">
        <el-input v-model="form.data.bookUri">
            <template #append>
                <el-icon class="el-icon-success" v-if="attrIsSuccess('bookId')">
                    <SuccessFilled />
                </el-icon>
                <el-icon class="el-icon-failed" v-else>
                    <CircleCloseFilled />
                </el-icon>
            </template>
        </el-input>
    </el-form-item>
</template>

<style scoped>
.el-icon-success {
    color: #409eff;
}

.el-icon-failed {
    color: #f56c6c;
}
</style>