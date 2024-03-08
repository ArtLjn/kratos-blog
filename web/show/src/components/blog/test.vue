<template>  
  <!-- <el-button @click="startCountdown" :disabled="dis">剩余{{countdown}}</el-button>   -->
  <el-form>
    <el-input v-model="data.username"></el-input>
    <el-input v-model="data.password"></el-input>
    <el-button>提交</el-button>
  </el-form>
</template>  
  
<!-- <script>  
import axios from 'axios';
let timer;

export default {  
  name: "test",  
  data() {
    return {
      countdown:5,
      dis:false,
      data:{
        username:"",
        password:""
      }
    }
  },
  methods:{
    startCountdown() {
      this.dis = true;
      if (timer) {
        clearInterval(timer);
      }
      this.countdown = 5;
      timer = setInterval(() => {
        this.countdown --;
        if (this.countdown === 0) {
          clearInterval(timer);
          this.countdown = 5;
          this.dis = false;
        }
      },1000)
    }
  }
}  
</script> -->
<script>
export default {
  name:"test",
  data() {
    return{
      data:{
        username:"",
        password:""
      }
    }
  },
  methods:{
  
  }
}
</script>
