<script lang="ts" setup>
import { computed, reactive, ref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus'
import { models } from '@wailsjs/go/models';
import { House } from '@element-plus/icons-vue'
import FormChapter from '@components/FormChapter.vue'
import FormBook from '@components/FormBook.vue';
import FormContent from '@components/FormContent.vue'
import { bookSiteRules, editRuleRow } from "@scripts/store";

const dialogArrayData = defineProps({
    submit: {
        type: Function,
        required: true,
    },
})

const isShow = ref(false)
const ruleFormRef = ref<FormInstance>()

defineExpose({
    isShow
})

const form = computed(() => {
    return editRuleRow.value || reactive(new models.BookSite({
    name: '',
    uri: '',
    searchUrl: new models.BookSiteUri(),
    chapterUrl: new models.BookSiteChapterUri(),
    contentUrl: new models.BookSiteContentUri(),
    }))
})

const nameValidator = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
    if (value.length < 2) {
        return callback(new Error('长度不能小于2！'));
    }
    if (bookSiteRules.value.find(book => book.name == value)) {
        return callback(undefined); 
    }
    callback(new Error(`源网址 <${value}> 没有记录,请先去添加！`));
}

const inputRules = reactive<FormRules<typeof form>>({
    name: [{ validator: nameValidator, trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
    formEl?.validate((valid) => {
        if (valid) {
            isShow.value = false;
            dialogArrayData.submit(form.value)
            return true;
        } else {
            return false
        }
    })
}
</script>

<template>
    <el-dialog v-model="isShow" title="编辑书源">
        <el-form ref="ruleFormRef" :model="form" :rules="inputRules" label-width="auto" label-position="right">
            <el-text> <el-icon> <House /> </el-icon>源站</el-text>
            <el-form-item prop="name" label="名称">
                <el-input v-model="form.name" />
            </el-form-item>
            <el-form-item prop="uri" label="网址">
                <el-input v-model="form.uri" />
            </el-form-item>
            <el-divider />
            <FormBook :data=form.searchUrl />
            <el-divider />
            <FormChapter :data=form />
            <el-divider />
            <FormContent :data=form />
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="isShow = false">取消</el-button>
                <el-button type="primary" @click="submitForm(ruleFormRef)">确认</el-button>
            </span>
        </template>
    </el-dialog>
</template>

<style scoped>
.dialog-footer button:first-child {
    margin-right: 10px;
}
</style>
