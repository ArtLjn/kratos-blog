<template>
  <el-container >
    <el-header class="header"  :style="{background:bgColor}">
      <div style="margin-top: 13px;float:left;" class="to">
        <font-awesome-icon :icon="['fas', 'bars']" @click="drawer = true" style="color:white;font-size:30px;" class="daohang"></font-awesome-icon>
        <span class="title">READ BEANS</span>
      </div>
      <el-drawer v-model="drawer" direction="ltr" 
      :lock-scroll="false"
        size="230" style="background-color:rgb(255, 255, 255,0.4);display:flex;flex-direction:column;"
        >
        <div class="lh">
          <router-link to="/" class="link">
            <font-awesome-icon icon="home" style="color:white;"></font-awesome-icon>
            首页
          </router-link>
         </div>
         <div class="lh">
          <router-link to="/tag" class="link">
            <font-awesome-icon :icon="['fas', 'tag']" /> 标签
          </router-link>
        </div>
        <div class="lh">
        <router-link to="/archives" class="link">
          <font-awesome-icon :icon="['fas', 'briefcase']" /> 归档
        </router-link></div>
       <div class="lh"> 
        <router-link to="/about" class="link">
          <font-awesome-icon :icon="['fas', 'circle-user']" /> 关于我
        </router-link>
        </div>
        <div class="lh"> 
            <router-link to="/message" class="link">
              <font-awesome-icon :icon="['fas', 'comment']" /> 留言
            </router-link>
          </div>
        <div class="lh">
          <router-link to="/photo" class="link">
            <font-awesome-icon :icon="['fas', 'image']" /> 相册
          </router-link>
        </div>
        <div class="lh">
          <router-link to="/friendlink" class="link">
            <font-awesome-icon :icon="['fas', 'address-card']" /> 友链
          </router-link>
        </div>
        <div class="lh" @click="searchBlog">
          <font-awesome-icon :icon="['fas', 'magnifying-glass']" style="color:white;"  />
          <span style="color:white;margin-left:5px;">搜索</span>
        </div>
        <div class="lh">
          <router-link to="#" class="link" @click="WantLogin" v-if="Logined">
            <font-awesome-icon :icon="['fas', 'user']" /> 登录
          </router-link>
        </div>
        <div class="lh">
          <router-link to="#" class="link" @click="logout" v-if="logouted">
            <font-awesome-icon :icon="['fas', 'user']" /> 退出
          </router-link>
        </div>
      </el-drawer>
      <div style="width: auto;margin-top: 20px;float:right;" class="dao">
        <router-link to="/" class="link">
          <font-awesome-icon icon="home" style="color:white;"></font-awesome-icon>
          首页
        </router-link>
        <router-link to="/tag" class="link">
          <font-awesome-icon :icon="['fas', 'tag']" /> 标签
        </router-link>
        <router-link to="/archives" class="link">
          <font-awesome-icon :icon="['fas', 'briefcase']" /> 归档
        </router-link>
        <router-link to="/about" class="link">
          <font-awesome-icon :icon="['fas', 'circle-user']" /> 关于我
        </router-link>
        <router-link to="/message" class="link">
          <font-awesome-icon :icon="['fas', 'comment']" /> 留言
        </router-link>
        <router-link to="/photo" class="link">
          <font-awesome-icon :icon="['fas', 'image']" /> 相册
        </router-link>
        <router-link to="/friendlink" class="link">
          <font-awesome-icon :icon="['fas', 'address-card']" /> 友链
        </router-link>
        <font-awesome-icon :icon="['fas', 'magnifying-glass']" style="color:white;" @click="searchBlog" />
        <router-link to="#" class="link" @click="WantLogin" style="margin-left:20px;" v-if="Logined">
          <font-awesome-icon :icon="['fas', 'user']" /> 登录
        </router-link>
        <router-link to="#" class="link" @click="logout" style="margin-left:20px;" v-if="logouted">
          <font-awesome-icon :icon="['fas', 'user']" /> 退出
        </router-link>
      </div>
    </el-header>
      <el-dialog v-model="isOK" style="min-height:450px;height:auto;width:90%;" class="search"
      :lock-scroll="false" accesskey="F"
      >
        <music class="audio"/>
        <div>
          <font-awesome-icon :icon="['fas', 'magnifying-glass']" style="color:black;font-size:30px;margin-right:10px;" />
          <span style="font-size:30px;">搜索</span>
        </div>
          <div>
            <input class="input" placeholder="输入关键字搜索" v-model="keyword.keyOk">
            <div v-for="blog in recommendBlogs" :key="blog.id" style="margin-top:20px;">
              <router-link :to="`/watch/${blog.id}`" class="link-hh" target="_blank">{{ blog.title }}</router-link>
              <p>{{ blog.preface }}</p>
            </div>
          </div>
      </el-dialog>
      <el-dialog v-model="isLogin" 
      :lock-scroll="false"
      style="border-radius:20px;
      width:400px;"
      >
        <login></login>
      </el-dialog>
    <el-main>
      <router-view></router-view>
    </el-main>
    <el-footer class="footer">
    </el-footer>
    <el-backtop :bottom="150" style="background-color:rgb(83, 75, 75,0);">
      <a href="https://support.qq.com/products/599947" target="_blank">
        <font-awesome-icon :icon="['fas', 'comments']" style="color: #e26a94;" />
      </a>
    </el-backtop>
  
    <el-backtop :bottom="80" class="backUp" >
      <el-icon><span style="font-size:20px;font-weight;bold;color:white;">
        <font-awesome-icon :icon="['fas', 'arrow-up']" /></span></el-icon>
    </el-backtop>
  </el-container>
</template>

<script>
import { onMounted, reactive, ref, watch } from 'vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faHome, fas } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
import music from './util/music.vue';
import login from '../authentication/login.vue'
import { ElMessage } from 'element-plus';
import { useRouter } from 'vue-router';
import { getAllBlog } from '@/api/blogFunc';
library.add(faHome, fas);

export default {
  components: {
    FontAwesomeIcon,
    music,
    login
  },
  setup() {
    const isLogin = ref(false);
    const drawer = ref(false);
    const isOK = ref(false);
    const searchBlog = () => {
      isOK.value = true;
      closeLogin();
    };
    const bgColor = ref('')

    const scrollPosition = () => {
      const position = window.pageYOffset
      if (position >= 80) {
        bgColor.value = 'linear-gradient(to top right,rgba(195, 234, 239, 0.5),rgb(214, 214, 254,0.5))'
      } else {
        bgColor.value = null
      }
    }

    const keyword = reactive({
      keyOk: ''
    });
    const recommendBlogs = ref([]);
    
    const searchKeyword = () => {
      recommendBlogs.value = []; 
      Blogs.value.forEach((blog) => {
        if (blog.title.includes(keyword.keyOk) || blog.content.includes(keyword.keyOk)) {
          recommendBlogs.value.push(blog);
        }
      });
    };

    watch(() => keyword.keyOk, () => {
      searchKeyword();
      if(keyword.keyOk == '' || keyword.keyOk == null) {
        recommendBlogs.value = [];
      }
    });
   
    const Blogs = ref([]);
    getAllBlog().then((res) => {
      Blogs.value = res.List;
    })
    onMounted(() => {
      window.addEventListener('scroll', scrollPosition)
      getAllBlog();
      LoginStatus();
    })
    const WantLogin = () => {
      isLogin.value = !isLogin.value;
    }
    const closeLogin = () => {
      if (isLogin.value) {
        isLogin.value = false;
      }
    };
    const Logined = ref(true);
    const logouted = ref(false);
    const router = useRouter();
    const logout = () => {
      localStorage.removeItem("token")
      localStorage.removeItem("authData")
      logouted.value = false;
      Logined.value = true;
      router.go(0);
      ElMessage.info("账号退出")
    }
    const token = localStorage.getItem("token");
    const LoginStatus = () => {
      if(token != null
       || token != "")
       {
        Logined.value = false;
        logouted.value = true;
       }
       if (token == null) {
        Logined.value = true;
        logouted.value = false;
       }
    }
    return {
      drawer,
      isOK,
      searchBlog,
      bgColor,
      scrollPosition,
      Blogs,
      searchKeyword,
      keyword,
      recommendBlogs,
      isLogin,
      WantLogin,
      closeLogin,
      Logined,
      logouted,
      logout
    };
  }
};
</script>
  <style scoped>
  @import url("../../assets/css/main.css");
  .header {
    position: fixed;
    top: 0;
    left: 0;
    width:calc(100%);
    z-index: 999;
  }
  .scrolled {
    background-color: red;
  }
  .audio{
    display: none;
  }
  .daohang{
    display: none;
    margin-top: 4px;
  }
  .title{
    flex-grow: 1;
    text-align: center;
    font-family: fantasy;
  }
  @media screen and (max-width: 1428px) {
    .dao {
      display: none;
    }
    .daohang{
      display: block;
      margin-left: 10px;
    }
    .drawer{
      display: block;
    }
    .to{
      display: flex;
      width: 100%;
    }
    .header{
      padding: 0;
      margin: 0;
      display: flex;
      justify-content: center;
      align-items: center;
    }
    
  }
  .header:not(.scrolled) {
    /* 滚动到顶部时的导航栏样式，移除背景颜色 */
    background-color: transparent;
  }
  .backUp{
    background:linear-gradient(to top right,rgb(34, 82, 171),rgb(67, 211, 208));
    display: flex;
    z-index: 10;
 }
 .input{
  outline: none;
  border-left: none;
  border-right: none;
  border-block-start: none;
  height: 60px;
  width: 100%;
  border-color: green;
  font-size: 20px;
  text-indent: 3px;
 }
 .lh{
  margin-top: 26px;
 }
 .link-hh{
  color: #42b983;
  font-size: 20px;
  text-decoration: none;
 }
  </style>
  