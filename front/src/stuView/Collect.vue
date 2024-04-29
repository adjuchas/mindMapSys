<script setup>
import avatarJpg from "@/assets/img/avatar.jpg";
import {onMounted, reactive} from "vue";
import axios from "axios";
import store from "@/stores/index.js";
import {useRouter} from "vue-router";
import {formatUpdateTime} from "@/utils/formatTime.js";

const router = useRouter()

const state = reactive({
  collect: []
})


onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/getStarDots", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "Id": store.state.userMsg.Yb_studentid,
    }
  }).then((res) => {
    state.collect = res.data.result
  })
})

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
  <div class="collect">
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
                <span class="infoItem">姚永康</span>
              </div>

              <div class="schoolInfo">
                <el-icon class="icons" style="padding-top: 7px"><Postcard /></el-icon>
                <div class="classinfo">
                  <span>人工智能学院</span>
                  <span>计算机科学与技术1班</span>
                </div>
              </div>
            </div>
          </div>

          <div class="noteNum">
            <el-statistic title="收藏思维导图数量" :value="state.collect.length" class="startNum"/>
          </div>
        </div>
      </el-card>
    </div>

    <el-divider />
    <div class="cards">
      <el-card v-for="(item, index) in state.collect" :key="index" shadow="hover" class="collects">
        <template #header>
          <div class="card-header">
            <span class="title">{{ item.title }}</span>
          </div>
        </template>
        <div class="card-body">
          <div class="body-msg">
            <span>作者：{{ item.auth }}</span>
            <span>更新时间：{{ formatUpdateTime(item.UpdateTime) }}</span>
            <span>tags：{{ item.tags }}</span>
          </div>
          <div class="preview-btn">
            <el-button type="info" @click="view(item.nodeTreePath, item.dot_id)" round>view</el-button>
          </div>
        </div>
        <template #footer>
          <div class="footer-msg">
            <span>描述：{{ item.description }}</span>
          </div>
        </template>
      </el-card>
    </div>
  </div>

</template>

<style scoped>
.collect{
  display: flex;
  flex-direction: column;
  padding-left: 25px;
  padding-right: 25px;
  height: 100%;
  width: 100%;
}

.info{
  display: flex;
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

.openNum{
  padding-right: 80px;
}

.top{
  padding-bottom: 10px;
}

.cards{
  padding: 0 10px 10px 10px;
  display: grid;
  grid-template-columns: repeat(3, 33.3%);
  grid-template-rows: repeat(2, 50%);
  grid-gap: 13px 10px;
}

.collects{
  height: 310px;
  width: 86%;
}

.card-body{
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}
.body-msg{
  display: flex;
  flex-direction: column;
  padding: 0 0 0 4px;
}

.preview-btn{
  display: flex;
  justify-content: end;
  padding-top: 18px;
}

</style>