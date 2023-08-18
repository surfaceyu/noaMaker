  
<script setup lang="ts">
import _ from 'lodash';
import { models } from '@wailsjs/go/models';
import { SearchBookByRule, SearchChapter } from '@backend/BookHandler';
import { editRule } from "@scripts/store";

const chapterData = editRule.chapterData

const form = defineProps({
    data: {
        type: Object as () => models.BookSite,
        required: true
    },
})

/**
 * 搜索书籍
 */
 async function handleSearch() {
    chapterData.value = undefined

    const books = await SearchBookByRule("剑仙", form.data.searchUrl)
    if (books && books.length > 0) {
        const c = await SearchChapter(books[0], form.data.chapterUrl)
        if (c && c.length > 0) {
            chapterData.value = c[0]
        }
    }
}
</script>

<template>
    <el-text> <el-icon> <Files /> </el-icon>章节搜索规则</el-text>
    <el-form-item prop="chapterUrlUri" label="搜索规则">
        <el-input v-model="form.data.chapterUrl.uri" >
            <template #append>
                <el-button icon="Search" @click="handleSearch" />
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="chapterUrlList" label="列表规则">
        <el-input v-model="form.data.chapterUrl.chapterList">
            <template #append>
                <el-tooltip effect="dark" :content="chapterData?.chapterName" placement="top-start">
                    <el-icon class="el-icon-success" v-if="_.get(chapterData,'chapterName')">
                        <SuccessFilled />
                    </el-icon>
                    <el-icon class="el-icon-failed" v-else>
                        <CircleCloseFilled />
                    </el-icon>
                </el-tooltip>
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="chapterUrlBookName" label="章节名称规则">
        <el-input v-model="form.data.chapterUrl.chapterName">
            <template #append>
                <el-tooltip effect="dark" :content="chapterData?.chapterName" placement="top-start">
                    <el-icon class="el-icon-success" v-if="_.get(chapterData,'chapterName')">
                        <SuccessFilled />
                    </el-icon>
                    <el-icon class="el-icon-failed" v-else>
                        <CircleCloseFilled />
                    </el-icon>
                </el-tooltip>
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="chapterUrlBookId" label="章节ID规则">
        <el-input v-model="form.data.chapterUrl.chapterId">
            <template #append>
                <el-tooltip effect="dark" :content="chapterData?.chapterId" placement="top-start">
                    <el-icon class="el-icon-success" v-if="_.get(chapterData,'chapterId')">
                        <SuccessFilled />
                    </el-icon>
                    <el-icon class="el-icon-failed" v-else>
                        <CircleCloseFilled />
                    </el-icon>
                </el-tooltip>
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="chapterUrlBookUri" label="章节内容规则">
        <el-input v-model="form.data.chapterUrl.chapterUrl">
            <template #append>
                <el-tooltip effect="dark" :content="chapterData?.chapterId" placement="top-start">
                    <el-icon class="el-icon-success" v-if="_.get(chapterData,'chapterId')">
                        <SuccessFilled />
                    </el-icon>
                    <el-icon class="el-icon-failed" v-else>
                        <CircleCloseFilled />
                    </el-icon>
                </el-tooltip>
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