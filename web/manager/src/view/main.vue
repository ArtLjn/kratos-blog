<template>
  <div class="common-layout">
    <el-container class="box">
      <!-- 左侧菜单栏 -->
      <el-aside :width="isCollapsed ? '64px' : '240px'" class="el-aside">
        <div class="logoBox" v-if="!isCollapsed"  @click="toggleCollapse">Blog</div>
        <div class="logoX"> <el-icon v-if="isCollapsed" @click="toggleCollapse"><Grid /></el-icon></div>
        <el-menu
            active-text-color="#ffd04b"
            background-color="#30363c"
            class="el-menu-vertical-demo"
            default-active="2"
            text-color="#fff"
            :router="true"
            :collapse="isCollapsed"
            :collapse-transition="false"
             @collapse-change="handleCollapseChange"
        >

        <template v-for="item in asideMenu" :key="item.title">

            <!-- 两级菜单 -->
            <template v-if="item.subs">
              <el-sub-menu :index="item.title">
                <!-- 一级菜单标题 -->
                <template #title>
                  <el-icon><component :is="item.icon" /></el-icon>
                  <span>{{ item.title }}</span>
                </template>
                <!-- 二级菜单标题 -->
                <template v-for="subItem in item.subs" :key="subItem.index">
                  <el-menu-item
                      :index="subItem.index"
                      class="is-submenu"
                      @click="() => handleMenuItem(subItem)"
                  >
                    <el-icon><component :is="subItem.icon" /></el-icon>
                    <span>{{ subItem.title }}</span>
                  </el-menu-item>
                </template>

              </el-sub-menu>
            </template>

            <!-- 一级菜单 -->
            <template v-else>
              <el-menu-item
                  :index="item.index"
                  @click="() => handleMenuItem(item)"
              >
                <el-icon><component :is="item.icon" /></el-icon>
                <span>{{ item.title }}</span>
              </el-menu-item>
            </template>
          </template>
        </el-menu>

      </el-aside>

      <el-container>
        <!-- 主体模块：标签页 + 当前路由内容 -->
        <el-main class="el-main">

            <el-tabs
              type="border-card"
              v-model="activeTabName"
              class="demo-tabs"
              @tab-remove="handleRemove"
              @tab-click="handleSwitchRoute"
          >
            <el-tab-pane
                v-for="item in editableTabs"
                :key="item.index"
                :label="item.title"
                :name="item.index"
                :closable="handleIsClose(item)"
            >
              <router-view />
            </el-tab-pane>
          </el-tabs>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script>
import {
  Document,
  Setting,
  ArrowLeft,
  ArrowRight,
  Edit,List,PriceTag,PictureFilled,
  Promotion,
  Grid,
  HomeFilled,
  User
} from "@element-plus/icons-vue";
import router from "@/router";

export default {
  name: "MainLayout",
  data() {
    return {
      //当前选项卡
      activeTabName: "/main/blog",
      //需要显示的标签数组
      editableTabs: [
        {
          title: "首页",
          index: "/main/blog",
        },
      ],
      //左侧菜单选项配置
      asideMenu: [
        {
          title: "首页",
          icon: "HomeFilled",
          index: "/main/blog",
        },
        {
          title: "文章",
          icon: "Document",
          subs: [
            {
              title: "文章列表",
              index: "/main/blog",
              icon: "List",
            },
            {
              title: "发布文章",
              index: "/main/edit",
              icon: "Edit",
            },
            {
              title: "标签管理",
              index: "/main/tagManager",
              icon: "PriceTag",
            },
            {
              title: "相册管理",
              index: "/main/photo",
              icon: "PictureFilled",
            },
            {
              title: "友链管理",
              index: "/main/friendManager",
              icon: "Promotion",
            }
          ],
        },
        {
          title: "用户管理",
          icon: "User",
          index: "/main/user",
        },
        {
          title: "配置中心",
          icon: "Setting",
          index: "/main/setting",
        },
      ],
      isCollapsed: false,
    };
  },
  components: {
    Document,
    Setting,
    ArrowLeft,
    ArrowRight,
    Edit,
    List,
    PriceTag,
    PictureFilled,
    Promotion,
    Grid,
    HomeFilled,
    User
  },
  methods: {
    toggleCollapse() {
      this.isCollapsed = !this.isCollapsed;
    },
    handleCollapseChange(collapsed) {
      const menu = document.querySelector('.el-menu-vertical-demo');
      if (collapsed) {
        menu.classList.add('is-collapsed');
      } else {
        menu.classList.remove('is-collapsed');
      }
    },
    handleIsClose(item) {
      return item.index !== "/main/blog";
    },
    handleMenuItem(obj) {
      this.activeTabName = obj.index;
      let result = this.editableTabs.findIndex((item) => {
        return item.index === obj.index;
      });
      if (result !== -1) {
        return;
      }
      this.editableTabs.push(obj);
    },
    handleSwitchRoute(tabsPaneContext) {
      let tabPaneName = tabsPaneContext.paneName;
      if (tabPaneName === 0) {
        tabPaneName = "";
      }
      router.push(tabPaneName);
    },
    handleRemove(tabPaneName) {
      let tempArr = this.editableTabs;
      let eleIndex = this.editableTabs.findIndex((item) => {
        return item.index === tabPaneName;
      });
      let routeIndex;
      for (let p in tempArr) {
        if (tempArr[p].index === tabPaneName) {
          routeIndex = tempArr[p - 1].index;
        }
      }
      this.activeTabName = routeIndex;
      router.push(routeIndex);
      this.editableTabs.splice(eleIndex, 1);
    },
  },
};
</script>

<style scoped>
.logoBox {
  position: absolute;
  top: 18px;
  left: 30px;
  font-size: 24px;
  color: #fff;
  display: flex;
}
.logoX{
  color: white;
  font-size: 24px;
  margin-left: 21px;
}
.box {
  width: 100vw;
  height: 100vh;
}

.collapse-button {
  position: absolute;
  top: 10px;
  left: 10px;
  z-index: 10;
}

.el-aside {
  background-color: #30363c;
  padding-top: 58px;
  transition: width 0.2s;
}

.el-main {
  background-color: #e9eef3;
  transition: padding-left 0.2s;
}

/* 默认状态下，一级菜单和二级菜单的缩进 */
.el-menu-item,
.el-sub-menu__title {
  padding-left: 20px !important;
}

/* 针对二级菜单项的缩进 */
.el-menu-item.is-submenu {
  padding-left: 40px !important;
}

/* 侧边栏收缩时的缩进调整 */
.is-collapsed .el-menu-item,
.is-collapsed .el-sub-menu__title {
  padding-left: 8px !important;
}

/* 针对收缩状态下二级菜单的缩进 */
.is-collapsed .el-menu-item.is-submenu {
  padding-left: 30px !important;
}

</style>
