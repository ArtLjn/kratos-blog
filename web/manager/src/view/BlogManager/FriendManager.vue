<template>
<el-container>
  <el-main>
    <el-button type="success" plain @click="addFalse">新增友链</el-button>
    <el-dialog v-model="isFalse" title="添加友链">
      <el-form>
        <el-form-item label="name">
          <el-input v-model="addForm.data.title"></el-input>
        </el-form-item>
        <el-form-item label="preface">
          <el-input v-model="addForm.data.preface"></el-input>
        </el-form-item>
        <el-form-item label="url">
          <el-input v-model="addForm.data.url" ></el-input>
        </el-form-item>
        <el-form-item label="图片">
          <el-input type="text" v-model="addForm.data.photo"></el-input>
        </el-form-item>
        <input type="file" @change="handleMainPhoto"  ref="fileInput">
        <el-button @click="addFriend" type="primary" plain>保存</el-button>
      </el-form>
    </el-dialog>
    <el-table :data="Friends" border style="margin-top:20px;">
      <el-table-column label="name" prop="title"></el-table-column>
      <el-table-column label="preface" prop="preface"></el-table-column>
      <el-table-column label="url">
        <template #default="{row}">
          <a class="link" :href="row.url" target="_blank">查看</a>
        </template>
      </el-table-column>
      <el-table-column label="操作">
        <template #default="{row}">
          <el-button type="danger" plain @click="deleteFriend(row.id)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </el-main>
</el-container>
</template>
<script>
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import {AddFriend, DeleteFriend, GetAllFriend} from "@/view/api/friend";
import {uploadFile} from "@/view/api/blog";
import {confirmFunc} from "@/view/api/util";
export default{
    name:"FriendManager",
    setup() {
        const file = ref('');
        const addForm = reactive({
          data:{
            title:'',
            preface:'',
            url:'',
            photo:''
          }
        })

        const addFriend = async() => {
            if(!addForm.data.title || !addForm.data.preface || !addForm.data.url) {
                ElMessage.warning("no");
                return;
            }
            AddFriend(addForm).then((res) => {
              if (res) {
                ElMessage.success(res.common.result)
                getAllFriend()
              }
            })
        }
        const handleMainPhoto = (event) => {
          uploadFile(event.target.files[0]).then((res) => {
            if (res) {
              addForm.data.photo = res.data;
              ElMessage.success("上传成功")
            }
          })
        }
        const Friends = ref([]);
        const getAllFriend = () => {
          GetAllFriend().then((res) => {
            if (res) {
              Friends.value = res.data;
            }
          });
        }
        const deleteFriend = (id) => {
          confirmFunc().then(() => {
            DeleteFriend(id).then((res) => {
              if (res) {
                getAllFriend()
                ElMessage.success("删除成功")
              }
            })
          })
        }
        const isFalse = ref(false);
        const addFalse = () => {
            isFalse.value = true;
        }
        onMounted(() => {
            getAllFriend();
        })
        return{
            addForm,
            addFriend,
            Friends,deleteFriend
            ,isFalse,addFalse,file,handleMainPhoto
        }
    }
}
</script>
