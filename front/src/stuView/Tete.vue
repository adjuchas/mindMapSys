<script setup>
import jsMind from 'jsmind'
import 'jsmind/style/jsmind.css'
import 'jsmind/draggable-node'
import {reactive, ref} from "vue";
import { onMounted } from "vue";
import {useRoute} from "vue-router";
import axios from "axios";
import store from "@/stores/index.js";
import 'highlight.js/styles/base16/darcula.css'
import {ElMessage, ElNotification} from "element-plus";


const route = useRoute();

const jsmind_container = ref()
const editor = ref()

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
  path: route.query.draftPath,
  draftId: route.query.draftId,
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
  mdData: "",
  toolbars: {
    imagelink: true, // 图片链接
  },
})

function generateUniqueFourDigitNumber() {
  const currentTime = new Date();
  const milliseconds = currentTime.getMilliseconds();
  const uniqueNumber = (milliseconds % 9000) + 1000;
  return uniqueNumber;
}


onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/stu/getDraftBody", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "draftPath": state.path,
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
  updateJsmindData()
  axios.post("http://127.0.0.1:8080/api/v1/stu/setDraftBody", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "draftPath": state.path,
      "draftId": state.draftId,
    },
    "NodeJson": state.jsmind_data
  }).then((res) => {
    if (res.data.result === true){
      ElMessage({
        message: '保存成功',
        type: 'success',
      })
    }
  })
}

function updateJsmindData() {
  state.jsmind_data = state.jm.get_data("node_tree")
}

const setMd = () => {
  axios.post("http://127.0.0.1:8080/api/v1/setMd", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "mdData": state.mdData,
      "mdPath": state.draftId + state.node.id.toString() + ".md",
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
  axios.post("http://127.0.0.1:8080/api/v1/getMd", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "mdPath": state.draftId + state.node.id + ".md",
    },
  }).then((res) => {
    state.mdData = res.data.result
    state.visible = true
  })
}

const imgAdd = (pos,file)=> {
  let imgData = new FormData()
  imgData.append('file',file)
  axios.post("http://127.0.0.1:8080/api/v1/uploadImg", imgData, {
    headers: {
      'Content-Type': 'multipart/form-data'
    }
  }).then((res) => {
    let replacedFilePath = res.data.result.replace(/\\/g, "/");
    editor.value.$img2Url(pos, `http://127.0.0.1:8080/${replacedFilePath}`)
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
          <el-descriptions-item label="叶节点"><el-tag size="default">{{ state.node.children.length === 0}}</el-tag></el-descriptions-item>
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
        <el-form label-width="auto" label-position="left" class="demo-form-inline" size="large" onsubmit="return false">
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
      <mavon-editor ref="editor" id="editor" :toolbars="state.toolbars" v-model="state.mdData" @imgAdd="imgAdd"></mavon-editor>
    </div>
  </el-drawer>
</template>

<style scoped>
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


.backMain{
  margin-bottom: 0;
}

.divider{
  margin: 13px 0 10px 0;
}

.notSelect{
  height: calc(100% - 120px);
}

.editor{
  height: calc(100%);
  width: calc(100%);
}

#editor{
  height: calc(100%);
  width: calc(100%);
}
</style>