<template>
    <div class="body">
        <el-main :style="{ 'background-image': `url(${RandomPhoto.photo})` }" class="main">
             <span class="main-title">留言区</span>
        </el-main>
        <el-card class="card">
            <div class="bottom-title">
                留言板
            </div>
            <div class="yan">欢迎留言，我会随时回复您的问题和意见，让我们共同探讨和理解更多的话题！</div>
            <img class="img" src="https://qiniu.lllcnm.cn/img/gif/%E6%9A%96%E8%90%8C%E7%86%8A%E7%8C%AB%E7%99%BD%E5%B0%8F%E8%83%96%E5%BE%AE%E4%BF%A1%E8%A1%A8%E6%83%850_%E7%88%B1%E7%BB%99%E7%BD%91_aigei_com.gif">
            <img class="img" style="float:right;" src="https://qiniu.lllcnm.cn/img/gif/%E6%9A%96%E8%90%8C%E7%86%8A%E7%8C%AB%E7%99%BD%E5%B0%8F%E8%83%96%E5%BE%AE%E4%BF%A1%E8%A1%A8%E6%83%852_%E7%88%B1%E7%BB%99%E7%BD%91_aigei_com.gif" alt="">
        </el-card>
        <comments style="margin-left:auto;margin-right:auto;"></comments>
        <newFooter class="footer"></newFooter>
    </div>
</template>
<script>
import newFooter from './util/newFooter.vue';
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome';
import { faHome, fas,faL } from '@fortawesome/free-solid-svg-icons';
import { library } from '@fortawesome/fontawesome-svg-core';
import comments from '../comments/comments.vue';
import { onMounted, reactive, ref } from 'vue';
import axios from 'axios';
library.add(faHome, fas,faL);
export default{
    name:"message",
    components:{
        newFooter,
        comments,
        FontAwesomeIcon,
    },
    setup() {
        onMounted(()=> {
            window.scrollTo(0,0);
        })
        const RandomPhoto = reactive({
            photo:''
        });
        RandomPhoto.photo = JSON.parse(sessionStorage.getItem("bingUrl")); 
        return{
            RandomPhoto
        }
    }
}
</script>
<style scoped>
.card{
    width: calc(68%);
    min-height: 180px;
    height: auto;
    border-radius: 10px;
    margin: 6px auto 20px auto;
    animation: slide-in 1s;
    margin-top: 320px;
    position: relative;
    vertical-align: bottom;
    border: none;
}
.bottom-title{
    color: rgb(12, 7, 32);
    font-size: 35px;
    font-weight: bold;
}
.yan{
    font-size:22px;text-align:center;margin-top:20px;
}
@media screen and (max-width: 600px) {
    .card{
        width: calc(98%);
    }
    .yan{
        font-size: 18px;
    }
}
.img{
    width: 100px;
    height: 100px;
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
</style>

