<script lang="ts" setup>
import { computed, provide, ref } from 'vue';
import { type Ref } from 'vue';
import { models } from '@wailsjs/go/models';
import { GetSites, AddSites, DelSites, EditSites } from '@wailsjs/go/backend/BookSitesHandler';
import DialogSiteAdd from '@components/DialogSiteAdd.vue'
import DialogSiteEdit from '@components/DialogSiteEdit.vue';

const search = ref('')
const editIndex = ref(0)
const tableData: Ref<models.BookSite[]> = ref([]);

// GetSites().then(result => {
//   tableData.value = result;
// })

const updateSitesData = async () => {
  tableData.value = await GetSites()
}

updateSitesData()

const filterTableData = computed(() => {
  return tableData.value.filter(data =>
    !search.value || data.name.toLowerCase().includes(search.value.toLowerCase()))
})

const handleEdit = async (index: number, row: models.BookSite) => {
  dialogFormEditVisible.value = !dialogFormEditVisible.value
  editIndex.value = index;
  updateSitesData()
}

const handleDelete = async (index: number, row: models.BookSite) => {
  await DelSites(row)
  updateSitesData()
}

const doAdd = async (site: models.BookSite) => {
  await AddSites(site)
  updateSitesData()
}

const dialogFormVisible = ref(false);
provide('dialogFormVisible', dialogFormVisible);
const dialogFormEditVisible = ref(false);
provide('dialogFormEditVisible', dialogFormEditVisible);
</script>

<template>
  <h5 class="main-content-title">书源管理</h5>
  <!-- 按钮组合 新增 -->
  <!-- 书源列表 -->
  <el-table :data="filterTableData" style="width: 100%" stripe>
    <el-table-column label="名称" prop="name" width="150" />
    <el-table-column label="网址" prop="uri"/>
    <el-table-column align="right" width="240">
      <template #header>
        <el-row>
          <el-col align="left" :span="4">
            <el-button size="small" @click="dialogFormVisible = true">新增</el-button>
          </el-col>
          <el-col :span="20">
            <el-input v-model="search" size="small" placeholder="输入书名进行搜索" style="width: 160px;" />
          </el-col>
        </el-row>
      </template>
      <template #default="scope">
        <el-button size="small" @click="handleEdit(scope.$index, scope.row)">编辑</el-button>
        <el-button size="small" type="danger" @click="handleDelete(scope.$index, scope.row)">删除</el-button>
      </template>
    </el-table-column>
  </el-table>
  <DialogSiteAdd :data="filterTableData" :submit="doAdd" />
  <DialogSiteEdit :data="filterTableData" :submit="updateSitesData" :index="editIndex" />
</template>

<style scoped>
.main-content-title {
  text-align: left;
  padding-left: 20px;
  font-size: large;
  color: #e7ac08;
}
</style>
