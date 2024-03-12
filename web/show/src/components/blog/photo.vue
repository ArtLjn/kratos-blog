<template>
    <div class="body" >
        <el-main :style="{ 'background-image': `url(${RandomPhoto.photo})` }" class="main">
            <span class="main-title">相册</span>
                <div class="hua">
                    当人们开始忙碌的生活，有时候会忘记珍惜生活中的美好瞬间。而照片便是记录美好瞬间的最好方式。</div>
        </el-main>
        <el-image-viewer
        v-if="showImageViewer"
        :url-list="photoUrl"
        @close="handleImageHide"
        :z-index="9999"
        :initial-index="selectedIndex"
        hide-on-click-modal
        placeholder="图片怎么丢失了😜" 
        ></el-image-viewer>
        <el-card class="card-dd"> 
            <div class="blog-p">
                <div v-for="item in PhotoList" :key="item.id" class="blog">
                    <div class="bp">
                        <el-tooltip :content="`${item.date} ${item.position}`" placement="top">
                            <div class="ok">
                                <img v-lazy="item.photo" class="photo" @click="handelEventPhoto"/>
                                <div class="lt-title">{{item.title}}</div>
                            </div>
                        </el-tooltip>
                    </div>
                </div>
            </div>
        </el-card>
        <newFooter></newFooter>
    </div>
</template>
<script>
import newFooter from './util/newFooter.vue';
import { onMounted, reactive, ref } from 'vue';
import {getAllPhoto} from '@/api/photoFunc'
export default{
    components:{
        newFooter
    },
    setup() {
      const PhotoList = ref([]);
      const photoUrl = ref([]);
      getAllPhoto().then((res) => {
        PhotoList.value = res.data
        res.data.forEach(val => {
          photoUrl.value.push(val.photo)
        })
      })
      const RandomPhoto = reactive({
          photo:''
      });

      RandomPhoto.photo = JSON.parse(sessionStorage.getItem("bingUrl"));
      const showImageViewer = ref(false);
      const handleImageHide = () => {
        showImageViewer.value = false
      }
      const selectedIndex = ref(0)
      const handelEventPhoto = () => {
          const index = photoUrl.value.indexOf(event.target.src); // 获取点击的图片在photoUrl列表中的索引值
          // console.logs(index); // 输出索引值
          selectedIndex.value = index
          showImageViewer.value = true;

      }

      onMounted(() => {
          getAllPhoto();
          window.scrollTo(0,0);
      })
      return{
          PhotoList,
          RandomPhoto,
          photoUrl,
          handleImageHide,
          showImageViewer,
          handelEventPhoto,
          selectedIndex,
      }
    }
}
</script>
<style scoped>
.card-dd{
    margin: 320px auto 20px auto;
    min-height: 400px;
    border-radius: 10px;
    width: calc(68%);
    min-width:68%;
    height: auto;
    position: relative;
    vertical-align: bottom;
    background-color: rgb(135, 109, 109,0);
    border: none;
}
.photo{
    width: 100%;
    border-radius: 10px;
    min-height: 250px;
    height: auto;
    animation: fadeInOut 2s ease-in-out;
  }
.ok{
    position: relative;
    width: 100%;
    display: flex;
}
.blog-p{
    border-radius: 10px;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
.lt-title{
    position: absolute;
    color: white;
    bottom: 30px;
    left: 10px;
    font-size: 25px;
    text-align: center;
}


.blog{
    width: 350px;
}
.hua{
    text-align:center;color:white;font-size:20px;font-weight:bold;
}
@media screen and (max-width:600px) {
    .card-dd{
        width: calc(89%);
    }
    .hua{
        font-size: 16px;
    }
}

.main{
    background-repeat: no-repeat;
    background-size: cover;
    background-position: center center;
    width: calc(100%);
    height: 500px;
    position: absolute;
    animation: slide-in2 1.8s
}
.bp{
    margin: 10px;
    position: relative;
}
</style>
