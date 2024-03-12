<template>
  <div class="container">
    <el-container class="el-container">
      <el-aside width="200px" class="aside">
        <div style="margin-top:80px;">
          <div class="ddd">
            <router-link class="link" to="/main/blog">
              <font-awesome-icon :icon="['fas', 'file']" /> 笔记</router-link>
          </div>
          <div class="ddd">
            <router-link class="link" to="/main/edit">
              <font-awesome-icon :icon="['fas', 'pen-nib']" /> 发布</router-link>
          </div>
          <div class="ddd">
            <router-link class="link" to="/main/photo">
              <font-awesome-icon :icon="['fas', 'image']" /> 相册</router-link>
          </div>
          <div class="ddd">
            <router-link class="link" to="/main/tagManager">
              <font-awesome-icon :icon="['fas', 'tag']" /> 标签</router-link>
          </div>
          <div class="ddd">
            <router-link class="link" to="/main/friendManager">
              <font-awesome-icon :icon="['fas', 'address-card']" /> 友链</router-link>
          </div>
        </div>
      </el-aside>
      <el-container>
        <el-header class="header">
          <div class="logout">
            <router-link to="/login" style="color:black;" class="link" @click="logout">退出</router-link>
          </div>
        </el-header>
        <el-main class="main">
          <el-card class="card-ff">
            <router-view></router-view>
          </el-card>
        </el-main>
      </el-container>
        <el-backtop :bottom="80" class="backUp" >
          <el-icon><span style="color:white;font-size:10px;font-weight:bold;"></span></el-icon>
        </el-backtop>
    </el-container>
  </div>
</template>
<script>
import { ref } from 'vue';
export default{
    name:"main",
    setup() {
      const drawer = ref(false);
      const logout = () => {
        localStorage.clear()
      }
      return{
        drawer,
        logout
      }
    }
}
</script>
<style scoped>
*{
  margin: 0;
  padding: 0;
}
.header{
  width: 100%;
  padding-top: 15px;
  background-color: #334157;
  border-left: 2px solid #333;
  position: fixed;
  z-index: 999;
}
.logout{
  display: flex;

}

.link{
  color: white;
  margin-left: 10px;
  font-size: 15px;
}
.container{
  height: 100%;
  width:100%;
  display: flex;
}
.el-container{
  height: 100%;
  width: 100%;
  flex: 1;
}
.daohang{
  display: none;
  float: left;
  color:rgb(20, 10, 86);
  font-size:25px;
}

 .main{
  background-color: #f5f5f5;
  height: auto;
  min-height: 100vh;
  margin-top: 60px;
  width: calc(100% - 180px);
  margin-left: 180px;
 }
 .card-ff{
    min-height: 90vh;
    height: auto;
 }
 .backUp{
  background:linear-gradient(to top right,rgb(221, 84, 106),rgb(67, 211, 208));
  display: flex;
  z-index: 10;
}
.aside{
  width: 180px;
  min-height: 100vh;
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: #334157;
  position: fixed;
}
.ddd{
  margin-bottom: 50px;
}
</style>