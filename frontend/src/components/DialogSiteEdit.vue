<script lang="ts" setup>
import { computed, inject, reactive, ref } from 'vue';
import { type Ref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus'
import { models } from '@wailsjs/go/models';
import { House } from '@element-plus/icons-vue'
import FormChapter from '@components/FormChapter.vue'
import FormBook from '@components/FormBook.vue';
import FormContent from '@components/FormContent.vue'
import { EditSites } from '@wailsjs/go/backend/BookSitesHandler';

const dialogArrayData = defineProps({
    data: {
        type: Array<models.BookSite>,
        required: true,
    },
    index: {
        type: Number,
        required: true,
    },
    submit: {
        type: Function,
        required: true,
    },
})

const dialogFormVisible: Ref<Boolean> | undefined = inject('dialogFormEditVisible')
const form = computed(() => {
    return dialogArrayData.data[dialogArrayData.index] || reactive(new models.BookSite({
    name: '',
    uri: '',
    searchUrl: new models.BookSiteUri(),
    chapterUrl: new models.BookSiteChapterUri(),
    contentUrl: new models.BookSiteContentUri(),
}))
})
const ruleFormRef = ref<FormInstance>()
const labelPosition = ref('right')

const nameValidator = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
    if (value.length < 2) {
        return callback(new Error('长度不能小于2'));
    }
    if (dialogArrayData.data.find(book => book.name == value)) {
        return callback(undefined); 
    }
    callback(new Error(`源网址 <${value}> 没有记录,请先添加`));
}

const inputRules = reactive<FormRules<typeof form>>({
    name: [{ validator: nameValidator, trigger: 'blur' }],
    // "searchUrl.uri": [{ validator: nameValidator, trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            if (dialogFormVisible) dialogFormVisible.value = false;
            doEdit(form.value)
            return true;
        } else {
            console.log('error submit!')
            return false
        }
    })
}

const doEdit = async (row: models.BookSite) => {
  await EditSites(row)
  dialogArrayData.submit()
}
</script>

<template>
    <el-dialog v-model="dialogFormVisible" title="编辑书源">
        <el-form ref="ruleFormRef" :model="form" :rules="inputRules" label-width="auto" :label-position="labelPosition">
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
            <FormChapter :data=form.chapterUrl />
            <el-divider />
            <FormContent :data=form.contentUrl />
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="dialogFormVisible = false">取消</el-button>
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
