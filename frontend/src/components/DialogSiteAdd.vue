<script lang="ts" setup>
import { reactive, ref } from 'vue';
import type { FormInstance, FormRules } from 'element-plus'
import { models } from '@wailsjs/go/models';
import { House } from '@element-plus/icons-vue'
import FormChapter from '@components/FormChapter.vue'
import FormBook from '@components/FormBook.vue';
import FormContent from '@components/FormContent.vue'

const dialogArrayData = defineProps({
    data: {
        type: Array as () => models.BookSite[],
        required: true,
    },
    submit: {
        type: Function,
        required: true,
    },
})

const isShow = ref(false)
const ruleFormRef = ref<FormInstance>()
const labelPosition = ref('right')

defineExpose({
    isShow
})

const form = reactive(new models.BookSite({
    name: '',
    uri: '',
    searchUrl: new models.BookSiteUri(),
    chapterUrl: new models.BookSiteChapterUri(),
    contentUrl: new models.BookSiteContentUri(),
}))

const nameValidator = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
    if (value.length < 2) {
        return callback(new Error('长度不能小于2'));
    }
    if (dialogArrayData.data.find(book => book.name == value)) {
        return callback(new Error(`源网址 <${value}> 已有重复名称`));
    }
    callback(undefined);
}

const inputRules = reactive<FormRules<typeof form>>({
    name: [{ validator: nameValidator, trigger: 'blur' }],
    // "searchUrl.uri": [{ validator: nameValidator, trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            isShow.value = false;
            dialogArrayData.submit(form)
            return true;
        } else {
            return false
        }
    })
}
</script>

<template>
    <el-dialog v-model="isShow" title="新增书源">
        <el-form ref="ruleFormRef" :model="form" :rules="inputRules" label-width="auto" :label-position="labelPosition">
            <el-text> <el-icon>
                    <House />
                </el-icon>源站</el-text>
            <el-form-item prop="name" label="名称">
                <el-input v-model="form.name" />
            </el-form-item>
            <el-form-item prop="uri" label="网址">
                <el-input v-model="form.uri" />
            </el-form-item>
        </el-form>
        <template #footer>
            <span class="dialog-footer">
                <el-button @click="isShow=false">取消</el-button>
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
