<template>
  <el-container>
    <el-main>
      <div>
        <el-button @click="deleteSomePhoto" type="danger">批量删除</el-button>
        <el-button @click="show = true" type="success">新增图片</el-button>
        <el-dialog title="新增图片" v-model="show" :lock-scroll="false" class="dialog" style="border-radius: 20px;" width="40%">
          <el-descriptions border column="1">
            <el-descriptions-item label="标题">
              <el-input type="text" v-model="data.data.title"></el-input>
            </el-descriptions-item>
            <el-descriptions-item label="定位">
              <el-input type="text" v-model="data.data.position"></el-input>
            </el-descriptions-item>
            <el-descriptions-item label="图片">
              <el-upload
                  class="upload-demo"
                  drag
                  :http-request="handleMainPhoto"
                  multiple
              >
                <div class="el-upload__text">
                  拖拽文件到此处，或<em>点击上传</em>
                </div>
                <template #tip>
                </template>
              </el-upload>
            </el-descriptions-item>
            <el-descriptions-item label="日期">
              <el-date-picker
                  v-model="data.data.date"
                  type="date"
                  placeholder="Pick a day"
                  format="YYYY-MM-DD"
                  value-format="YYYY-MM-DD"
                  :size="100"
              />
            </el-descriptions-item>
          </el-descriptions><br>
          <div style="float: right">
            <el-button @click="show = false">取消</el-button>
            <el-button @click="save" type="primary">提交</el-button>
          </div><br>
        </el-dialog>
      </div>
      <el-image-viewer
          v-if="showImageViewer"
          :url-list="photoUrl"
          @close="handleImageHide"
          :z-index="9999"
          :initial-index="selectedIndex"
      ></el-image-viewer>
      <el-row :gutter="20" class="photo-grid">
        <el-col v-for="photo in BlogPhoto" :key="photo.id" :span="6">
          <el-card :body-style="{ padding: '0px' }" class="photo-card" @dblclick="openOptions(photo)">
            <div :style="{ backgroundImage: `url(${photo.photo})` }" class="card-background">
              <div class="card-content">
                <h3>{{ photo.title }}</h3>
                <p>{{ photo.position }}</p>
                <p>{{ photo.date }}</p>
                <el-checkbox v-model="photo.selected" @change="handlePush(photo.id)" />
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>
      <el-dialog title="选项" v-model="optionsVisible" style="border-radius: 20px;" width="30%">
        <p>请选择操作：</p>
        <div class="actions">
          <el-button type="primary" @click="openLink(selectedPhoto.photo)">查看</el-button>
          <el-button type="danger" @click="deleteBlogPhoto(selectedPhoto.id)">删除</el-button>
        </div>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import {onMounted, reactive, ref} from 'vue';
import {ElMessage} from 'element-plus';
import {DeletePhoto, GetAllPhoto, savePhoto} from "@/view/api/photo";
import {confirmFunc} from "@/view/api/util";
import {UploadFile} from "@/view/api/tool";

export default {
  name: "blogPhoto",
  setup() {
    const BlogPhoto = ref([]);
    const photoUrl = ref([]);
    const getAllPhoto = () => {
      GetAllPhoto().then((res) => {
        BlogPhoto.value = res.data;
        BlogPhoto.value.forEach((item) => {
          photoUrl.value.push(item.photo);
        });
      });
    };
    const deleteBlogPhoto = (id) => {
      confirmFunc().then(() => {
        DeletePhoto(id).then((res) => {
          if (res) {
            ElMessage.success("删除成功");
            getAllPhoto();
          }
        });
      });
    };
    const selectedRows = ref([]);
    const handlePush = (id) => {
      if (selectedRows.value.includes(id)) {
        const index = selectedRows.value.indexOf(id);
        selectedRows.value.splice(index, 1);
      } else {
        selectedRows.value.push(id);
      }
      console.log(selectedRows.value);
    };

    const deleteSomePhoto = () => {
      confirmFunc().then(() => {
        selectedRows.value.forEach((row) => {
          DeletePhoto(row)
        });
      });
      getAllPhoto();
    };

    onMounted(() => {
      getAllPhoto();
    });

    const showImageViewer = ref(false);
    const handleImageHide = () => {
      showImageViewer.value = false;
    };
    const selectedIndex = ref(0);
    const handelEventPhoto = () => {
      selectedIndex.value = photoUrl.value.indexOf(event.target.src);
      showImageViewer.value = true;
    };
    const data = reactive({
      data: {
        title: '',
        position: '',
        photo: '',
        date: ''
      }
    });
    const handleMainPhoto = (param) => {
      UploadFile(param).then((res) => {
        ElMessage.success(res.data.info);
        data.data.photo = res.data.data;
      });
    };
    const save = () => {
      savePhoto(data).then(() => {
        ElMessage.success("保存成功");
        show.value = false;
        data.data = {
          title: '',
          position: '',
          photo: '',
          date: ''
        };
        getAllPhoto();
      });
    };
    const show = ref(false);

    const optionsVisible = ref(false);
    const selectedPhoto = ref({});

    const openOptions = (photo) => {
      selectedPhoto.value = photo;
      optionsVisible.value = true;
    };

    const openLink = (url) => {
      window.open(url, '_blank');
    };

    return {
      BlogPhoto,
      deleteBlogPhoto,
      selectedRows,
      handlePush,
      deleteSomePhoto,
      showImageViewer,
      handleImageHide,
      selectedIndex,
      handelEventPhoto,
      photoUrl,
      data,
      handleMainPhoto,
      show,
      save,
      optionsVisible,
      selectedPhoto,
      openOptions,
      openLink,
    };
  },
};
</script>

<style scoped>
@import url('../../assets/css/main.css');

.photo-grid {
  margin-top: 20px;
}

.photo-card {
  transition: transform 0.2s;
  border-radius: 10px;
  overflow: hidden;
}

.photo-card:hover {
  transform: translateY(-5px);
}

.card-background {
  background-size: cover;
  background-position: center;
  height: 200px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.card-content {
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  padding: 20px;
  text-align: center;
  width: 100%;
}

.actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}
</style>
