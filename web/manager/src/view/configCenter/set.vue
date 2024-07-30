<script>
import router from "@/router";

export default {
  methods:{
    logout(){
      localStorage.clear()
      router.go(0)
    }
  },
  data() {
    return {
      activeName: 'first',
    };
  },
}
</script>

<template>
  <el-container>
    <el-main>
      <el-tabs v-model="activeName" class="demo-tabs">
        <el-tab-pane label="账户设置" name="first">
          <el-form>
            <el-form-item>
              <el-button type="danger" @click="logout" plain>退出登录</el-button>
            </el-form-item>
          </el-form>
        </el-tab-pane>
        <el-tab-pane label="系统设置" name="second"></el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
</template>

<style scoped>

</style>