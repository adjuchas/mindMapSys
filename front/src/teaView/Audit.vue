<script setup>
import {Plus} from "@element-plus/icons-vue";
import {onMounted, reactive} from "vue";
import axios from "axios";
import store from "@/stores/index.js";
import {useRouter} from "vue-router";

const state = reactive({
  draftData: []
})
const router = useRouter()
onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/audit", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
    }
  }).then((res) => {
    state.draftData = res.data.result
  })
})

const toAudit = (nodePath, dotId) => {
  router.push({
    path: '/check',
    query: {
      "nodePath": nodePath,
      "dotId": dotId,
    }
  })
}
const stateFormat = (row,column,cellValue) => {
  if(!cellValue) return "";
  if(cellValue.length >12){
    return cellValue.slice(0,12) + "......"
  }
  return cellValue
}
</script>

<template>
  <div class="main">
    <el-table :data="state.draftData" stripe highlight-current-row table-layout="fixed" max-height="414">
      <el-table-column fixed prop="draftID" label="序号" min-width="100" />
      <el-table-column fixed prop="title" label="标题" min-width="170" />
      <el-table-column fixed prop="auth" label="作者" min-width="170" />
      <el-table-column prop="description" label="描述" min-width="240" :formatter="stateFormat"/>
      <el-table-column prop="tags" label="标签" min-width="200" :formatter="formatState"/>
      <el-table-column fixed="right" label="操作" min-width="220" align="center">
        <template #default="scope">
          <el-button link @click="toAudit(scope.row.nodeTreePath, scope.row.draftID)" type="primary">查看</el-button>
        </template>
      </el-table-column>
    </el-table>


  </div>
</template>

<style scoped>
.main{
  margin-right: 9px;
  margin-left: 9px;
}
</style>