<template>
    <div class="body" >
        <el-main :style="{ 'background-image': `url(${RandomPhoto.photo})` }" class="main">
            <span class="main-title">友链</span>
                <div style="text-align:center;color:white;font-size:20px;font-weight:bold;">
               </div>
        </el-main>
        <el-card class="card-dd">
            <div class="friends">
                <div v-for="item in Friends" :key="item.id" class="friend" >
                    <div class="blog">
                        <a :href="`${item.url}`" class="link" target="_blank">
                            <img :src="item.photo" alt="" class="img2" style="height:100%;"/>
                            <div class="blog-t2">{{item.title}}</div>
                            <div class="preface2">{{item.preface}}</div>
                            <button class="button2">前往观摩</button>
                        </a>
                    </div>
                </div>
            </div>
        </el-card>
        <comments style="margin-left:auto;margin-right:auto;"></comments>
        <newFooter style="footer"></newFooter>
    </div>
</template>
<script>
import comments from '../comments/comments.vue';
import newFooter from './util/newFooter.vue';
import { onMounted, reactive, ref } from 'vue';
import {getAllFriend} from '../../api/friendFunc'
export default{
    name:"friendLink",
    components:{
        newFooter,
        comments,
    },
    setup() {
        const Friends = ref([]);
        getAllFriend().then((res) => {
            Friends.value = res.list;
        })
        const RandomPhoto = reactive({
            photo:''
        });

        RandomPhoto.photo = JSON.parse(sessionStorage.getItem("bingUrl"));
        onMounted(() => {
            getAllFriend();
            window.scrollTo(0,0);
        })

        return{
            Friends,
            RandomPhoto
        }
    }
}
</script>
<style scoped>
.main{
    animation: slide-in2 1.8s

}
.bottom-title{
    font-size: 25px;
    font-weight: bold;
}
.card-dd{
    margin-top:320px;
    width: calc(68%);
    height: auto;
    min-height: 300px;
    border: none;
    background-color:rgb(181, 180, 180,0);
}
@media screen and (max-width: 600px) {
    .card-dd{
        width: calc(98%);
    }
    .friend{
        width: 100%;
        min-width: 350px;
    }
}
.friend{
    display: flex;
    position: relative;
}
.friends{
    display: flex;
    flex-wrap: wrap;
    margin: 10px auto 10px auto;
    justify-content: center;
}

.blog{
    margin: 20px;
    width: 500px;
}
.link{
    text-align: center;
    margin-top: 30px;
    margin: auto;
    display: flex;
    position: relative;
}
.toggle {
    margin-bottom: 20px;
    border: 1px solid #f0f0f0;
    border-top-width: 1px;
    border-right-width: 1px;
    border-bottom-width: 1px;
    border-left-width: 1px;
    border-top-style: solid;
    border-right-style: solid;
    border-bottom-style: solid;
    border-left-style: solid;
    border-top-color: rgb(240, 240, 240);
    border-right-color: rgb(240, 240, 240);
    border-bottom-color: rgb(240, 240, 240);
    border-left-color: rgb(240, 240, 240);
    border-image-source: initial;
    border-image-slice: initial;
    border-image-width: initial;
    border-image-outset: initial;
    border-image-repeat: initial;
}
.toggle > .toggle-button {
    padding: 6px 15px;
    background: #f0f0f0;
    color: #1f2d3d;
    cursor: pointer;
}
.toggle > .toggle-content {
    margin: 30px 24px;
}
.fa, .fa-brands, .fa-classic, .fa-regular, .fa-sharp, .fa-solid, .fab, .far, .fas {
    -moz-osx-font-smoothing: grayscale;
    -webkit-font-smoothing: antialiased;
    display: var(--fa-display,inline-block);
    font-style: normal;
    font-variant: normal;
    line-height: 1;
    text-rendering: auto;
}
.note.warning.modern {
    border-color: #fae4cd;
    background: #fcf4e3;
    color: #8a6d3b;
}
.note.modern {
    border: 1px solid transparent !important;
    background-color: #f5f5f5;
    color: #202124;
}
.note:not(.no-icon) {
    padding-left: 3em;
}
.note {
    position: relative;
    margin: 0 0 20px;
    padding: 15px;
    border-radius: 3px;
}
ol {
    display: block;
    list-style-type: decimal;
    margin-block-start: 1em;
    margin-block-end: 1em;
    margin-inline-start: 0px;
    margin-inline-end: 0px;
    padding-inline-start: 40px;
}
.note.danger.modern {
    border-color: #ebcdd2;
    background: #f2dfdf;
    color: #a94442;
}
a {
    color: #99a9bf;
    text-decoration: none;
    word-wrap: break-word;
    -webkit-transition: all 0.2s;
    -moz-transition: all 0.2s;
    -o-transition: all 0.2s;
    -ms-transition: all 0.2s;
    transition: all 0.2s;
    overflow-wrap: break-word;
}
.note.info.modern {
    border-color: #b3e5ef;
    background: #d9edf7;
    color: #31708f;
}

</style>
