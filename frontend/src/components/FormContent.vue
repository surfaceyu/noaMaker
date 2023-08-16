  
<script setup lang="ts">
import { models } from '@wailsjs/go/models';
import { Document, Search } from '@element-plus/icons-vue'
import { onUpdated, ref } from 'vue';
import _, { isEmpty } from 'lodash';
import { SearchBookByRule, SearchChapter, SearchContent } from '@backend/BookHandler';

const form = defineProps({
    data: {
        type: Object as () => models.BookSite,
        required: true
    },
})

const contentData = ref<models.Content>()

async function handleSearch() {
    contentData.value = undefined;
    form.data.checkStep = 0;

    const books = await SearchBookByRule("剑仙", form.data.searchUrl)
    if (books && books.length > 0) {
        form.data.checkStep = 1;
        const c = await SearchChapter(books[0], form.data.chapterUrl)
        
        _.filter(c, (_, index) => index <= 10).forEach(async v => {
            if (contentData.value) {
                return
            }
            form.data.checkStep = 2;
            const result = await SearchContent(v, form.data.contentUrl)
            if (!_.isEmpty(result?.content)) {
                contentData.value = result
                form.data.checkStep = 3;
            }
        })
    }
}

onUpdated(_.debounce(() => {
    handleSearch()
}, 300))
</script>

<template>
    <el-text> <el-icon>
            <Document />
        </el-icon>内容搜索规则</el-text>
    <el-form-item prop="contentUrlBookId" label="搜索规则">
        <el-input v-model="form.data.contentUrl.uri">
            <template #append>
                <el-button :icon="Search" @click="handleSearch" />
            </template>
        </el-input>
    </el-form-item>
    <el-form-item prop="contentUrlBookUri" label="内容规则">
        <el-input v-model="form.data.contentUrl.content">
            <template #append>                
                <el-tooltip effect="dark" :content="contentData?.content.slice(1, 50)" placement="top-start">
                    <el-icon class="el-icon-success" v-if="_.get(contentData, 'content')">
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