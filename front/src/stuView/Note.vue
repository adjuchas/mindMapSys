<script setup>
import avatarJpg from "@/assets/img/avatar.jpg";
import {ref} from "vue";
const msg1 = ref("Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点")
const states = [
  {
    state: 0,
    value: "私有"
  },
  {
    state: 1,
    value: "公开"
  }
]
</script>



<template>
  <div class="note">
    <div class="top">
      <el-card shadow="never">
        <div class="info">
          <div class="info-left">
            <div class="pic">
              <el-avatar :size="100" :src="avatarJpg"/>
            </div>
            <div class="usrInfo">
              <div class="name">
                <el-icon class="icons"><User /></el-icon>
                <span class="infoItem">姚永康</span>
              </div>

              <div class="schoolInfo">
                <el-icon class="icons" style="padding-top: 7px"><Postcard /></el-icon>
                <div class="classinfo">
                  <span>人工智能学院</span>
                  <span>计算机科学与技术1班</span>
                </div>
              </div>
            </div>
          </div>

          <div class="noteNum">
            <el-statistic title="公开笔记数量" :value="12" class="openNum"/>
            <el-statistic title="编写笔记数量" :value="13" class="startNum"/>
          </div>
        </div>
      </el-card>
    </div>
    <el-divider />

    <div class="notes">
      <div class="noteInfo">
        <el-card shadow="hover">
          <template #header>
            <el-row class="authMsg">
              <el-col class="AuthMsgItem" :span="18">
                <div class="authName">
                  <el-icon class="icons2"><User /></el-icon>
                  <span class="authSpan">笔记作者：姚永康</span>
                </div>
                <div class="createTime">
                  <el-icon class="icons2"><Calendar /></el-icon>
                  <span class="authSpan">创建日期：2020-10-44</span>
                </div>
              </el-col>

              <div class="AuthMsgItem">
                <div class="authName">
                  <el-icon class="icons2"><User /></el-icon>
                  <span class="authSpan">思维导图作者：姚永康</span>
                </div>
                <div class="createTime">
                  <el-icon class="icons2"><Calendar /></el-icon>
                  <span class="authSpan">更新日期：2020-10-44</span>
                </div>
              </div>
            </el-row>
          </template>


          <el-row class="noteBody" :gutter="20">
            <el-col :span="18">
              <div class="title">
                <el-text size="large">
                  关于xxxx的笔记
                </el-text>
              </div>
              <div class="description">
                <el-text size="large">描述：</el-text>
                <el-text truncated>
                  Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点
                  对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的net/http足够简单，性能也非常不错
                  借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范
                </el-text>
              </div>
              <div class="body">
                <el-text>正文：</el-text>
                <el-text truncated size="small">
                  {{msg1}}
                </el-text>
              </div>
            </el-col>
            <el-col class="dot" :span="3">
              <div class="dotTitle"></div>
              <div class="preview">
                <el-image :src="avatarJpg" fit="fit"/>
              </div>
              <div class="tags"></div>
            </el-col>
          </el-row>

        </el-card>

      </div>
    </div>
  </div>
</template>

<style scoped>
.note{
  display: flex;
  flex-direction: column;
  padding-left: 25px;
  padding-right: 25px;
  height: 100%;
  width: 100%;
}
.body{
  padding-top: 6px;
}

.description{
  padding-top: 3px;
}
.info{
  display: flex;
  flex-direction: row;
  justify-content: start;
  align-items: center;
}

.info-left{
  display: flex;
  padding-right: 100px;
}

.pic{
  padding-right: 23px;
}

.usrInfo{
  display: flex;
  flex-direction: column;
  align-items: start;
  justify-content: center;
}

.name{
  padding-bottom: 19px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.icons{
  transform: scale(1.5);
}

.infoItem{
  padding-left: 13px;
  font-size: 20px;
}

.schoolInfo{
  display: flex;
  justify-content: center;
  align-items: start;
}

.classinfo{
  font-size: 20px;
  display: flex;
  padding-left: 13px;
  flex-direction: column;
  justify-content: start;
}

.noteNum{
  display: flex;
  padding-left: 50px;
}

.openNum{
  padding-right: 80px;
}

.top{
  padding-bottom: 10px;
}

.icons2{
  transform: scale(1.2);
  padding-top: 4px;
}

.authMsg{
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.AuthMsgItem{
  display: flex;
  flex-direction: column;
  align-items: start;
  justify-content: space-between;
}

.authName{
  display: flex;
  justify-content: center;
}

.authSpan{
  padding-left: 7px;
}

.noteBody{
  display: flex;
  justify-content: space-between;
}

</style>