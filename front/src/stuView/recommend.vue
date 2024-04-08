<script setup>
import {onMounted, reactive} from "vue";
import axios from "axios";
import stores from "@/stores/index.js";

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

</script>

<template>
  <div class="recommend">
    <div class="cards">

      <el-card v-for="(item, index) in state.cardData" :key="index" shadow="hover">
        <template #header>
          <div class="card-header">
            <span class="title">{{ item.title }}</span>
          </div>
        </template>
        <div class="card-body">
          <div class="body-msg">
            <span>作者：{{ item.author }}</span>
            <span>更新时间：{{ item.UpdateTime }}</span>
            <span>相关笔记：{{ item.relatedNotes }}</span>
            <span>tags：{{ item.tags }}</span>
          </div>
          <div class="preview-btn">
            <el-button type="info" round>preview</el-button>
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
.recommend{
  position: relative;
}
.cards{
  padding: 0 10px 10px 10px;
  display: grid;
  grid-template-columns: repeat(3, 33.3%);
  grid-template-rows: repeat(2, 50%);
  grid-gap: 13px 10px;
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

.limit-btn{
  padding: 10px;
  display: block;
  position: absolute;
  bottom: 0;
  right: 0;
}
</style>