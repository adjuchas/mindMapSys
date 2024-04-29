<script setup>
import {useRoute} from "vue-router";
import {computed, onMounted, reactive, ref} from "vue";
import axios from "axios";
import store from "@/stores/index.js";
import jsMind from 'jsmind'
import 'jsmind/style/jsmind.css'
import 'highlight.js/styles/base16/darcula.css'
import {ElMessage} from "element-plus";
const route = useRoute();

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
  path: route.query.nodePath,
  dotId: route.query.dotId,
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
  mdNote: "",
  visible: false,
  isStar: false,
  toolbars: {
    imagelink: true, // 图片链接
  },
  containerHeight: "",
  showToolbar: {
  }
})

onMounted(() => {
  axios.post("http://127.0.0.1:8080/api/v1/getDotBody", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
      "dotId": state.dotId,
      "nodePath": state.path,
    }
  }).then((res) => {
    state.jsmind_data = JSON.parse(res.data.result)
    let option = Object.assign(jsmindOptions, {
      container: jsmind_container.value
    })
    state.jm = new jsMind(option);
    state.jm.show(state.jsmind_data);
  })
  axios.post("http://127.0.0.1:8080/api/v1/checkStar", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
      "dotId": state.dotId
    }
  }).then((res) => {
    state.isStar = res.data.result
  })
  if ( store.state.userMsg.Yb_identity === "学生"){
    state.containerHeight = "calc((100% - 76px)*0.95)"
  }else {
    state.containerHeight = "calc(100% - 76px)"
  }
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
  axios.post("http://127.0.0.1:8080/api/v1/getMd", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "mdPath": state.dotId + state.node.id + ".md",
    },
  }).then((res) => {
    state.mdData = res.data.result
    state.centerDialogVisible = true
  })
}

const setStar = () => {
  axios.post("http://127.0.0.1:8080/api/v1/startDot", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "Id": store.state.userMsg.Yb_studentid,
      "dotId": state.dotId
    },
  }).then((res) => {
    if (res.data.result){
      state.isStar = true
      ElMessage({
        message: '收藏成功',
        type: 'success',
      })
    }
  })
}


const quitStar = () => {
  axios.post("http://127.0.0.1:8080/api/v1/quitStar", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
      "dotId": state.dotId
    },
  }).then((res) => {
    if (res.data.result){
      state.isStar = false
      ElMessage('已取消收藏')
    }
  })
}

const save = ()=> {
  axios.post("http://127.0.0.1:8080/api/v1/stu/setNote", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
      "dotId": state.dotId,
      "mdData": state.mdNote,
    },
  }).then((res) => {
    if (res.data.result === "success"){
      ElMessage({
        message: '保存成功',
        type: 'success',
      })
    }
  })
}

const getMd = () => {
  axios.post("http://127.0.0.1:8080/api/v1/stu/getNote", {
    "accessToken": store.state.Token,
    "Id": store.state.userMsg.Yb_studentid,
    "Identity": store.state.userMsg.Yb_identity,
    "Data": {
      "id": store.state.userMsg.Yb_studentid,
      "dotId": state.dotId,
    },
  }).then((res) => {
    state.mdNote = res.data.result
  })
  state.visible = true
}


</script>

<template>
    <div class="show">
      <div ref="jsmind_container" class="jsmind-container" v-on:click="onNodeSelected" :style="{ height: state.containerHeight }"></div>
      <div class="btns" v-if="store.state.userMsg.Yb_identity === '学生'">
        <el-button class="btn" @click="getMd" round>编辑笔记</el-button>
        <el-button v-if="!state.isStar" class="btn" @click="setStar" round>添加收藏</el-button>
        <el-button v-else class="btn" @click="quitStar" round>取消收藏</el-button>
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
      <mavon-editor id="editor" :toolbars="state.showToolbar" v-model="state.mdData" :subfield="false" :editable="false" defaultOpen="preview" codeStyle="github-dark-dimmed" :toolbarsFlag="false"></mavon-editor>
    </el-dialog>

  <el-drawer size="80%" v-model="state.visible" :show-close="false" :lock-scroll="false">
    <div class="editor">
      <mavon-editor id="editor" :toolbars="state.toolbars" v-model="state.mdNote">
        <template v-slot:right-toolbar-before>
          <div class="btns">
            <el-button title="关闭" type="info" class="btn" @click="state.visible = false">
              <el-icon><CloseBold /></el-icon>
              close
            </el-button>
            <el-button title="保存" type="primary" class="btn" @click="save">
              <el-icon><Upload /></el-icon>
              save
            </el-button>
          </div>
        </template>
      </mavon-editor>
    </div>
  </el-drawer>
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

}

.btns{
  display: flex;
  flex-direction: row-reverse;
  margin-top: 4px;
  align-items: center;
}

.btn{
  margin-left: 10px;
  margin-right: 10px;
}

.editor{
  height: calc(100%)
}

#editor{
  height: calc(100%)
}

</style>