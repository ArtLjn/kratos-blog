<script>
export default {
  name:"ASide",
}
</script>
<template>
  <el-container class="admin-page">
    <el-aside class="aside">
      <el-menu class="menu" text-color="#fff" background-color="#282C34" router :default-active="$route.path" >
        <el-menu-item class="menu-item"
                      style="height: 78px;font-weight: bold;margin-left: 10px;">
          <span style="font-size:20px;line-height: 55px;">
           博客系统
          </span>
        </el-menu-item>
        <el-menu-item  style="display: flex; justify-content:space-between">
          <div class="left">
            <i class="el-icon-house left-icon" style="color: #fff;"></i>
            <span style="font-size:15px">
            <router-link class="link" to="/main/blog">
              文章管理
            </router-link>
           </span>
          </div>
        </el-menu-item>
        <el-menu-item  style="display: flex; justify-content:space-between">
          <div class="left">
            <i class="el-icon-s-management left-icon" style="color: #fff;"></i>
            <span style="font-size:15px">
              <router-link class="link" to="/main/edit">发布文章</router-link>
            </span>
          </div>
        </el-menu-item>
        <el-menu-item style="display: flex; justify-content:space-between">
          <div class="left">
            <i class="el-icon-folder left-icon" style="color: #fff;"></i>
            <span style="font-size:15px">
              <router-link class="link" to="/main/photo">
              相册管理</router-link>
            </span>
          </div>
        </el-menu-item>
        <el-menu-item style="display: flex; justify-content:space-between">
          <div class="left">
            <i class="el-icon-s-promotion left-icon" style="color: #fff"></i>
            <span style="font-size:15px">
              <router-link class="link" to="/main/tagManager">
              标签管理</router-link>
            </span>
          </div>
        </el-menu-item>
        <el-menu-item  style="display: flex; justify-content:space-between">
          <div class="left">
            <i class="el-icon-s-flag left-icon" style="color: #fff"></i>
            <span style="font-size:15px">
              <router-link class="link" to="/main/friendManager">
                友链管理</router-link>
            </span>
          </div>
        </el-menu-item>
        <el-menu-item  style="display: flex; justify-content:space-between">
          <div class="left">
            <i class="el-icon-s-flag left-icon" style="color: #fff"></i>
            <span style="font-size:15px">
                <router-link class="link" to="/main/setting">配置中心</router-link>
            </span>
          </div>
        </el-menu-item>
      </el-menu>
    </el-aside>
  </el-container>
</template>

<style scoped>
.admin-page {
  height: 100vh;

  .aside {
    .menu {
      height: 100%;
    }
  }}
.link{
  color: white;
  margin-left: 10px;
  font-size: 15px;
}
</style>