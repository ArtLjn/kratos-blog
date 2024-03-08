<template>
    <div class="body" @resize="handel" >
            <el-main class="main" :style="{ 'background-image': `url(${image.photo})` }" >
                <span class="main-title">关于我</span>
            </el-main>
            <div class="follow">
                <img src="../../assets/image/touxiang.jpg" class="photo"/>
                <el-card class="card-dd">
                    <div style="float:left;" class="tag"><span class="bq" style="background-color:rgb(22, 14, 53);color:white;padding:2px;
                        font-size:14px;border-top-left-radius: 5px;border-bottom-left-radius:5px;">
                          <font-awesome-icon :icon="['fas', 'droplet']" /> 文章</span><span style="background-color:rgb(11, 156, 167);
                          color:white;padding:2px;font-size:14px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
                            {{blogCount}} 篇</span>
                          </div>
                   <div class="bottom-title">红豆南墙</div>
                   <div style="width:80%;margin:auto;">
                    <p class="jshao">我是红豆南墙，一个对生活充满好奇心和热情的人。

                        我一直对探索新事物和挑战自己充满激情。通过我的博客，我希望与大家分享我的思考和经历，以及对各种各样话题的见解。
                        
                        我的兴趣广泛，并不局限于特定领域。我喜欢阅读各种类型的书籍，包括小说、科幻、历史等。我也喜欢观看电影和探索不同的艺术形式。这些经历激发了我对文化、艺术和人类历史的深思。
                        
                        此外，旅行也是我生活中不可或缺的一部分。通过走访不同的地方，我能够体验不同的文化和风俗，拓宽自己的眼界。这些旅行经历也让我更加珍惜现在的生活，并对全球化的世界有更深刻的理解。
                        
                        在我的博客中，你会找到关于个人成长、心理学、健康生活、社会问题等方面的文章。我希望通过分享我的思考和见解，能够激发读者们的思考，并为大家提供一些有意义的讨论。
                        
                        感谢大家阅读我的博客，并期待与你们的交流和互动！谢谢！</p>
                   </div>
                   <div class="tu">
                        <zhuanTu class="zt"/>
                        <bingTu class="bt"/>
                   </div>
                   <Niandu class="niandu"/>
                </el-card>
            </div>
        <newFooter class="footer"></newFooter>
    </div>
</template>
<script>
import bingTu from './util/bingTu.vue'
import newFooter from './util/newFooter.vue'
import zhuanTu from './util/zhuanTu.vue'
import { onMounted, reactive, ref } from 'vue'
import { getAllBlog } from '@/api/blogFunc'
export default{
    name:"followme",
    components:{
        newFooter,
        bingTu,
        zhuanTu,
    },
    setup() {
        const image = reactive({
            photo:''
        })
        image.photo = JSON.parse(sessionStorage.getItem("bingUrl"));
        onMounted(() => {
            window.scrollTo(0,0);
        })
        const Blogs = ref([]);
        const blogCount = ref('')
        getAllBlog().then(data => {
            Blogs.value = data.list;
            blogCount.value = Blogs.value.length;
        })
      return{
        image,
        Blogs,
        blogCount
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
    height: 500px;
    position: absolute;
    animation: slide-in2 1.8s
}
.main-title{
    margin-left: auto;
    margin-right: auto;
    display: flex;
    margin-bottom: 0;
}
.follow{
    display: flex;
    justify-content: center;
    animation: slide-in3 1s;
    margin-top:320px;
}
.card-dd{
    margin-top: 80px;
    background-color: white;
}
.bottom-title{
    color: rgb(25, 12, 74);
    margin-top: 100px;
    font-size: 30px;
}

.jshao{
    font-size: 16px;
    text-indent: 10px;
    letter-spacing: 5px;
}

.tu{
    width: 100%;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
}
.zt{
    width: 48%;
}
.bt{
    width: 48%;
}
@media screen and (max-width: 600px) {
    .tu{
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
    .zt{
        width: 90%;
    }
    .bt{
        width: 90%;
    }
    .card-dd{
        width: 98%;
    }
}
.photo {
    width: 160px;
    height: 160px;
    border-radius: 50%;
    display: flex;
    z-index: 2;
    position: absolute;
    transition: transform 0.3s ease;
  }
  
  .photo:hover {
    animation: rotate 2s infinite;
  }
  
  @keyframes rotate {
    0% {
      transform: rotate(0deg);
    }
    100% {
      transform: rotate(360deg);
    }
  }
  
  .photo:not(:hover) {
    transform: rotate(0deg) !important;
  }
</style>