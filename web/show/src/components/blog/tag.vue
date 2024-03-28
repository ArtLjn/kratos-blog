<template>
    <div class="body">
        <el-main :style="{ 'background-image': `url(${RandomPhoto.photo})` }" class="main">
            <span class="main-title">TAG</span>
                <div style="text-align:center;color:white;font-size:20px;font-weight:bold;">
                    点亮未来，追逐梦想</div>
        </el-main>
        <el-card class="card-dd">
            <div class="bottom-title">
                <font-awesome-icon :icon="['fas', 'tags']" style="padding-top:5px;padding-right:10px;" />文章标签</div>
            <div class="all-tag">
                <div v-for="tag in tags" :key="tag.id" class="tags">
                    <div class="tag" :style="`background-color: rgb(${Math.floor(Math.random() * 256)},
                     ${Math.floor(Math.random() * 256)}, ${Math.floor(Math.random() * 256)}, 0.7)`" @click="getTagName(tag.tagName)">{{ tag.tagName }}</div>
                </div>
            </div>
        </el-card>         
        <div class="div-card">
            <div class="card3">
                <div v-for="blog in Blogs" :key="blog.id" class="blog">
                    <router-link :to="`/watch/${blog.id}`" class="top" target="_blank">
                        <img v-lazy="`${blog.photo}`" class="img2" v-if="blog.photo">
                        <div class="blog-t2">{{ blog.title }}</div>
                        <div class="preface2">{{blog.preface}}</div>
                        <button class="button2" v-if="blog.title
                        ">阅读更多</button>
                    </router-link>
                </div>
            </div>  
        </div>
        <newFooter class="footer"></newFooter>
    </div>
</template>
<script>
import newFooter from './util/newFooter.vue';
import { onMounted, reactive, ref } from 'vue';
import {GetTag, GetTagByName} from "@/api/tag";
export default{
    components:{
        newFooter
    },
    setup() {
        const tags = ref([])
        const Blogs = ref([]);

        GetTag().then((res) => {
          tags.value = res.data;
        })
        const getTagName = (tag) => {
          GetTagByName(tag).then((res) => {
            Blogs.value = res.List;
          })
        }

        onMounted(() => {
            GetTag();
            window.scrollTo(0,0);
        })
        const RandomPhoto = reactive({
            photo:''
        });

        RandomPhoto.photo = JSON.parse(sessionStorage.getItem("bingUrl"));

        return{
            tags,Blogs,getTagName,RandomPhoto
            
        }
    }
}
</script>
<style scoped>
.main{
    animation: slide-in2 1.8s

}
.card-dd{
    width: 68%;
    background-color: white;
    height: auto;
    min-height: auto;
    margin-top:320px;
}
.bottom-title{
    font-size: 33px;
    color:rgb(20, 10, 86) ;
}
.all-tag{
    display: flex;
    justify-content: center;
    align-items: center;
    margin-top: 15px;
    flex-wrap: wrap;
    flex-direction: row;
}
.tags{
    display: inline-flex;
    flex-wrap: wrap;
    flex-direction: row;
}
.tag{
    color: rgb(8, 5, 29);
    font-size: 15px;
    font-weight: bold;
    margin-right: 20px;
    margin-bottom: 20px;
    padding: 10px;
    border-radius: 10px;
}

.top{
    display: flex;
    position: relative;
  }
.div-card{
    width: calc(68%);
    margin: auto;
    height: auto;
}
.blog{
    margin: 20px;
    width: 500px;
    animation: fadeInOut 2s ease-in-out;
}
.card3{
    display: flex;
    flex-wrap: wrap;
}
@media screen and (max-width: 600px) {
    .card-dd{
        width: calc(86%);
    }
    .blog{
        width: 100%;
    }
    .blog{
        width: calc(100%);
    }
    .div-card{
        width: calc(100%);
    }
}

</style>
