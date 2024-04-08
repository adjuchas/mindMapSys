<script setup>

import avatarJpg from "@/assets/img/avatar.jpg";
</script>

<template>
  <div class="collect">
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
            <el-statistic title="收藏笔记数量" :value="12" class="openNum"/>
            <el-statistic title="收藏思维导图数量" :value="13" class="startNum"/>
          </div>
        </div>
      </el-card>
    </div>

    <el-divider />
    <div class="startInfo">
      <div class="startNote">

      </div>

      <div class="startDot">

      </div>
    </div>
  </div>

</template>

<style scoped>
.collect{
  display: flex;
  flex-direction: column;
  padding-left: 25px;
  padding-right: 25px;
  height: 100%;
  width: 100%;
}

.info{
  display: flex;
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
</style>