<script setup>
import {reactive, computed, onMounted} from 'vue'
import { Search, Refresh, Plus } from '@element-plus/icons-vue'
import { useRouter } from "vue-router";
import axios from "axios";
import store from "@/stores/index.js";
import { formatETableCreateTime, formatETableUpdateTime} from "@/utils/formatTime.js";

const router = useRouter()

// 格式化：状态码 --> 状态名
const formatState = (row, colum) =>{
  const foundState = states.find(item => item.state === parseInt(row.state))
  return foundState ? foundState.value : "状态未知";
}
const getFormatDate = () => {
  const currentDate = new Date()
  return currentDate.getFullYear() + "-" + currentDate.getMonth() + "-" + currentDate.getDate()
}

const tagOptions = computed(() => {
  return store.state.tags
})

const state = reactive({
  newSelectTags: [],
  searchForm: {
    title: "",
    state: "",
    fuzzy: false
  },
  dotData: [],
  centerDialog: false,
  dialogForm: {
    tags: [],
    title: '',
    description: "",
    newTags: '',
    createTime: getFormatDate(),
    updateTime: getFormatDate()
  },
  dialogRule: {
    title: [{
      required: true,
      message: "请填写对应信息",
      trigger: "change",
      type: "string",
    }],
    description: [{
      required: true,
      message: "请填写对应信息",
      trigger: "change",
      type: "string",
    }],
    tags: [{
      required: true,
      message: "请选择标签",
      trigger: "change",
    }],
  }
})

const handleTagSelectionChange = (selectedValues) => {
  state.dialogForm.tags = selectedValues.join(',');
}

const resetForm = ()=> {
  state.searchForm.title = ''
  state.searchForm.state = ''
  state.searchForm.fuzzy = false
  axios.post("http://127.0.0.1:8080/api/v1/dots", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "Id": store.state.userMsg.Yb_studentid
    }
  }).then((res) => {
    state.dotData = res.data.result
  })
}
const resetDialog = () => {
  state.centerDialog = false
  state.dialogForm.title = ""
  state.dialogForm.tags = []
  state.dialogForm.description= ""
  state.dialogForm.newTags = ""
}
// 稿件的所有状态映射
const states = [
  {
    state: 0,
    value: "上架"
  },
  {
    state: 1,
    value: "下架"
  }
]

onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/dots", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "Id": store.state.userMsg.Yb_studentid
    }
  }).then((res) => {
    state.dotData = res.data.result
  })
})

const addItem = () => {
  axios.post("http://127.0.0.1:8080/api/v1/createDraft", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "title": state.dialogForm.title,
      "description": state.dialogForm.description,
      "tags": state.dialogForm.tags,
      "id": store.state.userMsg.Yb_studentid
    }
  }).then((res) => {
    if (res.data.result === true){
      location.reload()
    }else {
      alert("create is error")
    }
  })
}

const editDraft = (draftPath, draftId) => {
  router.push({
    path: '/Tete',
    query: {
      "draftPath": draftPath,
      "draftId": draftId
    }
  })
}

const deleteDot = (dotId) => {
  axios.post("http://127.0.0.1:8080/api/v1/dotDel", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "dotId": dotId.toString(),
    }
  }).then((res) => {
    if (res.data.result === true){
      state.dotData = state.dotData.filter(obj => obj.dotID !== dotId);
    }else {
      console.log(res.data.result)
    }
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

const setDot = (dotId, tar) => {
  axios.post("http://127.0.0.1:8080/api/v1/setDotState", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "dotId": dotId.toString(),
      "state": tar
    }
  }).then((res) => {
    if (res.data.result === true){
      for (let i = 0; i < state.dotData.length; i++) {
        if (state.dotData[i].dotID === dotId) {
          state.dotData[i].state = tar;
          break;
        }
      }
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
          label-width="70px"
          label-position="right"
          label-suffix=" : "
        >
          <el-row gutter=20 >
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

          <el-row class="btn" gutter=20 justify="end">

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
        <el-table-column fixed prop="dotID" label="序号" min-width="100" />
        <el-table-column fixed prop="title" label="标题" min-width="170" />
        <el-table-column prop="description" label="描述" min-width="240" truncated/>
        <el-table-column prop="tags" label="标签" min-width="220"/>
        <el-table-column prop="state" label="状态" min-width="100" :formatter="formatState"/>
        <el-table-column prop="createTime" label="创建时间" min-width="200" :formatter="formatETableCreateTime" />
        <el-table-column prop="updateTime" label="更新时间" min-width="200" :formatter="formatETableUpdateTime"/>
        <el-table-column fixed="right" label="操作" min-width="220" align="center">
          <template #default="scope">
            <el-button v-if="scope.row.state === '1'" link type="primary" @click="editDraft(scope.row.nodeTreePath, scope.row.dotID)">编辑</el-button>
            <el-button v-else link type="primary" :disabled="true">编辑</el-button>
            <el-button v-if="scope.row.state === '1'" link type="primary" @click="setDot(scope.row.dotID, '0')">上架</el-button>
            <el-button v-else link type="primary" :disabled="true">上架</el-button>
            <el-button v-if="scope.row.state === '0'" link type="primary" @click="setDot(scope.row.dotID, '1')">下架</el-button>
            <el-button v-else link type="primary" :disabled="true">下架</el-button>
            <el-button v-if="scope.row.state === '1'" link type="primary" @click="deleteDot(scope.row.dotID)">删除</el-button>
            <el-button v-else link type="primary" :disabled="true">删除</el-button>
          </template>
        </el-table-column>
      </el-table>


      <el-button size="large" class="mt-4" style="width: 100%" @click="() => state.centerDialog = true">
        <el-icon class="createIcon"><Plus/></el-icon>
        <span>新增草稿</span>
      </el-button>

      <el-dialog
        v-model="state.centerDialog"
        title="新增草稿"
        width="500"
        align-center
        :lock-scroll="false"
      >
        <el-form label-position="top" :rules="state.dialogRule" :model="state.dialogForm">

          <div class="draft-form">
            <div class="draft-form-item">
              <div>
                <el-form-item label="标题：" prop="title">
                  <el-input v-model="state.dialogForm.title"/>
                </el-form-item>
              </div>

              <div>
                <el-form-item label="创建日期：" >
                  <el-input v-model="state.dialogForm.createTime" disabled/>
                </el-form-item>
              </div>
            </div>

            <div class="draft-form-item">
              <div>
                <el-form-item label="描述：" prop="description">
                  <el-input
                    v-model="state.dialogForm.description"
                    type="textarea"
                    :rows="9"
                    maxlength="50"
                    show-word-limit
                    placeholder="描述你的稿件内容"/>
                </el-form-item>
              </div>

              <div class="draft-form-bottom-right">

                <div>
                  <el-form-item label="标签：" prop="tags">
                    <el-select
                      v-model="state.newSelectTags"
                      multiple
                      collapse-tags
                      collapse-tags-tooltip
                      placeholder="Select"
                      style="width: 200px"
                      @change="handleTagSelectionChange"
                    >
                      <el-option
                        v-for="item in tagOptions"
                        :key="item.Name"
                        :label="item.Name"
                        :value="item.Name"
                      />
                    </el-select>
                  </el-form-item>
                </div>

                <div class="btns">
                  <el-button class="form-btn" type="primary" round @click="addItem">提交</el-button>
                  <el-button class="form-btn" type="info" round @click="resetDialog">取消</el-button>
                </div>
              </div>
            </div>
          </div>
        </el-form>
      </el-dialog>
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

.draft-form{
  display: flex;
  flex-direction: column;
  justify-content: space-around;
}

.draft-form-item{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.draft-form-bottom-right{
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.btns{
  display: flex;
  justify-content: end;
}

.form-btn{
  margin-bottom: 20px;
  margin-right: 14px;
}
</style>