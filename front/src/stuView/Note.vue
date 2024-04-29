<script setup>
import avatarJpg from "@/assets/img/avatar.jpg";
import {onMounted, reactive} from "vue";
import store from "@/stores/index.js";
import axios from "axios";
import {useRouter} from "vue-router";

const router = useRouter()
const state = reactive({
  noteList: [],
  noteData: "",
})


onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/stu/notes", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
    }
  }).then((res) => {
    state.noteList = res.data.result
  })
})

const handleCollapseChange = (name) => {
  axios.post("http://127.0.0.1:8080/api/v1/stu/getNote", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
      "dotId": name.toString()
    }
  }).then((res) => {
    state.noteData = res.data.result
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
  <div class="note">
    <div class="top">
      <el-card shadow="never">
        <div class="info">
          <div class="info-left">
            <div class="pic">
              <el-avatar :size="100" :src="avatarJpg"/>
            </div>
            <div class="usrInfo">
              <div class="name">
                <el-icon class="icons"><User /></el-icon>
                <span class="infoItem">{{ store.state.userMsg.Yb_realname }}</span>
              </div>

              <div class="schoolInfo">
                <el-icon class="icons" style="padding-top: 7px"><Postcard /></el-icon>
                <div class="classinfo">
                  <span>{{ store.state.userMsg.Yb_collegename }}</span>
                  <span>{{ store.state.userMsg.Yb_classname }}</span>
                </div>
              </div>
            </div>
          </div>

          <div class="noteNum">
            <el-statistic title="编写笔记数量" :value="state.noteList.length" class="startNum"/>
          </div>
        </div>
      </el-card>
    </div>

    <el-divider />

    <div class="notes">
      <el-collapse accordion @change="handleCollapseChange" >
        <el-collapse-item class="collapse" v-for="(item, index) in state.noteList" :title="item.title" :name="item.dotID">
          <el-button @click="view(item.notePath, item.dotID)">进入</el-button>
          <mavon-editor id="editor" :toolbars="state.showToolbar" v-model="state.noteData" :subfield="false" :editable="false" defaultOpen="preview" codeStyle="github-dark-dimmed" :toolbarsFlag="false"></mavon-editor>
        </el-collapse-item>
      </el-collapse>
    </div>
  </div>
</template>

<style scoped>
.note{
  display: flex;
  flex-direction: column;
  padding-left: 25px;
  padding-right: 25px;
  height: 100%;
  width: 100%;
}
.body{
  padding-top: 6px;
}

.info{
  display: flex;
  flex-direction: row;
  justify-content: start;
  align-items: center;
}

.info-left{
  display: flex;
  padding-right: 100px;
}

.pic{
  padding-right: 23px;
}

.usrInfo{
  display: flex;
  flex-direction: column;
  align-items: start;
  justify-content: center;
}

.name{
  padding-bottom: 19px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.icons{
  transform: scale(1.5);
}

.infoItem{
  padding-left: 13px;
  font-size: 20px;
}

.schoolInfo{
  display: flex;
  justify-content: center;
  align-items: start;
}

.classinfo{
  font-size: 20px;
  display: flex;
  padding-left: 13px;
  flex-direction: column;
  justify-content: start;
}

.noteNum{
  display: flex;
  padding-left: 50px;
}

.top{
  padding-bottom: 10px;
}

/deep/ .el-collapse-item__header{
    font-size: 17px;
}
</style>