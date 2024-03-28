<template>
    <div class="body" style="display:flex;justify-content: center;">
        <el-container class="main" :style="{ 'background-image': `url(${getData.photo})` }">
            <el-header></el-header>
            <span class="origin-title">{{getData.title}}</span>
        </el-container>
        <div class="div" style="width:calc(100%);">
          <el-card class="card">
              <div>
                  <span class="hd">
                    <font-awesome-icon :icon="['fas', 'calendar-days']" /> 发布日期:
                     <span style="font-size:13px;"> {{getData.createTime}}</span></span> 
                  <span class="hd">
                    <font-awesome-icon :icon="['fas', 'calendar-days']" />
                     更新日期:  <span style="font-size:13px;">{{getData.updateTime}}</span>
                  </span>
                  <span class="hd">
                    <font-awesome-icon :icon="['fas', 'eye']" /> 浏览量: {{getData.visits}}</span>
                  <el-tooltip :content="`发布日期：${getData.createTime} 更新日期: ${getData.updateTime} 浏览量: ${getData.visits}`" placement="top">
                    <font-awesome-icon :icon="['fas', 'signature']" class="date" />
                  </el-tooltip>
                  <div style="float:right;" class="tag"><span class="bq" style="background-color:rgb(20, 9, 59);color:white;padding:2px;
                  font-size:14px;border-top-left-radius: 5px;border-bottom-left-radius:5px;">
                    <font-awesome-icon :icon="['fas', 'droplet']" /> 标签</span><span style="background-color:rgb(207, 115, 130);
                    color:white;padding:2px;font-size:14px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
                      {{getData.tag}}</span>
                    </div>
              </div>
              <el-divider/>
              <el-image-viewer
              v-if="showImageViewer"
              :url-list="imageList"
              @close="handleImageHide"
              :z-index="9999"
              :initial-index="selectedIndex"
              hide-on-click-modal
              ></el-image-viewer>
              <div>
                  <span style="color:#34495e;font-size:20px;margin-left:35px;font-weight:bold;">{{getData.preface}}</span>
                  <v-md-editor :model-value="getData.content" mode="preview"  :include-level="[3, 4]"
                  @image-click="handleImageClick" @copy-code-success="handleCopyCodeSuccess"></v-md-editor>
              </div>
              <div class="footer-ban">
                <div class="quote">
                  <p><span class="statement"><font-awesome-icon :icon="['fas', 'user']" /> 
                    版权归属:</span> <span style="color:rgb(37, 173, 87);">红豆南墙</span></p>
                  <p><span class="statement"><font-awesome-icon :icon="['fas', 'link']" /> 本文链接:</span>
                     <a :href="`https://lllcnm.cn/watch/${getData.id}`" style="color:rgb(37, 173, 87);text-decoration: none;" target="_blank">
                      {{`https://lllcnm.cn/watch/${getData.id}`}}</a></p>
                  <p><span class="statement"><font-awesome-icon :icon="['fas', 'copyright']" /> 版权声明:</span>
                    若无注明均为<span class="highlight">红豆南墙</span>原创，转载请注明出处，感谢您的支持!</p>
                </div>
              </div>
  
              <div style="display:flex;align-items:center;justify-content: center;margin-top:10px;">
                <Tipping/>
              </div>
          </el-card>
         <comments :message="id" style="margin-left:auto;margin-right:auto;" class="comments"></comments>
         <newFooter></newFooter>
        </div>
    </div>
</template>
<script>
import {useRoute} from 'vue-router'
import {onMounted, reactive, ref} from 'vue'
import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome';
import {faHome, fas} from '@fortawesome/free-solid-svg-icons';
import {library} from '@fortawesome/fontawesome-svg-core';
import {ElImageViewer, ElMessage} from 'element-plus';
import 'prismjs/themes/prism-dark.css';
import 'prismjs/components/prism-javascript';
import 'prismjs/components/prism-css';
import 'prismjs/plugins/line-numbers/prism-line-numbers'
import comments from '../comments/comments.vue'
import newFooter from './util/newFooter.vue'
import Tipping from './util/Tipping.vue';
import {GetBlogById} from "@/api/blogFunc";

library.add(faHome,fas)
export default{
    components:{
        FontAwesomeIcon,
        comments,
        newFooter,
        ElImageViewer,
        Tipping
     },
    setup() {
        const getData = reactive({
            id:"",
            title:"",
            content:"",
            preface:"",
            updateTime:"",
            createTime:"",
            photo:"",
            tag:'',
            visits:""
        })
        const showImageViewer = ref(false);
        const selectedIndex = ref(0)
        const imageList = ref([])
        const url = `data:image/png;base64,${getData.photo}`;
        const route = useRoute();
        const id = route.params.id;

        onMounted(() => {
            getId()
            window.scrollTo(0,0);
        })
        const regex = /!\[.*?\]\((.*?)\)/g;
        const getId = async () => {
          const res = await GetBlogById(id);
          Object.keys(getData).forEach(key => {
            // eslint-disable-next-line no-prototype-builtins
            if (res.data.hasOwnProperty(key)) {
              getData[key] = res.data[key];
            }
          });
          let match;
          const extractedImages = [];
          while ((match = regex.exec(getData.content)) !== null) {
            extractedImages.push(match[1]);
          }
          // 将解析的图片地址添加到 imageList
          imageList.value = extractedImages;
        };

        const handleImageClick = () => {
          selectedIndex.value = imageList.value.indexOf(event.target.src)
          showImageViewer.value = true;
        }
        const handleImageHide = () => {
          showImageViewer.value = false
        }
        const handleCopyCodeSuccess = () => {
          ElMessage.success("复制成功")
        }

        return {
            getData,
            url,
            id,
            showImageViewer,
            imageList,
            selectedIndex,
            handleImageClick,
            handleImageHide,
            handleCopyCodeSuccess,
        }
    }
}
</script>
<style scoped>
.main{
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center center;
    width: calc(100%);
    height: 400px;
    position: absolute;
    animation: slide-in2 1.8s
}
.card{
    margin: 250px auto 20px auto;
    min-height: 400px;
    border-radius: 10px;
    width: calc(68%);
    min-width:68%;
    height: auto;
    position: relative;
    vertical-align: bottom;
}
.statement{
  font-weight: bold;
}
.hd{
    margin-right: 20px;
}
.origin-title{
    font-size: 45px;
    font-weight: bold;
    color: white;
    display: flex;
    align-items: center;
    justify-content: center;
    margin-top: 100px;
}

.date{
  display: none;
  float: left;
}

@media screen and (max-width: 600px) {
  .origin-title{
    font-size: 25px;
  }
  .main{
    width: calc(100%);
  }
  .hd{
    display: none;
  }
  .div{
    width: calc(100%);
  }
  .card{
    width: calc(98%);
  }
  .date{
    display: block;
  }
}
.footer-ban {
  margin-top: 20px;
  background-color: #f5f5f5;
  padding: 10px;
  border-left: 4px solid #ddd;
  width: calc(90%);
}

.quote {
  padding-left: 10px;
  color: #555555;
}

.highlight {
  color: rgb(37, 173, 87);
  font-weight: bold;
}
</style>