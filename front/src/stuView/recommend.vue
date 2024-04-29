<script setup>
import {onMounted, reactive} from "vue";
import axios from "axios";
import stores from "@/stores/index.js";
import {useRouter} from "vue-router";
import {formatUpdateTime} from "@/utils/formatTime.js";

const router = useRouter()

const state = reactive({
  cardData: [],
})

onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/recommend", {
    "accessToken": stores.state.Token
  }).then((res) => {
    state.cardData = res.data.result
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
  <div class="recommend" >
    <el-scrollbar height="calc(100vh - 76px )">
    <div class="cards">
      <el-card v-for="(item, index) in state.cardData" :key="index" shadow="hover" >
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
    </el-scrollbar>
  </div>
</template>

<style scoped>
.recommend{
  height: 100%;
}
.cards{
  padding: 0 10px 10px 10px;
  display: grid;
  grid-template-columns: repeat(4, 25%);
  grid-gap: 13px 10px;
  height: 100%
}

.el-card{
  height: 310px;
  width: 86%;
}

.card-body{
  display: flex;
  flex-direction: column;
  justify-content: space-between;
}

.preview-btn{
  display: flex;
  justify-content: end;
  padding-top: 18px;
}

.body-msg{
  display: flex;
  flex-direction: column;
  padding: 0 0 0 4px;
}

</style>