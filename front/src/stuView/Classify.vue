<script setup>
import {computed, onMounted, reactive, ref} from "vue"
import axios from "axios";
import store from "@/stores/index.js";


const state = reactive({
  classifyData: [],
  tags: [],
  selectTag: {}
})


const fe3 = () => {
  axios.get('http://127.0.0.1:8080/api/v1/hasTags')
    .then(res => {
      store.commit("setTags", res.data.result)
      state.tags = res.data.result
      state.selectTag = state.tags[0]
    })
}
const fe4 = () => {
  axios.post('http://127.0.0.1:8080/api/v1/classify', {
    "accessToken": store.state.Token,
    "Identity": store.state.userMsg.Yb_identity,
    "Id": store.state.userMsg.Yb_studentid,
    "Data": {
      "classifyId": state.selectTag.TagID,
    }
  }).then((res) => {
      state.classifyData = []
      state.classifyData = res.data.result
    })
}

const changeSelect = (item) => {
  state.classifyData = []
  state.selectTag = item
  fe4()
}

onMounted(async () => {
  await fe3()
  await fe4()
})
</script>

<template>
  <div class="classify">
    <div class="tags">
      <el-button
        v-for="item in state.tags"
        size="large"
        disable-transitions
        @click="changeSelect(item)"
        :class="{'active': state.selectTag.TagID === item.TagID}"
      >{{ item.Name }}</el-button>

    </div>
    <el-divider />
    <div class="dir">
      <div class="tagName">
<!--        其实还可以加上一个小标来显示数量-->
        <el-card style="width: 310px" shadow="never">{{ state.selectTag.Name }}</el-card>
      </div>
      <div class="cards">
        <el-card v-for="item in state.classifyData" :key="index" shadow="hover">
          <template #header>
            <div class="card-header">
              <span class="title">{{ item.TITLE }}</span>
            </div>
          </template>
          <div class="card-body">
            <div class="body-msg">
              <span>作者：{{ item.KsID }}</span>
              <span>更新时间：{{ item.UpdateTime }}</span>
              <span>相关笔记：{{ item.notes }}</span>
              <span>tags：{{item.TAGS}}</span>
            </div>
            <div class="preview-btn">
              <el-button type="info" round>preview</el-button>
            </div>
          </div>
          <template #footer>
            <div class="footer-msg">
              <span>描述：{{ item.DESCRIPTION }}</span>
            </div>
          </template>
        </el-card>

      </div>
    </div>
  </div>
</template>

<style scoped>
.tags{
  border-style: solid ;
  border-width: 1px;
  border-color: #7febff;
  border-radius: 16px;
  display: flex;
  flex-wrap: wrap;
  justify-content: start;
  padding: 0 10px 0 10px;
  margin-left: 9px;
  margin-right: 9px;
  user-select: none;
}

.el-button{
  margin: 11px;
}
.active {
  background-color: #d2e4f6;
}
.el-divider{
  margin-left: 9px;
}

.el-card{
  margin: 0 0 0 20px;
}

.cards{
  margin-top: 30px;
  padding: 0 10px 10px 0;
  display: grid;
  grid-template-columns: repeat(3, 33.3%);
  grid-template-rows: repeat(2, 50%);
  grid-gap: 13px 10px;
}

.cards .el-card{
  height: 360px;
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