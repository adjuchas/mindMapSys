<script setup>
import {nextTick, onMounted, reactive, ref} from "vue"
import axios from "axios";
import store from "@/stores/index.js";
import {useRouter} from "vue-router";
import {formatUpdateTime} from "@/utils/formatTime.js";

const router = useRouter()
const scrollbar = ref()
const state = reactive({
  classifyData: [],
  tags: [],
  selectTag: {},
  scrollbarHeight: "auto"
})


const fe3 = () => {
  axios.get('http://127.0.0.1:8080/api/v1/hasTags')
    .then(res => {
      store.commit("setTags", res.data.result)
      state.tags = res.data.result
      state.selectTag = state.tags[0]
      fe4()
    })
}
const fe4 = () => {
  axios.post('http://127.0.0.1:8080/api/v1/classify', {
    "accessToken": store.state.Token,
    "Identity": store.state.userMsg.Yb_identity,
    "Id": store.state.userMsg.Yb_studentid,
    "Data": {
      "classifyId": state.selectTag.TagID.toString(),
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

const getDynamicHeight = () => {
  const myDiv = document.getElementsByClassName('tags')[0]
  const myDiv2 = document.getElementsByClassName('tagName')[0]
  const divHeight = myDiv.offsetHeight + myDiv2.offsetHeight;
  return `calc(100vh - 76px - 140px - ${divHeight}px)`;
}


onMounted(() => {
  fe3()
  state.scrollbarHeight = getDynamicHeight();
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
        <el-badge :value="state.classifyData.length" class="badge">
          <el-card style="width: 310px" shadow="never">{{ state.selectTag.Name }}</el-card>
        </el-badge>
      </div>
      <el-scrollbar ref="scrollbar" :height="state.scrollbarHeight">

        <div class="cards">
          <el-card v-for="item in state.classifyData" :key="index" shadow="hover">
            <template #header>
              <div class="card-header">
                <span class="title">{{ item.TITLE }}</span>
              </div>
            </template>
            <div class="card-body">
              <div class="body-msg">
                <span>作者：{{ item.auth }}</span>
                <span>更新时间：{{ formatUpdateTime(item.UpdateTime) }}</span>
                <span>tags：{{item.TAGS}}</span>
              </div>
              <div class="preview-btn">
                <el-button type="info" @click="view(item.nodeTreePath, item.DotID)" round>view</el-button>
              </div>
            </div>
            <template #footer>
              <div class="footer-msg">
                <span>描述：{{ item.DESCRIPTION }}</span>
              </div>
            </template>
          </el-card>
        </div>
      </el-scrollbar>
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
  margin: 0 7px 0 20px;
}

.cards{
  margin-top: 30px;
  padding: 0 10px 10px 0;
  display: grid;
  grid-template-columns: repeat(4, 25%);
  grid-gap: 13px 10px;
}

.cards .el-card{
  height: 320px;
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