<script lang="ts" setup>
import { computed, ref } from 'vue';
import { Check, Delete } from '@element-plus/icons-vue'
import { models } from '@wailsjs/go/models';
import { AddSites, DelSites, EditSites } from '@backend/BookSitesHandler';
import DialogSiteAdd from '@components/DialogSiteAdd.vue'
import DialogSiteEdit from '@components/DialogSiteEdit.vue';
import { bookSiteRules, editRuleRow, useRule } from "@scripts/store";

const search = ref('')
const dialogSiteEdit = ref<InstanceType<typeof DialogSiteEdit>>()
const dialogFormAdd = ref<InstanceType<typeof DialogSiteAdd>>()

const siteStepType = ["danger", "warning", "warning"];
const siteStepText = ["搜索", "章节", "内容", "完成"]

const filterTableData = computed(() => {
    return bookSiteRules.value.filter(data =>
        !search.value || data.name.toLowerCase().includes(search.value.toLowerCase()))
})

const handleAdd = async () => {
    dialogFormAdd.value!.isShow = true
}

const handleEdit = async (index: number, row: models.BookSite) => {
    dialogSiteEdit.value!.isShow = true;
    editRuleRow.value = row;
}

const handleDelete = async (index: number, row: models.BookSite) => {
    await DelSites(row)
    useRule()
}

const doAdd = async (site: models.BookSite) => {
    await AddSites(site)
    useRule()
}

const doEdit = async (row: models.BookSite) => {
    await EditSites(row);
    useRule();
}

useRule()
</script>

<template>
    <h5 class="main-content-title">书源管理</h5>
    <!-- 按钮组合 新增 -->
    <!-- 书源列表 -->
    <el-table :data="filterTableData" style="width: 100%" stripe>
        <el-table-column type="index" width="50" />
        <el-table-column label="名称" prop="name" width="150" />
        <el-table-column label="网址" prop="uri" width="300" />
        <el-table-column label="状态">
            <template #default="scope">
                <el-tag :type="siteStepType[scope.row.checkStep] || 'success'" disable-transitions>
                    {{ siteStepText[scope.row.checkStep] }}
                </el-tag>
            </template>
        </el-table-column>
        <el-table-column align="right" width="240">
            <template #header>
                <el-row>
                    <el-col align="left" :span="4">
                        <el-button size="small" @click="handleAdd">新增</el-button>
                    </el-col>
                    <el-col :span="20">
                        <el-input v-model="search" size="small" placeholder="输入书名进行搜索" style="width: 160px;" />
                    </el-col>
                </el-row>
            </template>
            <template #default="scope">
                <el-button size="small" type="success" :icon="Check"
                    @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
                <el-button :disabled="scope.$index == 0" size="small" type="danger" :icon="Delete"
                    @click="handleDelete(scope.$index, scope.row)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
    <DialogSiteAdd :data="filterTableData" :submit="doAdd" ref="dialogFormAdd" />
    <DialogSiteEdit :submit="doEdit" ref="dialogSiteEdit" />
</template>

<style scoped>
.main-content-title {
    text-align: left;
    padding-left: 20px;
    font-size: large;
    color: #e7ac08;
}
</style>
