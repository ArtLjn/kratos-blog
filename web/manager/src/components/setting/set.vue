<script>
import router from "@/router";

export default {
  methods:{
    logout(){
      localStorage.clear()
      router.go(0)
    }
  }
}
</script>

<template>
  <el-form>
    <el-form-item>
      <el-button type="danger" @click="logout" plain>退出登录</el-button>
    </el-form-item>
  </el-form>
</template>

<style scoped>

</style>