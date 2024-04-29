<script setup>
import {ref, reactive, computed, onMounted} from 'vue'
import { Search, Refresh, Plus } from '@element-plus/icons-vue'
import axios from "axios";
import store from "@/stores/index.js";
import {useRouter} from "vue-router";
import {formatETableUpdateTime} from "@/utils/formatTime.js";

const router = useRouter()
const formatState = (row, colum) => {
  const foundState = states.find(item => item.state === parseInt(row.state))
  return foundState ? foundState.value : "状态未知";
}

const states = [
  {
    state: 0,
    value: '展示'
  },
  {
    state: 1,
    value: "下架"
  }
]

const state = reactive({
  searchForm: {
    title: '',
    state: '',
    fuzzy: false
  },
  dotData: []
})


const TakeOff = (dotID) => {
  axios.post('http://127.0.0.1:8080/api/v1/setDotState', {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "dotId": dotID.toString(),
      "state": "1"
    }
  }).then((res) => {
    if (res.data.result === true){
      for (let i = 0; i < state.dotData.length; i++) {
        if (state.dotData[i].dotID=== dotID) {
          state.dotData[i].state = "1";
          break;
        }
      }
    }else {
      console.log(res.data)
    }
  })
}

onMounted(() => {
  axios.post('http://127.0.0.1:8080/api/v1/dots', {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
    }
  }).then((res) => {
    state.dotData = res.data.result
  })
})


const resetForm = ()=> {
  state.searchForm.title = ''
  state.searchForm.state = ''
  state.searchForm.fuzzy = false
  axios.post("http://127.0.0.1:8080/api/v1/dots", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
    }
  }).then((res) => {
    state.dotData = res.data.result
  })
}

const selectDot = ()=> {
  axios.post("http://127.0.0.1:8080/api/v1/dotSelect", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "title": state.searchForm.title,
      "state": state.searchForm.state.toString(),
      "fuzzy": state.searchForm.fuzzy.toString(),
      "id": store.state.userMsg.Yb_studentid,
    }
  }).then((res) => {
    state.dotData = []
    state.dotData = res.data.result

  })
}

const view = (nodeTreePath, dotId) => {
  router.push({
    path: '/Show',
    query: {
      "nodePath": nodeTreePath,
      "dotId": dotId
    }
  })
}

</script>

<template>
  <div class="draft">
    <div class="top">
      <div class="searchBorder">
        <el-form
          class="searchForm"
          :model="state.searchForm"
          :inline
          label-width="70px"
          label-position="right"
          label-suffix=" : "
        >
          <el-row gutter="20" >
            <el-col class="search" span="7">
              <el-form-item label="标题">
                <el-input type="text" size="large" placeholder="请输入标题" clearable v-model="state.searchForm.title" style="min-width: 260px"/>
              </el-form-item>
            </el-col>

            <el-col class="search" span="7">
              <el-form-item label="状态">
                <el-select v-model="state.searchForm.state" placeholder="所有（默认）" clearable size="large" style="min-width: 260px">
                  <el-option v-for="item in states" :label="item.value" :value="item.state"/>
                </el-select>
              </el-form-item>
            </el-col>

            <el-col span="7">
              <el-form-item label-width="127px" label="标题模糊搜索" class="fuzzy">
                <el-switch v-model="state.searchForm.fuzzy"/>
              </el-form-item>
            </el-col>
          </el-row>

          <el-row class="btn" gutter="20" justify="end">
            <el-col span="7">
              <el-button type="primary" size="large" @click="selectDot" :icon="Search">查询</el-button>
            </el-col>
            <el-col span="7">
              <el-button size="large" @click="resetForm" :icon="Refresh">重置</el-button>
            </el-col>
          </el-row>
        </el-form>
      </div>
    </div>
    <el-divider />

    <div class="main">
      <el-table :data="state.dotData" stripe highlight-current-row table-layout="fixed" max-height="414">
        <el-table-column fixed prop="dotID" label="序号"/>
        <el-table-column fixed prop="title" label="标题" min-width="170"/>
        <el-table-column prop="description" label="描述" min-width="240"/>
        <el-table-column prop="tags" label="标签" min-width="220"/>
        <el-table-column prop="state" label="状态" :formatter="formatState"/>
        <el-table-column prop="updateTime" label="更新时间" min-width="150" :formatter="formatETableUpdateTime"/>
        <el-table-column fixed="right" label="操作" min-width="220" align="center">
          <template #default="scope">
            <el-button v-if="scope.row.state === '0'" link type="primary" @click="view(scope.row.nodeTreePath, scope.row.dotID)">查看（以读者模式）</el-button>
            <el-button v-else link type="primary" :disabled="true">查看（以读者模式）</el-button>
            <el-button v-if="scope.row.state === '0'" link type="primary" @click="TakeOff(scope.row.dotID)">下架</el-button>
            <el-button v-else link type="primary" :disabled="true">下架</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<style scoped>
.searchBorder{
  border-style: solid ;
  border-width: 1px;
  border-color: #7febff;
  border-radius: 16px;
  margin-left: 12px;
  margin-right: 20px;
}

.searchForm{
  margin: 10px 20px 12px 0;
}

.el-divider{
  margin-left: 9px;
}

.main{
  margin-left: 9px;

}
</style>