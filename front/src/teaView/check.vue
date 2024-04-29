<script setup>
import {useRoute, useRouter} from "vue-router";
import {marked} from "marked";
import {markedHighlight} from "marked-highlight";
import hljs from "highlight.js";
import {computed, onMounted, reactive, ref} from "vue";
import axios from "axios";
import store from "@/stores/index.js";
import jsMind from 'jsmind'
import 'jsmind/style/jsmind.css'
import {debounce} from "lodash-es";
// 注意引入样式，你可以前往 node_module 下查看更多的样式主题
import 'highlight.js/styles/base16/darcula.css'
const route = useRoute();
const router = useRouter()
marked.use(markedHighlight({
  langPrefix: 'hljs language-',
  highlight(code, lang) {
    const language = hljs.getLanguage(lang) ? lang : 'shell'
    return hljs.highlight(code, { language }).value
  }
}))

const output = computed(() => marked(state.mdData))


const jsmind_container = ref()
const jsmindOptions = {
  container: 'jsmind_container',
  editable:false,
  log_level: 'debug',
  support_html: true,
  mode: 'full',
  view: {
    engine: 'canvas',
    line_style:'curved',
    hide_scrollbars_when_draggable: true,
    draggable: true,
  },
  layout: {
    pspace:13,
  }
}

const state = reactive({
  centerDialogVisible: false,
  jsmind_data: {},
  jm: null,
  node: {
    id: "",
    parent: {},
    topic: "",
    children: [],
    expanded: false,
    data: {
      "background-color": "",
    }
  },
  mdData: "",
  result: false,
  pass: false,
  msg: ""
})

onMounted(() => {
  const path = route.query.nodePath
  axios.post("http://127.0.0.1:8080/api/v1/getDotBody", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "nodePath": path,
    }
  }).then((res) => {
    state.jsmind_data = JSON.parse(res.data.result)
    let option = Object.assign(jsmindOptions, {
      container: jsmind_container.value
    })
    state.jm = new jsMind(option);
    state.jm.show(state.jsmind_data);
  })
})

const onNodeSelected = () => {
  if (state.jm.get_selected_node() != null){
    state.mdData = ""
    let selectNode = state.jm.get_selected_node()
    state.node.id = selectNode.id
    state.node.children = selectNode.children
    state.node.parent = selectNode.parent
    state.node.topic = selectNode.topic
    state.node.expanded = selectNode.expanded
    state.node.data = selectNode.data
    findMd()
  }else {
    state.node.id = ""
    state.node.data = {}
    state.node.expanded = false
    state.node.topic = ""
  }
}

const findMd = () => {
  const dotId = route.query.dotId
  axios.post("http://127.0.0.1:8080/api/v1/getMd", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "mdPath": dotId + state.node.id + ".md",
    },
  }).then((res) => {
    state.mdData = res.data.result
    state.centerDialogVisible = true
  })
}

const passAudit = () => {
  state.result = true
  state.pass = true
}

const noAudit = () => {
  state.result = true
  state.pass = false
}

const setPass = () => {
  const dotId = route.query.dotId
  axios.post("http://127.0.0.1:8080/api/v1/passAudit", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "draftId": dotId.toString()
    },
  }).then((res) => {
    if (res.data.result){
      router.push({
        path: '/Audit',
      })
    }
  })
}

const setUnPass = () => {
  const dotId = route.query.dotId
  axios.post("http://127.0.0.1:8080/api/v1/unPassAudit", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "draftId": dotId.toString(),
      "msg": state.msg,
      "tea": store.state.userMsg.Yb_studentid,
    },
  }).then((res) => {
    if (res.data.result){
      router.push({
        path: '/Audit',
      })
    }
  })
}

</script>

<template>
  <div class="show">
    <div ref="jsmind_container" class="jsmind-container" v-on:click="onNodeSelected"></div>
    <div class="btns">
      <el-button class="btn" round @click="noAudit">不予通过</el-button>

      <el-button class="btn" round @click="passAudit">审核通过</el-button>
    </div>
  </div>

  <el-dialog
    v-model="state.centerDialogVisible"
    title="Notice"
    destroy-on-close
    :lock-scroll="false"
  >

    <template #header="{ close, titleId, titleClass }">
      <div class="my-header">
        <h4 :id="titleId" :class="titleClass">{{ state.node.topic }}</h4>
      </div>
    </template>

    <div class="output" v-html="output"></div>
  </el-dialog>

  <el-dialog
    v-model="state.result"
    title="Notice"
    destroy-on-close
    align-center
    width="500"
    :lock-scroll="false"
    v-if="state.pass"
  >

    <template #header="{ close, titleId, titleClass }">
      <div class="my-header">
        <h4 :id="titleId" :class="titleClass">是否确认通过</h4>
      </div>
    </template>

    <div class="isPass">
      <el-button type="primary" @click="setPass" round>是</el-button>
      <el-button type="info" @click="state.result = false" round>否</el-button>
    </div>
  </el-dialog>

  <el-dialog
    v-model="state.result"
    title="Notice"
    destroy-on-close
    align-center
    width="500"
    :lock-scroll="false"
    v-else
  >

    <template #header="{ close, titleId, titleClass }">
      <div class="my-header">
        <h4 :id="titleId" :class="titleClass">输入驳回意见：</h4>
        <el-input
          class="unpassInput"
          v-model="state.msg"
          :autosize="{ minRows: 4, maxRows: 7 }"
          maxlength="100"
          show-word-limit
          type="textarea"
          placeholder="请输入驳回修改的意见"
        ></el-input>
      </div>
    </template>

    <div class="isPass">
      <el-button type="primary" @click="setUnPass" round>是</el-button>
      <el-button type="info" @click="state.result = false" round>否</el-button>
    </div>
  </el-dialog>
</template>

<style scoped>
.show{
  display: flex;
  flex-direction: column;
  height: 100%;
  width: calc(100% - 24px);
  justify-content: start;
}

.jsmind-container{
  border-style: solid ;
  border-width: 1px;
  border-color: #7febff;
  border-radius: 16px;
  margin-left: 12px;
  margin-right: 12px;
  width: calc(100%);
  height: calc((100% - 76px)*0.95);
}

.btns{
  display: flex;
  flex-direction: row-reverse;
  margin-top: 4px;
}

.btn{
  margin-left: 5px;
  margin-right: 5px;
}

.isPass{
  display: flex;
  justify-content: end;
}

.unpassInput{
  margin-top: 10px;
}
</style>