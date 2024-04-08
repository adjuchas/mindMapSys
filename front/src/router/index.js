import { createRouter, createWebHistory } from 'vue-router'
import store from "@/stores/index.js";

const Home = () => import('@/stuView/home.vue')
const Recommend = () => import('@/stuView/recommend.vue')
const Classify = () => import('@/stuView/Classify.vue')
const Draft = () => import('@/stuView/Draft.vue')
const Uploaded = () => import('@/stuView/Uploaded.vue')
const Note = () => import('@/stuView/Note.vue')
const Collect = () => import('@/stuView/Collect.vue')
const Tete = () => import('@/stuView/Tete.vue')
const Login = () => import('@/stuView/login.vue')
const MyDot = () => import('@/teaView/myDot.vue')

const router = createRouter({
  history: createWebHistory(),
  base: '/',
  routes: [
    {
      path: '/Login',
      name: 'Login',
      component: Login,
    },
    {
      path: '/',
      redirect:{
        name: 'Login'
      }
    },
    {
      path: '/Home',
      name: 'Home',
      component: Home,
      children: [
        {
          path: '/Recommend',
          name: 'Recommend',
          meta: {
            title: "推荐"
          },
          component: Recommend
        },
        {
          path: '/Classify',
          name: 'Classify',
          meta: {
            title: "分类"
          },
          component: Classify
        },
        {
          path: '/Draft',
          name: 'Draft',
          meta: {
            title: "草稿"
          },
          component: Draft
        },
        {
          path: '/Uploaded',
          name: 'Uploaded',
          meta: {
            title: "上传管理"
          },
          component: Uploaded
        },
        {
          path: '/Note',
          name: 'Note',
          meta: {
            title: "笔记"
          },
          component: Note
        },
        {
          path: '/Collect',
          name: 'Collect',
          meta: {
            title: "推荐"
          },
          component: Collect
        },
        {
          path: '/tete',
          name: "tete",
          meta: {
            title: "编辑"
          },
          component: Tete,
        },
        {
          path: '/MyDot',
          name: 'myDot',
          meta: {
            title: "我的稿件"
          },
          component: MyDot
        },
      ]
    }
  ]
})

router.beforeEach((to, form, next) => {
  document.title = to.meta.title === undefined ? "求精知识点分享平台" : to.meta.title + " | 求精"
  if (store.state.Token === "" && to.name !== "Login" ){
    next({
      path: "/Login"
    })
  } else {
    next()
  }
})

export default router