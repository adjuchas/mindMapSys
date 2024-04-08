<script setup>
import {onBeforeMount, onMounted} from "vue";
import {sendRedirect} from "@/api/ims/index.js";
import store from "@/stores/index.js";
import {useRouter} from "vue-router";
import axios from "axios";
import {APPID, SOURCE, UIDNAME} from "@/api/global.js";

const router = useRouter()
const redirectRecommend = () => {
  router.push({
    path: "/Recommend"
  })
}

onMounted(() => {
  let uid = window.localStorage.getItem(UIDNAME)
  console.log(uid)

  let token = store.state.Token
  if (token == null || token === ""){
    if (uid == null || uid === ""){
      axios({
        headers: {
          'Content-Type': "application/x-www-form-urlencoded",
        },
        url: "https://dgcsxyyiban.cn/api/ims/userDistribute",
        method: "GET",
      }).then((res) => {
        window.localStorage.setItem(UIDNAME, res.data)
        sendRedirect(res.data)
      })
    }else {
      console.log(uid)
      axios({
        headers: {
          'Content-Type': "application/x-www-form-urlencoded",
        },
        url: "https://dgcsxyyiban.cn/api/ims/requestToken",
        method: "POST",
        data: "uid=" + uid + "&source=" + SOURCE +"&appId=" +APPID,
      }).then((res) => {
        store.commit("setToken", res.data.accessToken)
        redirectRecommend()
      })
    }
  }else {
    redirectRecommend()
  }
})






</script>

<template>
</template>

<style scoped>

</style>