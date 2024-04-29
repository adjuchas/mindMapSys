import { createApp } from 'vue'
import '@/assets/css/main.css'
import App from './App.vue'
import router from "./router/index.js";
import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'
import {zhCn} from "element-plus/es/locale/index";
import * as ElementIcons from '@element-plus/icons-vue'
import store from "@/stores"
import mavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'

const app = createApp(App)

app.use(ElementPlus, {
  locale: zhCn,
})
app.use(router)
app.use(store)
app.use(mavonEditor)
for(const name in ElementIcons){
  app.component(name, ElementIcons[name])
}

app.mount('#app')
