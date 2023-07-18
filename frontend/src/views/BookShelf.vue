<script lang="ts" setup>
import { computed, provide, ref } from 'vue';
import { type Ref } from 'vue';
import { backend } from '@wailsjs/go/models';
import DialogBookAdd from '@components/DialogBookAdd.vue'

const search = ref('')
const tableData: Ref<backend.Book[]> = ref([]);

const filterTableData = computed(() => {
  return tableData.value.filter(data =>
    !search.value || data.name.toLowerCase().includes(search.value.toLowerCase()))
})

const handleEdit = (index: number, row: backend.Book) => {
  console.log(index, row);
}
const handleDelete = (index: number, row: backend.Book) => {
  console.log(index, row);
}

const dialogFormVisible = ref(false);

provide('dialogFormVisible', dialogFormVisible);
</script>

<template>
  <h5 class="main-content-title">书架管理</h5>
  <!-- 按钮组合 新增 -->
  <!-- 书源列表 -->
  <el-table :data="filterTableData" style="width: 100%" stripe>
    <el-table-column label="名称" prop="name" width="150" />
    <el-table-column label="作者" prop="author" width="100" />
    <el-table-column label="介绍" prop="desc"  />
    <el-table-column align="right" width="240">
      <template #header>
        <el-row>
          <el-col align="left" :span="4">
            <el-button size="small" @click="dialogFormVisible = true">新增</el-button>
          </el-col>
          <el-col :span="20">
            <el-input v-model="search" size="small" placeholder="输入名称进行搜索" style="width: 160px;" />
          </el-col>
        </el-row>
      </template>
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
        <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <DialogBookAdd :data="tableData"/>
</template>

<style scoped>
.main-content-title {
  text-align: left;
  padding-left: 20px;
  font-size: large;
  color: #e7ac08;
}
</style>
