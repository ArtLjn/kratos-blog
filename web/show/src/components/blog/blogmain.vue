<template>
    <div class="body" >
      <el-container class="mains" :style="{ 'background-image': `url(${RandomPhoto.photo})` }">
        <el-header></el-header>
          <div class="ls">
            <span class="main-title">红豆南墙</span>
            <span class="bottom-title poetry">{{poetry.content}}</span>
          </div>
          <div class="util">
            <button class="util-button" @click="utilButton">    
              <font-awesome-icon :icon="['fas', 'angles-down']" style="color: #ffffff;margin-right:5px;" />开始阅读</button>
          </div>
          <div class="bottom-title " style="margin-top: 80px;">
            <div style="margin-right: 30px;">
              <el-tooltip content="QQ联系我:1521505611" placement="top">
                <a href="
                https://qiniu.lllcnm.cn/img/3e7a1966-1df3-4044-8b12-6d41f11c09a0.jpg" style="color:white;" target="_blank">
                <font-awesome-icon :icon="['fab', 'qq']" /></a>
              </el-tooltip>
            </div>
            <div style="margin-right: 30px;">
              <el-tooltip content="微信扫一扫" placement="top">
                <a href="https://qiniu.lllcnm.cn/img/3d7f9074-b831-484b-a274-b5569a758f8f.jpg" style="color:white;" target="_blank">
                <font-awesome-icon :icon="['fab', 'weixin']" /></a>
              </el-tooltip>
            </div>
            <div style="margin-right: 30px;">
              <el-tooltip content="邮件联系我" placement="top">
                <a href="mailto:1521505611@qq.com" style="color:white;" target="_blank">
                <font-awesome-icon :icon="['fas', 'envelope']" /></a>
              </el-tooltip>
            </div>
            <div style="margin-right: 30px;">
              <el-tooltip content="访问github" placement="top">
                <a href="https://github.com/620324" style="color:white;text-direction:none;" target="_blank">
                  <font-awesome-icon :icon="['fab', 'github']" /></a>
              </el-tooltip>
            </div>
          </div>
      </el-container>
        <div class="blog-main">
          <el-card class="blog-referrer">
            <div class="referrer-title"><font-awesome-icon :icon="['fas', 'thumbs-up']" style="padding-top:3.5px;padding-right:3px;" ></font-awesome-icon>推荐文章</div>
              <div class="suggest">
                <div v-for="blog in SuggestList" :key="blog.id" class="blog2">
                  <div class="card2">
                    <router-link :to="`/watch/${blog.id}`" class="top">
                      <img :src="`${blog.photo}`" class="img2" alt="图片怎么丢失了😜" >
                      <div class="blog-t2">{{ blog.title }}</div>
                      <div class="preface2">{{blog.preface}}</div>
                      <button class="button2">阅读更多</button>
                    </router-link>
                  </div>
                </div>
            </div>
          </el-card>
          <div class="blogs-container">
            <div v-for="blog in displayedItems" :key="blog.id" class="blog animate__bounce" style="margin-top: 20px; animation: fadeInOut 2s ease-in-out;">
                  <div class="card">
                    <router-link :to="`/watch/${blog.id}`" class="top">
                          <img v-lazy="`${blog.photo}`" alt="图片怎么丢失了😜" class="img">
                          <div class="blog-t">{{ blog.title }}</div> 
                    </router-link>
                    <div class="content">
                        <span class="preface">{{ blog.preface }}</span>
                        <div style="margin-top:10px;">
                          <font-awesome-icon :icon="['fas', 'calendar-days']" />
                          <span class="create-time">{{ blog.createTime }}</span>
                            <hr style="background-color:aqua;">
                        </div>
                        <div style="margin-bottom:5px;"><span class="bq" style="background-color:rgb(20, 9, 59);color:white;padding:2px;
                          font-size:14px;border-top-left-radius: 5px;border-bottom-left-radius:5px;">
                            <font-awesome-icon :icon="['fas', 'droplet']" /> 标签</span><span style="background-color:rgb(207, 115, 130);
                            color:white;padding:2px;font-size:14px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
                              {{blog.tag}}</span>
                              <span style="float:right;margin-right:4px;" class="saw">
                                <font-awesome-icon :icon="['fas', 'eye']" /> {{blog.visits}}</span>
                        </div>
                    </div>
                </div>
            </div>
          </div> 
          <el-pagination  :total="Blogs.length" @current-change="handleCurrentChange" pager-count="3"
                          layout="prev, pager, next"
          background class="page" >
        </el-pagination>
      </div>
      <newFooter></newFooter>
    </div>
  </template>
  
  <script>
  import { computed, onMounted, reactive, ref } from 'vue';
  import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
  import { faHome, fas } from '@fortawesome/free-solid-svg-icons';
  import { library } from '@fortawesome/fontawesome-svg-core';
  import { fab } from '@fortawesome/free-brands-svg-icons';
  import newFooter from './util/newFooter.vue';
  import {getAllBlog,getShici,getBingPhoto, getAllSuggest} from '../../api/blogFunc'
  import {utilButton} from '../../api/util/util'
  import {SUCCESS_REQUEST} from "@/api/status";
  library.add(faHome, fas, fab);
  export default {
    // eslint-disable-next-line vue/multi-word-component-names
    components: {
      FontAwesomeIcon,
      newFooter
    },
    setup() {
      const itemsPerPage = 9;
      const currentPage = ref(1);
      const Blogs = ref([]);
      const SuggestList = ref([]);
      const poetry = reactive({
        content:''
      })
      const RandomPhoto = reactive({
        photo:''
      });
      const displayedItems = computed(() => {
        const startIndex = (currentPage.value - 1) * itemsPerPage;
        const endIndex = startIndex + itemsPerPage;
        return Blogs.value.slice(startIndex, endIndex);
      });
      const handleCurrentChange = (page) => {
        currentPage.value = page;
        window.scrollTo(0,850);
      };

      getAllBlog().then(res => {
        Blogs.value = res.List;
      })

      getShici().then(data => {
        poetry.content = data.content;
      })
      getAllSuggest().then((res) => {
        if (res.common.code === SUCCESS_REQUEST) {
          SuggestList.value = res.List
        }
      })

      onMounted(() => {
        getAllBlog();
        getAllSuggest();
        window.scrollTo(0, 0);
        getBing();
        getShici();
      });

      const getBing = () => {   //Bing每日图片接口
        if(sessionStorage.getItem("bingUrl")) {
          RandomPhoto.photo = JSON.parse(sessionStorage.getItem("bingUrl"));
        } else {
          getBingPhoto().then(data => {
            RandomPhoto.photo = data.data;
            sessionStorage.setItem("bingUrl",JSON.stringify(RandomPhoto.photo));
          })
        }
      }
      return {
        Blogs,
        displayedItems,
        RandomPhoto,
        handleCurrentChange,
        utilButton,
        SuggestList,
        poetry,
      };
    }
  };
  </script>
  
  <style scoped>
  .mains {
    height: 850px;
    width: 100%;
    text-align: center;
    background-size: cover;
    background-repeat: no-repeat;
    animation: slide-in2 1.8s
  }
  .content {
    border-bottom-left-radius: 10px;
    border-bottom-right-radius: 10px;
    background-color: white;
    padding: 10px;
  }
  .img {
    border-top-left-radius: 10px;
    border-top-right-radius: 10px;
    width: 100%;
    height: 200px;
  }
  .top{
    display: flex;
    position: relative;
  }
  .top .content{
    margin: 0;
  }
  .card{
    width: 350px;
    height: auto;
    transition: transform 0.3s;
    
  }
  .card:hover {
    transform: translateY(-5px);
  }


  .blog-referrer {
    background-color: rgba(255, 255, 255, 0.5);
    width: calc(70%);
    height: auto;
    min-height: 600px;
    margin: 10px auto 10px auto;
    border-radius: 10px;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    position: relative;
    vertical-align: bottom;
  }
  .blogs-container {
    width: calc(68.6%);
    margin-left: auto;
    margin-right: auto;
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
  }
  .suggest{
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    align-items: center;
  }
  @media screen and (max-width: 600px){
    .blogs-container{
      display: flex;
      flex-wrap: wrap;
      width: calc(100%);
    }
    .mains{
      background-position: center center;
    }
    .card{
      width: 99%;
      max-width: 350px;
      margin: auto;
    }
    .card2{
      width: 100%;
      margin: auto;
      margin-top: 20px;
    }
    .img .content{
      width: calc(100%);
    }
    .blog{
      margin-left: auto;
      margin-right: auto;
      width: calc(100%);
    }
    .blog-main{
      overflow: hidden;
      width: 100%;
    }
    .top-div .content{
      width: auto;
    }
    .img{
      background-size: cover;
    }
    .blog-referrer{
      width: calc(100%);
    }
    .poetry{
      margin-left: 10px;
      margin-right: 10px;
    }
  }
  
  .referrer-title {
    justify-content: center;
    display: flex;
    font-weight: bold;
    font-size: 24px;
    padding-top: 10px;
    color: rgb(24, 10, 58);
    width: 100%;
  }


  .preface{
    word-wrap: break-word; /* 换行规则 */
    overflow-wrap: break-word; /* 兼容性更好的换行规则 */
  }
  .create-time{
    margin-left: 5px;
    font-size: 13px;

  }

  .util{
    margin: 80px auto 0px auto;
  }
  .util-button {
    padding: 13px 27px;
    border-radius: 30px;
    color: white;
    background-color:transparent;
    border: none;
    font-size: 15px;
    position: relative;
    animation: bounce 1.8s ease-in-out infinite;
    outline: none;
  }
  .util-button:hover,
  .util-button:active {
    box-shadow: 0 4px 6px rgba(13, 12, 12, 0.3);
  }
  @keyframes bounce {
    0%, 100% {
      transform: translateY(0);
    }
    50% {
      transform: translateY(20px);
      opacity: 0.6;
    }
  }
.saw{
  font-size: 12px;
  color: #4b4a4a;
  margin-top: 4px;
}
  </style>
  