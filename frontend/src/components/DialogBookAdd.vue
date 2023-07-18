<script lang="ts" setup>
import { inject, reactive, ref } from 'vue';
import { type Ref } from 'vue';
import { backend } from '@wailsjs/go/models';
import type { FormInstance, FormRules } from 'element-plus'

const formLabelWidth = '140px'

const dialogArrayData = defineProps({
    data: Array as () => backend.Book[],
})

const dialogFormVisible: Ref<Boolean> | undefined = inject('dialogFormVisible')
const form = reactive({
    name: '',
})

const nameValidator = (rule: any, value: string, callback: (arg0: Error | undefined) => void) => {
    if (value.length < 2) {
        return callback(new Error('长度不能小于2'));
    }
    if (dialogArrayData.data?.find(book => book.name == value)) {
        return callback(new Error(`书本 <${value}> 已有重复名称`));
    }
    callback(undefined);
}

const ruleFormRef = ref<FormInstance>()
const ruleForm = reactive({
    name: '',
})
const inputRules = reactive<FormRules<typeof ruleForm>>({
    name: [{ validator: nameValidator, trigger: 'blur' }],
})

const submitForm = (formEl: FormInstance | undefined) => {
    if (!formEl) return
    formEl.validate((valid) => {
        if (valid) {
            if (dialogFormVisible) dialogFormVisible.value = false
            return true;
        } else {
            console.log('error submit!')
            return false
        }
    })
}
</script>

<template>
    <el-dialog v-model="dialogFormVisible" title="新增书本">
        <el-form ref="ruleFormRef" :model="form" :rules="inputRules">
            <el-form-item prop="name" label="书籍名称" :label-width="formLabelWidth">
                <el-input v-model="form.name" autocomplete="off" />
            </el-form-item>
            <!-- <el-form-item label="Zones" :label-width="formLabelWidth">
        <el-select v-model="form.region" placeholder="Please select a zone">
          <el-option label="Zone No.1" value="shanghai" />
          <el-option label="Zone No.2" value="beijing" />
        </el-select>
      </el-form-item> -->
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
.el-button--text {
    margin-right: 15px;
}

.el-select {
    width: 300px;
}

.el-input {
    width: 300px;
}

.dialog-footer button:first-child {
    margin-right: 10px;
}</style>
