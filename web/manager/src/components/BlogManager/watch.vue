<template>
      <div class="card">
        <div>
          <span class="hd">发布日期: {{ getData.createTime }}</span>
          <span class="hd">更新日期: {{ getData.updateTime }}</span>
          <div style="float:right;" class="tag"><span class="bq" style="background-color:rgb(20, 9, 59);color:white;padding:2px;
            font-size:14px;border-top-left-radius: 5px;border-bottom-left-radius:5px;">
              <font-awesome-icon :icon="['fas', 'droplet']" /> 标签</span><span style="background-color:rgb(207, 115, 130);
              color:white;padding:2px;font-size:14px;border-top-right-radius:5px;border-bottom-right-radius:5px;">
                {{getData.tag}}</span>
              </div>
          <el-divider/>
        </div>
        <el-image-viewer
        v-if="showImageViewer"
        :url-list="imageList"
        @close="showImageViewer=false"
        :z-index="9999"
        :initial-index="selectedIndex"
        hide-on-click-modal
        ></el-image-viewer>
        <div>
          <h1 style="margin-left:36px;">{{ getData.preface }}</h1>
          <v-md-editor :model-value="getData.content" mode="preview" style="width:100%;" @image-click="handleImageClick"
          @copy-code-success="handleCopyCodeSuccess"></v-md-editor>
        </div>
      </div>
  </template>
  
<script>
  import { useRoute } from 'vue-router';
  import { onMounted, reactive, ref } from 'vue';
  import { ElMessage } from 'element-plus';
  import {GetBlogById} from "@/components/api/blog";

  export default {
    name: 'watch',
    setup() {
      let getData = reactive({
        title: '',
        content: '',
        preface: '',
        updateTime: '',
        createTime: '',
        tag: ''
      });
      const route = useRoute();
      const id = route.params.id;
  
      onMounted(() => {
        getId(id);
        window.scrollTo(0, 0);
      });
      const regex = /!\[.*?\]\((.*?)\)/g;

      const imageList = ref([])
      const showImageViewer = ref(false);

      const selectedIndex = ref(0)

      const getId = async () => {
        const res = await GetBlogById(id);
        Object.keys(getData).forEach(key => {
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
          const index = imageList.value.indexOf(event.target.src);
          selectedIndex.value = index
          showImageViewer.value = true;
      }

      const handleCopyCodeSuccess = () => {
          ElMessage.success("复制成功")
        }
      return {
        getData,
        imageList,
        showImageViewer,
        selectedIndex,
        handleImageClick,
        handleCopyCodeSuccess
      };
    },
  };
</script>
<style scoped>
.card{
    margin: 0px auto auto auto;
    min-height: 400px;
    border-radius: 10px;
    width: 100%;
    height: auto;
}
.hd{
    margin-right: 20px;
}
</style>