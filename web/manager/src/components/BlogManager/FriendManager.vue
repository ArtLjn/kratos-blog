<template>
    <el-button type="success" plain @click="addFalse">新增友链</el-button>
    <el-dialog v-model="isFalse" title="添加友链">
        <el-form>
            <el-form-item label="name">
                <el-input v-model="addForm.title"></el-input>
            </el-form-item>
            <el-form-item label="preface">
                <el-input v-model="addForm.preface"></el-input>
            </el-form-item>
            <el-form-item label="url">
                <el-input v-model="addForm.url" ></el-input>
            </el-form-item>
            <el-form-item label="图片">
                <el-input type="text" v-model="addForm.photo"></el-input>
            </el-form-item>
            <input type="file" @change="handleMainPhoto"  ref="fileInput">
            <el-button @click="addFriend" type="primary" plain>保存</el-button>
        </el-form>
    </el-dialog>
    <el-table :data="Friends" border style="margin-top:20px;">
        <el-table-column label="name">
            <template #default="{row}">
                <span>{{row.title}}</span>
            </template>
        </el-table-column>
        <el-table-column label="preface">
            <template #default="{row}">
                <span>{{row.preface}}</span>
            </template>
        </el-table-column>
        <el-table-column label="url">
            <template #default="{row}">
                <a :href="row.url">查看</a>
            </template>
        </el-table-column>
        <el-table-column label="操作">
            <template #default="{row}">
                <el-button type="danger" plain @click="deleteFriend(row.id)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>
<script>
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import axios from 'axios'
import {safeMath} from "@/components/api/util";
import {DeleteFriend} from "@/components/api/friend";
export default{
    name:"FriendManager",
    setup() {
        const file = ref('');
        const addForm = reactive({
            title:'',
            preface:'',
            url:'',
            photo:''
        })

        const addFriend = async() => {
            if(!addForm.title || !addForm.preface || !addForm.url) {
                ElMessage.warning("no");
                return;
            }
            axios
                .post("/api/addFriend",addForm)
                .then((res) => {
                  if (res.data.status === 200) {
                    ElMessage.success(res.data.result)
                    getAllFriend()
                  } else {
                    ElMessage.error(res.data.error)
                  }
                }).catch((err) => {
                    ElMessage.error(err.res.data.err)
                })
        }
        const handleMainPhoto = (event) => {
            file.value = event.target.files[0];
            const formData = new FormData();
            formData.append("file",file.value)
            axios
                .post('/api/upload', formData,{
                    headers: {
                        'Content-Type': 'multipart/form-data'
                    }
                })
                .then((response) => {
                    console.log(response.data.result)
                    addForm.photo = response.data.result;
                })
                .catch((error) => {
                    console.error(error);
                });
        }
        const Friends = ref([]);
        const getAllFriend = () => {
            axios
                .get("/api/getAllFriends")
                .then((response) => {
                    Friends.value = response.data.list;
                })
        }
        const oka = reactive({
            key:""
        });
        const deleteFriend = (id) => {
          safeMath().then((r) => {
            DeleteFriend(id,r).then((res) => {
              if (res.status === 200) {
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
            Friends,oka,deleteFriend
            ,isFalse,addFalse,file,handleMainPhoto
        }
    }
}
</script>
