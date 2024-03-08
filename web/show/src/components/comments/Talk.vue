<template>
    <div class="body" >
        <el-main :style="{ 'background-image': `url(${RandomPhoto.photo})` }" class="main">
            <span class="main-title">
                说说</span>
                <div style="text-align:center;color:white;font-size:20px;font-weight:bold;">
                    {{shici.content}}</div>
        </el-main>
        <el-card class="card-dd">
           <div v-for="(item,index) in talkList" :key="index">
            <div class="talk">
                <img src="../../assets/image/touxiang.jpg" class="photo"/>
                <div class="shuoshuo">
                    <div class="content">{{item.content}}</div>
                    <span class="date">{{item.date}}</span>
                </div>
            </div>
           </div>
        </el-card>   
        <comments style="margin:auto;margin-bottom:150px;" class="comments"/>      
        <newFooter class="footer"></newFooter>
    </div>
</template>
<script>
import { onMounted, reactive, ref } from 'vue';
import axios from 'axios';
import newFooter from '../blog/util/newFooter.vue';
import comments from './comments.vue';
import {getShici} from '../../api/blogFunc';
export default {
  name: "talk",
  components: {
    newFooter,
    comments
  },
  setup() {
    const talkList = ref([]);
    const shici = reactive({
      content: ""
    });

    getShici().then(data => {
        shici.content = data.content;
      }) 

    const getAllTalk = () => {
      axios.get("/api/getAllTalk").then((response) => {
        talkList.value = response.data.list.sort((a, b) =>
          new Date(b.date) - new Date(a.date)
        );
      });
    };

    onMounted(() => {
      getAllTalk();
      getShici();
      window.scrollTo(0, 0);
    });
    const RandomPhoto = reactive({
            photo:''
    });
    RandomPhoto.photo = JSON.parse(sessionStorage.getItem("bingUrl")); 
    return {
      talkList,
      shici,
      RandomPhoto
    };
  }
}
</script>
<style scoped>
.card-dd{
    margin: 0px auto 20px auto;
    min-height: 400px;
    border-radius: 10px;
    width: calc(68%);
    min-width:68%;
    height: auto;
    margin-top:320px;
    background-color: rgb(135, 109, 109,0);
    border: none;
    animation: slide-in 1s;
}
.photo{
    border-radius: 50%;
    height: 50px;
    width: 50px;
  }
  .talk{
    width:100%;display:flex;
    margin: 10px;
  }
  .content{
    word-wrap: break-word;
    overflow-wrap: break-word;
    padding: 20px;
    margin-left: 10px;
    margin-right: 10px;
  }
  .shuoshuo{
    height: auto;
    min-height: 50px;
    margin-left: 20px;
    width: 90%;
    background-color: rgb(247, 233, 233,0.3);
    border-radius: 5px;
    text-indent: 20px;
  }
  .date{
    padding-right: 20px;
    padding-bottom: 5px;
    font-size: 10px;
    float: right;
  }
  @media screen and (max-width:600px) {
    .card-dd{
        width: 98%;
    }
  }
  .comments{
    background-color: rgb(234, 230, 230,0);
    border: none;
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