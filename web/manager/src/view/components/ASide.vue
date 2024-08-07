
<template>
  <el-container class="admin-page">
    <el-aside class="aside">
      <el-menu
          class="menu"
          text-color="#fff"
          background-color="#282C34"
          router
          :default-active="$route.path"
      >
        <el-menu-item class="menu-item" style="height: 78px; font-weight: bold; margin-left: 10px;">
          <span style="font-size: 20px; line-height: 55px;">
            博客系统
          </span>
        </el-menu-item>
        <el-menu-item v-for="item in menuItems" :key="item.label" class="menu-item" style="height: 50px; display: flex; align-items: center; padding-left: 20px;">
          <router-link :to="item.path" class="link">
            <i :class="item.icon" style="color: #fff; margin-right: 10px;"></i>
            <span>{{ item.label }}</span>
          </router-link>
        </el-menu-item>
      </el-menu>
    </el-aside>
  </el-container>
</template>

<script>
export default {
  data() {
    return {
      menuItems: [
        { label: "文章管理", path: "/main/blog", icon: "el-icon-house" },
        { label: "发布文章", path: "/main/edit", icon: "el-icon-s-management" },
        { label: "相册管理", path: "/main/photo", icon: "el-icon-folder" },
        { label: "标签管理", path: "/main/tagManager", icon: "el-icon-s-promotion" },
        { label: "友链管理", path: "/main/friendManager", icon: "el-icon-s-flag" },
        { label: "配置中心", path: "/main/setting", icon: "el-icon-s-flag" },
      ],
    };
  },
};
</script>

<style scoped>
.admin-page {
  height: 100vh;
}

.aside .menu {
  height: 100%;
}

.link {
  color: white;
  display: flex;
  align-items: center;
  width: 100%;
  height: 100%;
  text-decoration: none;
  font-size: 15px;
}

.menu-item:hover .link {
  text-decoration: none;
  color: #409eff;
}
</style>
