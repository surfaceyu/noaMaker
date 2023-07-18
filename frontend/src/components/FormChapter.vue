  
<script setup lang="ts">
import { models } from '@wailsjs/go/models';
import { Files } from '@element-plus/icons-vue'

const form = defineProps({
    data: {
    type: models.BookSiteChapterUri,
    required: true
  }, 
})
</script>

<template>
    <el-text> <el-icon> <Files /> </el-icon>章节搜索规则</el-text>
    <el-form-item prop="chapterUrlUri" label="搜索规则">
        <el-input v-model="form.data.uri" />
    </el-form-item>
    <el-form-item prop="chapterUrlList" label="列表规则">
        <el-input v-model="form.data.chapterList" />
    </el-form-item>
    <el-form-item prop="chapterUrlBookName" label="章节名称规则">
        <el-input v-model="form.data.chapterName" />
    </el-form-item>
    <el-form-item prop="chapterUrlBookId" label="章节ID规则">
        <el-input v-model="form.data.chapterId" />
    </el-form-item>
    <el-form-item prop="chapterUrlBookUri" label="章节内容规则">
        <el-input v-model="form.data.chapterUrl" />
    </el-form-item>
</template>
