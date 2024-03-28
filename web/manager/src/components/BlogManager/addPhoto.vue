<template>
<el-button @click="isokOrNo = true" type="success" plain>新增图片</el-button>
  <el-dialog v-model="isokOrNo" :lock-scroll="false" class="dialog" style="width:80%;"> 
        <el-form>
            <el-form-item label="标题">
                <el-input type="text" v-model="data.title"></el-input>
            </el-form-item>
            <el-form-item label="定位">
                <el-input type="text" v-model="data.position"></el-input>
            </el-form-item>
            <el-form-item label="图片">
              <el-input type="text" v-model="data.photo"></el-input>
          </el-form-item>
            <input type="file" @change="handleMainPhoto"  ref="fileInput">
            <el-button @click="save" type="primary" plain>提交</el-button>
        </el-form>
  </el-dialog>
</template>
<script>
import { reactive, ref } from 'vue'
import axios from 'axios'
import { ElMessage } from 'element-plus'
import {uploadFile} from "@/components/api/blog";
export default{
    name:"addPhoto",
    setup() {
    const data = reactive({
        title:'',
        position:'',
        photo:""
    })
    const file2 = ref('')
    const handleMainPhoto = (event) => {
        uploadFile(event).then((res) => {
          data.photo = res.result;
        })
     }
     const save = () => {
      axios
        .post("/api/addPhoto",data)
        .then((res) => {
          ElMessage.success(res.data.message)
        }).catch((err) => {
          ElMessage.error(err.res.data.error)
        })
     }
      const isokOrNo = ref(false);
        return{
            data,
            file2,
            handleMainPhoto,
            isokOrNo
            ,save
        }
    }
}
</script>
<style scoped>

</style>
