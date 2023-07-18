  
<script setup lang="ts">
import { models } from '@wailsjs/go/models';
import { Collection } from '@element-plus/icons-vue'

const form = defineProps({
    data: {
    type: models.BookSiteUri,
    required: true
  }, 
})

</script>

<template>
    <el-text> <el-icon> <Collection /> </el-icon>书本搜索规则</el-text>
    <el-form-item prop="searchUrl.uri" label="搜索规则" >
        <el-input v-model="form.data.uri" />
    </el-form-item>
    <el-form-item prop="searchUrl.bookList" label="列表规则">
        <el-input v-model="form.data.bookList" />
    </el-form-item>
    <el-form-item prop="searchUrl.author" label="作者规则">
        <el-input v-model="form.data.author" />
    </el-form-item>
    <el-form-item prop="searchUrl.bookName" label="书名规则">
        <el-input v-model="form.data.bookName" />
    </el-form-item>
    <el-form-item prop="searchUrl.bookId" label="书本ID规则">
        <el-input v-model="form.data.bookId" />
    </el-form-item>
    <el-form-item prop="searchUrl.bookUri" label="章节列表规则">
        <el-input v-model="form.data.bookUri" />
    </el-form-item>
</template>
