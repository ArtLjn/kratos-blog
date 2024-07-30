<template>
  <el-container>
    <el-main>
      <div >
        <el-button @click="deleteSomePhoto" type="danger" plain>批量删除</el-button>
        <addPhoto></addPhoto>
      </div>
      <el-image-viewer
          v-if="showImageViewer"
          :url-list="photoUrl"
          @close="handleImageHide"
          :z-index="9999"
          :initial-index="selectedIndex"
      ></el-image-viewer>
      <el-table :data="BlogPhoto" class="table" border :lazy="true" stripe height="600">
        <el-table-column type="selection" label="序号" width="80" >
          <template #default="{ row }">
            <el-checkbox v-model="row.selected" @change="handlePush(row.id)"></el-checkbox>
          </template>
        </el-table-column>
        <el-table-column label="图标">
          <template #default="{row}">
            <el-image :src="`${row.photo}`" class="img" style="width:130px;" @click="handelEventPhoto"/>
          </template>
        </el-table-column>
        <el-table-column  label="标题" prop="title"></el-table-column>
        <el-table-column label="时间" prop="date"></el-table-column>
        <el-table-column fixed="right" label="操作">
          <template #default="{row}">
            <el-button  @click="deleteBlogPhoto(row.id)" type="danger" round>删除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>
<script>
import {onMounted, ref} from 'vue'
import addPhoto from './addPhoto.vue'
import {ElMessage} from 'element-plus'
import {DeletePhoto, GetAllPhoto} from "@/view/api/photo";
import {confirmFunc} from "@/view/api/util";

export default{
    name:"blogPhoto",
    components:{
        addPhoto
    },
    setup() {
        const BlogPhoto = ref([])
        const photoUrl = ref([])
        const getAllPhoto =()=> {
          GetAllPhoto().then((res) => {
            BlogPhoto.value = res.data
            BlogPhoto.value.map((item) => {
              photoUrl.value.push(item.photo);
            })
          })
        }
        const deleteBlogPhoto = (id) => {
          confirmFunc().then(() => {
            DeletePhoto(id).then((res) => {
              if (res) {
                ElMessage.success("删除成功");
                getAllPhoto();
              }
            })
          })
        }
        const selectedRows = ref([]);
        const handlePush = (id) => {
            if (selectedRows.value.includes(id)) {
            let index = selectedRows.value.indexOf(id);
            selectedRows.value.splice(index, 1);
            } else {
            selectedRows.value.push(id);
            }
            console.log(selectedRows.value)
        }

        const deleteSomePhoto = () => {
          console.log(selectedRows.value)
            selectedRows.value.forEach(row => {
                DeletePhoto(row).then((res) => {
                  if (res.status === 200) {
                    getAllPhoto();
                  }
                })
            });
        }

        onMounted(() => {
            getAllPhoto();
        })
        const showImageViewer = ref(false);
        const handleImageHide = () => {
          showImageViewer.value = false
        }
        const selectedIndex = ref(0)
        const handelEventPhoto = () => {
             // 获取点击的图片在photoUrl列表中的索引值
            selectedIndex.value = photoUrl.value.indexOf(event.target.src)
            showImageViewer.value = true;
        }
        return{
            BlogPhoto,
            deleteBlogPhoto,
            selectedRows,
            handlePush,
            deleteSomePhoto,
            showImageViewer,
            handleImageHide,
            selectedIndex,
            handelEventPhoto,
            photoUrl
        }
    }
}
</script>
<style scoped>
.table{
    margin-top: 20px;
}
</style>

