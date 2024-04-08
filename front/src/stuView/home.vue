<script setup>
import avatarJpg from '@/assets/img/avatar.jpg'
import {computed, onBeforeMount, onMounted, onUpdated, reactive, ref} from "vue";
import {useStore} from "vuex";
import axios from "axios";

const store = useStore()
const state = reactive({
  title: "求精知识点分享平台",
  menu: [],
  popovers: [],
  isCollapse: false,
  userMsg: {},
  history: []
})

const menuItems = [
  { index: '/Recommend', icon: 'Share', label: '推荐' },
  { index: '/Classify', icon: 'Menu', label: '分类' },
  { index: '/Draft', icon: 'EditPen', label: '草稿箱子' },
  { index: '/Uploaded', icon: 'UploadFilled', label: '已上传箱子' },
  { index: '/Note', icon: 'Notebook', label: '笔记' },
  { index: '/Collect', icon: 'Star', label: '收藏' }
]

const teaMenuItems = [
  { index: '/Recommend', icon: 'Share', label: '推荐' },
  { index: '/Classify', icon: 'Menu', label: '分类' },
  { index: '/MyDot', icon: 'EditPen', label: '我的稿件' },
  { index: '/Audit', icon: 'Finished', label: '审核' },
  { index: '/ClassifyManage', icon: 'Wallet', label: '分类管理' }
]

const popovers = [
  { id: 'history', icon: 'Clock', label: '历史' },
  { id: 'msg', icon: 'ChatLineSquare', label: '消息' }
]


onMounted(() => {
  axios.post('http://127.0.0.1:8080/api/v1/info', {
    accessToken: store.state.Token
  }).then((res) => {
      state.userMsg = res.data.info
      store.commit("setUserMsg", res.data.info)
      state.popovers = popovers
      if (state.userMsg.Yb_identity !== "学生"){
        state.menu = teaMenuItems
        state.title = "求精知识点分享平台 [ 教师端 ]"
        store.commit("setTeaId", res.data.info.Yb_employid)
      }else {
        state.menu = menuItems
        axios.post('http://127.0.0.1:8080/api/v1/stu/history', {
          accessToken: store.state.Token
        }).then((res) => {
          state.history = res.data.history
      })}
    })

})

</script>

<template>
  <div class="home">
    <el-container>
      <el-header>
        <div class="header">
          <div class="left">
            <router-link :to="{path : '/Recommend'}" class="leftLink">
              <div class="logo">
                <img src="../assets/img/avatar.jpg" alt="">
              </div>
              <div class="sysName">{{ state.title }}</div>

            </router-link>
          </div>

          <div class="right">
            <el-popover v-for="item in popovers" :key="item.index">
              <template #reference>
                <div :id="item.id" class="headerLink">
                  <el-icon class="icons" style="margin-bottom: 4px">
                    <component :is="item.icon"/>
                  </el-icon>
                  <span style="font-size: 13px">{{ item.label }}</span>
                </div>
              </template>
            </el-popover>


            <div id="info">
              <div id="avatar">
                <el-avatar :size="50" :src="avatarJpg"/>
              </div>
              <div id="userName" >
                <span>{{ state.userMsg.Yb_realname }}</span>
              </div>
            </div>
          </div>
        </div>
      </el-header>


      <el-main>
        <el-container>
          <el-aside width="auto">
            <el-menu
              :router="true"
              :collapse="state.isCollapse"
              :collapse-transition="true"
              class="el-menu-vertical"
              :default-active="$route.path"
            >
              <template v-for="item in state.menu">
                <el-menu-item :index="item.index">
                  <el-icon>
                    <component :is="item.icon"/>
                  </el-icon>
                  <span>{{ item.label }}</span>
                </el-menu-item>
              </template>

              <el-button class="aside_control" @click="() => state.isCollapse=!state.isCollapse">
                <div v-if="state.isCollapse">
                  展开<el-icon><ArrowRightBold /></el-icon>
                </div>
                <div v-else>
                  收起<el-icon><ArrowLeftBold /></el-icon>
                </div>
              </el-button>

            </el-menu>
          </el-aside>
          <el-main>
            <router-view/>
          </el-main>
        </el-container>
      </el-main>

    </el-container>
  </div>
</template>

<style scoped>
.home{
  height: 100%;
  display: flex;
  flex-direction: column;
}

.header{
  user-select: none;
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 64px;
  padding: 0;
  background-color: #fff;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.12);
}


.leftLink{
  display: flex;
  align-items: center;
}

.logo img{
  width: auto;
  height: 60px;
  margin-left: 16px;
  margin-right: 25px;
}

.sysName{
  font-size: 20px;
  color: #333;
  font-weight: 500;
}


.right{
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.headerLink{
  margin-left: 20px;
  margin-right: 10px;
}

#info{
  margin-left: 15px;
  margin-right: 10px;
  display: flex;
  align-items: center;
}

#userName{
  margin-left: 15px;
  margin-right: 20px;
}

.icons{
  transform: scale(1.5);
}

.el-main{
  padding: 10px 0 0 0;
}

.el-menu-vertical{
  height: 100vh;
}

.el-menu-vertical:not(.el-menu--collapse){
  width: 200px;
}

.aside_control{
  display: block;
  position: fixed;
  bottom: 0;
  width: inherit;
  height: 40px;
}

button.el-button {
  transition: all 0.03s ease;
}

.headerLink{
  display: flex;
  flex-direction: column;
  justify-content: space-around;
  align-items: center;
  margin-top: 10px;
}

</style>