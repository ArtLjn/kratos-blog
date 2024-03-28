<template>
  <div class="body">
    <el-main :style="{ 'background-image': `url(${RandomPhoto.photo})` }" class="main">
      <span class="main-title">归档</span>
      <div style="text-align:center;color:white;font-size:20px;font-weight:bold;"></div>
    </el-main>
    <el-card class="card-dd">
      <div class="time-dd">
        <el-timeline class="time-tt">
          <el-timeline-item v-for="item in displayedItems" :key="item.id" style="display:flex;" :timestamp="item.createTime" center hide-timestamp>
            <div class="blog">
              <router-link :to="`/watch/${item.id}`" class="top">
                  <img  v-lazy="`${item.photo}`" class="img2">
                <div class="blog-t2">{{item.title}}</div>
                <div class="preface2">{{item.preface}}</div>
                <button class="button2">阅读更多</button>
              </router-link>
            </div>
            <div class="time">{{item.createTime}}</div>
          </el-timeline-item>
        </el-timeline>
      </div>
      <div class="phone">
        <div v-for="item in displayedItems" :key="item.id" style="display:flex;" :timestamp="item.createTime" center hide-timestamp>
          <div class="blog">
            <router-link :to="`/watch/${item.id}`" class="top">
                <img  v-lazy="`${item.photo}`" class="img2">
              <div class="blog-t2">{{item.title}}</div>
              <div class="preface2">{{item.preface}}</div>
              <button class="button2">阅读更多</button>
            </router-link>
            <div class="time">{{item.createTime}}</div>
          </div>
        </div>
      </div>
    </el-card>
    <el-pagination  :total="Blogs.length" @current-change="handleCurrentChange" pager-count="3"  layout="prev, pager, next"
    background class="page" >
  </el-pagination>
    <newFooter></newFooter>
  </div>
</template>

<script>
import newFooter from './util/newFooter.vue';
import { computed, onMounted, reactive, ref } from 'vue';
import { getAllBlog } from '@/api/blogFunc';

export default {
  // eslint-disable-next-line vue/multi-word-component-names
  name: "archives",
  components: {
    newFooter
  },
  setup() {
    const Blogs = ref([]);
    getAllBlog().then(data => {
        Blogs.value = data.List;
    })
    const RandomPhoto = reactive({
      photo: ''
    });

    RandomPhoto.photo = JSON.parse(sessionStorage.getItem("bingUrl"));

    onMounted(() => {
      window.scrollTo(0,0);
    });

    /**
     * 分页
     */
    const itemsPerPage = 10;
    const currentPage = ref(1);
    const displayedItems = computed(() => {
      const startIndex = (currentPage.value - 1)  * itemsPerPage;
      const endIndex = startIndex + itemsPerPage;
      return Blogs.value.slice(startIndex,endIndex);
    })
    const handleCurrentChange = (page) => {
      currentPage.value = page;
      window.scrollTo(0,0);
    }
    return {
      RandomPhoto,
      displayedItems,
      handleCurrentChange,
      Blogs
    }
  }
}
</script>
<style scoped>
.main{
  animation: slide-in2 1.8s
}
.bottom-title{
    font-size: 30px;
    color: #021120;
    font-weight: bold;
}

.blog{
    margin: 20px;
    width: 500px;
}
.top{
  display: flex;
  position: relative;
}
.card-dd{
  background-color: rgb(225, 218, 218,0);
  border: none;
  margin-top:320px;
  animation: slide-in3 1s
}
.time{
  color: white;
  font-size: 20px;
  margin-top: 10px;
}
.phone{
  display: none;
}
@media screen and (max-width: 600px) {
  .card-dd{
    width: calc(100%);
  }
  .blog{
    width: 100%;
    margin-right: 10px;
  }
  .phone{
    display: block;
  }
  .time-dd{
    display: none;
  }

}
.time-tt{
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>