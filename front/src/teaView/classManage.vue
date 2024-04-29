<script setup>
import {Plus} from "@element-plus/icons-vue";
import {onMounted, reactive} from "vue";
import store from "../stores/index.js";
import axios from "axios";
import { formatETableCreateTime } from "@/utils/formatTime.js";

const getFormatDate = () => {
  const currentDate = new Date()
  return currentDate.getFullYear() + "-" + currentDate.getMonth() + "-" + currentDate.getDate()
}

const state = reactive({
  name: store.state.userMsg.Yb_realname,
  createTime: getFormatDate(),
  centerDialog: false,
  tagsData: [],
  dialogRule: {
    tagName: [{
      required: true,
      message: "请填写新增的分类名称",
      trigger: "change",
      type: "string",
    }]
  },
  dialogForm: {
    tagName: "",
  }
})

onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/getClassifyInfo", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {}
  }).then((res) => {
    state.tagsData = res.data.result
  })
})

const resetForm = ()=> {
  state.centerDialog = false
  state.dialogForm.tagName = ""
}

const createNewTag = () => {
  axios.post("http://127.0.0.1:8080/api/v1/createClassify", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "tagName": state.dialogForm.tagName,
      "id": store.state.userMsg.Yb_studentid,
    }
  }).then((res) => {
    if (res.data.result === true){
      location.reload()
    }
  })
}
</script>

<template>
  <div class="main">
    <el-table :data="state.tagsData" stripe highlight-current-row table-layout="fixed" max-height="calc(100%)">
      <el-table-column fixed prop="tagID" label="序号" min-width="100" />
      <el-table-column fixed prop="name" label="分类名" min-width="170" />
      <el-table-column prop="dotNum" label="上架作品数" min-width="240" truncated/>
      <el-table-column prop="leaderName" label="创建人" min-width="240" truncated/>
      <el-table-column prop="createTime" label="创建时间" min-width="150" :formatter="formatETableCreateTime" />
    </el-table>

    <el-button size="large" class="mt-4" style="width: 100%" @click="() => state.centerDialog = true">
      <el-icon class="createIcon"><Plus/></el-icon>
      <span>新增一个分类</span>
    </el-button>

    <el-dialog
      v-model="state.centerDialog"
      title="新增标签"
      width="500"
      align-center
      :lock-scroll="false"
    >
      <el-form label-position="top" :rules="state.dialogRule" :model="state.dialogForm">
        <div class="tagInfo">
          <div>
            <el-form-item label="分类名：" prop="tagName">
              <el-input v-model="state.dialogForm.tagName"/>
            </el-form-item>
          </div>
          <div>
            <el-form-item label="创建者：">
              <el-input v-model="state.name" :disabled="true" />
            </el-form-item>
          </div>
        </div>
        <div>
          <el-form-item label="创建日期：" >
            <el-input v-model="state.createTime" disabled/>
          </el-form-item>
        </div>

        <div class="btns">
          <el-button class="btn" type="info" round @click="resetForm">取消</el-button>

          <el-button class="btn" type="primary" round @click="createNewTag">新建</el-button>
        </div>
      </el-form>
    </el-dialog>
  </div>
</template>

<style scoped>
.main{
  padding-left: 9px;
  padding-right: 9px;
}

.tagInfo{
  display: flex;
  justify-content: space-between;
}

.btns{
  display: flex;
  flex-direction: row-reverse;
}

.btn{
  margin-top: 10px;
  margin-right: 9px;
  margin-left: 9px;
}
</style>