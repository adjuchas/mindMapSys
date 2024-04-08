<script setup>
import jsMind from 'jsmind'
import 'jsmind/style/jsmind.css'
import {computed, reactive, ref} from "vue";
import { onMounted } from "vue";
import {useRoute} from "vue-router";
import axios from "axios";
import store from "@/stores/index.js";
import {debounce} from "lodash-es";
import {marked} from "marked";
import { markedHighlight } from "marked-highlight"
import hljs from 'highlight.js'
// 注意引入样式，你可以前往 node_module 下查看更多的样式主题
import 'highlight.js/styles/base16/darcula.css'
import {ElNotification} from "element-plus";


const route = useRoute();
marked.use(markedHighlight({
  langPrefix: 'hljs language-',
  highlight(code, lang) {
    const language = hljs.getLanguage(lang) ? lang : 'shell'
    return hljs.highlight(code, { language }).value
  }
}))

const jsmind_container = ref()
const jsmindOptions = {
  container: 'jsmind_container',
  editable:true,
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
  jsmind_data: {},
  color1: '#409EFF',
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
  visible: false,
  mdData: {
    input: "# hello"
  }
})

const output = computed(() => marked(state.mdData.input))
const update = debounce((e) => {
  state.mdData.input.value = e.target.value
}, 100)

function generateUniqueFourDigitNumber() {
  const currentTime = new Date();
  const milliseconds = currentTime.getMilliseconds();
  const uniqueNumber = (milliseconds % 9000) + 1000;
  return uniqueNumber;
}


onMounted(() => {
  const path = route.query.draftPath
  axios.post("http://127.0.0.1:8080/api/v1/stu/getDraftBody", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "draftPath": path,
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
    let selectNode = state.jm.get_selected_node()
    state.node.id = selectNode.id
    state.node.children = selectNode.children
    state.node.parent = selectNode.parent
    state.node.topic = selectNode.topic
    state.node.expanded = selectNode.expanded
    state.node.data = selectNode.data
  }else {
    state.node.id = ""
    state.node.data = {}
    state.node.expanded = false
    state.node.topic = ""
  }
}

const changeTopic = (val) => {
  state.jm.update_node(state.node.id, val)
}

const changeExpanded = (val) => {
  if (val) {
    state.jm.expand_node(state.node.id)
  }else {
    state.jm.collapse_node(state.node.id)
  }
}

const changeColor = (val) => {
  state.jm.set_node_color(state.node.id, val)
}

const addNewNode = () => {
  state.jm.add_node(state.node.id, generateUniqueFourDigitNumber(), "请输入内容")
}

const delNode = () => {
  state.jm.remove_node(state.node.id)
}

const delChildrenNode = () => {
  let parId = state.node.parent.id
  let node = {
    id: state.node.id,
    topic: state.node.topic,
    expanded: false,
    data: state.node.data
  }
  state.jm.remove_node(state.node.id)
  state.jm.add_node(parId, node.id, node.topic, node.data)
}


const saveNodeTree = () => {
  const path = route.query.draftPath
  updateJsmindData()
  axios.post("http://127.0.0.1:8080/api/v1/stu/setDraftBody", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "draftPath": path,
    },
    "NodeJson": state.jsmind_data
  }).then((res) => {
    if (res.data.result === true){
      location.reload()
    }
  })
}

function updateJsmindData() {
  // 更新jsmind_data
  state.jsmind_data = state.jm.get_data("node_tree")
}

const setMd = () => {
  const draftId = route.query.draftId
  console.log(draftId)
  console.log(state.mdData)
  axios.post("http://127.0.0.1:8080/api/v1/setMd", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "mdData": state.mdData.input,
      "mdPath": draftId + state.node.id + ".md",
    },
  }).then((res) => {
    if (res.data.result === true){
      console.log(res)
      ElNotification({
        title: '结果：',
        message: 'markdown笔记已保存',
        duration: 1000,
      })
    }
  })
}

const editMd = () => {
  const draftId = route.query.draftId
  axios.post("http://127.0.0.1:8080/api/v1/getMd", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "mdPath": draftId + state.node.id + ".md",
    },
  }).then((res) => {
    state.mdData.input = res.data.result
    state.visible = true
  })
}


</script>

<template>
  <div class="Tete">
    <div ref="jsmind_container" class="jsmind-container" v-on:click="onNodeSelected"></div>
    <div class="setStyle">
      <div class="nodeInfo" v-if="state.node.id !== ''">
        <el-descriptions
          title="节点数据："
          direction="vertical"
          :column="3"
          border
        >
          <el-descriptions-item label="节点ID"><el-tag size="default">{{ state.node.id }}</el-tag></el-descriptions-item>
          <el-descriptions-item label="根节点"><el-tag size="default">{{ state.node.parent == null }}</el-tag></el-descriptions-item>
          <el-descriptions-item label="叶节点"><el-tag size="default">{{ state.node.children.length == null}}</el-tag></el-descriptions-item>
          <el-descriptions-item label="子节点">
            <el-tag size="default">{{ state.node.children.length }}</el-tag>
          </el-descriptions-item>
          <el-descriptions-item label="父节点内容">
            <el-tag size="large">{{ state.node.parent?.topic || "无"}}</el-tag>
          </el-descriptions-item>
        </el-descriptions>
        <el-divider class="divider" />
      </div>
<!--      或许可以用骨架屏来展示,目前的样式写法也太搞笑了吧-->
      <div v-else class="notSelect">请选择节点</div>

      <div class="nodeOperate" v-if="state.node.id !== ''">
        <div class="operaTitle">节点操作：</div>
        <el-form label-width="auto" label-position="left" class="demo-form-inline" size="large">
          <el-form-item label="节点内容：">
            <el-input clearable size="default" v-model="state.node.topic" @input="changeTopic"/>
          </el-form-item>

          <el-form-item label="展开节点：" v-if="state.node.children.length !== 0">
            <el-switch size="default" v-model="state.node.expanded" @change="changeExpanded"/>
          </el-form-item>

          <div class="backCol">
            <el-form-item label="背景色：" class="backMain">
              <el-color-picker size="default" v-model="state.node.data['background-color']" @active-change="changeColor" />
            </el-form-item>
          </div>
        </el-form>
        <el-divider class="divider" />
      </div>

      <div class="nodeBtns" v-if="state.node.id !== ''">
        <el-button class="nodebtn" type="primary" @click="addNewNode" plain>新建节点</el-button>
        <el-button class="nodebtn" type="primary" @click="editMd" plain>md编辑</el-button>
        <el-button class="nodebtn" type="primary" v-if="state.node.parent != null" @click="delChildrenNode" plain>删除子节点</el-button>
        <el-button class="nodebtn" type="primary" v-if="state.node.parent != null" @click="delNode" plain>删除当前节点</el-button>
      </div>
      <el-divider class="divider" />

      <el-button type="primary" @click="saveNodeTree">保存思维导图</el-button>
    </div>
  </div>

  <el-drawer size="80%" v-model="state.visible" :show-close="false" :lock-scroll="false">
    <template #header="{ close, titleId, titleClass }">
      <h1>{{ state.node.topic }}</h1>
      <el-button type="primary" @click="setMd">
        <el-icon><Upload /></el-icon>
        <span>Save</span>
      </el-button>
      <el-button type="danger" @click="close">
        <el-icon><Close /></el-icon>
        <span>Close</span>
      </el-button>
    </template>
    <div class="editor">
      <el-input
        v-model="state.mdData.input"
        placeholder="Please input"
        type="textarea"
        @input="update"
        class="input"
        :autosize="{ minRows: 25, maxRows: 999 }"
      />
      <div class="output" v-html="output"></div>
    </div>
  </el-drawer>


</template>

<style scoped>

@import "v1.css";

.Tete{
  display: flex;
  height: 100%;
  width: 100%;
  justify-content: space-between;
}
.jsmind-container{
  border-style: solid ;
  border-width: 1px;
  border-color: #7febff;
  border-radius: 16px;
  margin-left: 12px;
  margin-right: 12px;
  width: calc(80%);
  height: calc(100% - 76px);
}

.setStyle{
  height: calc(100% - 76px);
  width: calc(18%);
  border-style: solid ;
  border-width: 1px;
  border-color: #7febff;
  border-radius: 16px;
  margin-right: 8px;
}

.setStyle{
  display: flex;
  flex-direction: column;
  justify-content: start;
  padding: 10px 5px 0 5px;
}

.nodeInfo{
  padding-left: 5px;
  margin-right:  5px;
}

.el-divider{
  padding-left: 9px;
  padding-right: 9px;
  width: calc(100%);
}

.operaTitle{
  font-weight: bold;
  margin-bottom: 7px;
}

.nodeBtns{
  display: flex;
  flex-direction: row;
  flex-wrap: wrap;
  justify-content: space-around;
}

.nodebtn{
  width: calc(50% - 8px);
  margin: 0 0 7px 0 ;
}

.backCol{
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.saveNode{
  margin: 0 7px 0 0;
}

.backMain{
  margin-bottom: 0;
}

.divider{
  margin: 13px 0 10px 0;
}

.notSelect{
  height: calc(100% - 120px);
}
</style>