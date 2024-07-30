<template>
  <div class="container">
    <el-container class="el-container">
      <el-aside style="width:180px" class="aside">
        <ASide/>
      </el-aside>
      <el-container>
        <el-main class="main">
          <router-view></router-view>
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
import ASide from "@/view/components/ASide.vue";
export default{
    components:{
      ASide
    },
    setup() {
      const drawer = ref(false);
      return{
        drawer,
      }
    }
}
</script>
<style scoped>
*{
  margin: 0;
  padding: 0;
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

 .main{
  background-color: #f5f5f5;
    height: auto;
    min-height: 100vh;
    width: calc(100% - 180px);
    margin-left: 180px;
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
</style>