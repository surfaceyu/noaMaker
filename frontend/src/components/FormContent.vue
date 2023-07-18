  
<script setup lang="ts">
import { models } from '@wailsjs/go/models';
import { Document } from '@element-plus/icons-vue'

const form = defineProps({
    data: {
    type: models.BookSiteContentUri,
    required: true
  }, 
})
</script>

<template>
    <el-text> <el-icon> <Document /> </el-icon>内容搜索规则</el-text>
    <el-form-item prop="contentUrlBookId" label="搜索规则">
        <el-input v-model="form.data.uri" />
    </el-form-item>
    <el-form-item prop="contentUrlBookUri" label="内容规则">
        <el-input v-model="form.data.content" />
    </el-form-item>
</template>
