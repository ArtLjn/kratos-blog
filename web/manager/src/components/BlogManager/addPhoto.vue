<template>
<el-button @click="show = true" type="success" plain>新增图片</el-button>
  <el-dialog v-model="show" :lock-scroll="false" class="dialog" style="width:80%;"> 
        <el-form>
            <el-form-item label="标题">
                <el-input type="text" v-model="data.data.title"></el-input>
            </el-form-item>
            <el-form-item label="定位">
                <el-input type="text" v-model="data.data.position"></el-input>
            </el-form-item>
            <el-form-item label="图片">
              <el-input type="text" v-model="data.data.photo"></el-input>
             </el-form-item>
            <el-form-item label="日期">
              <el-date-picker
                  v-model="data.data.date"
                  type="date"
                  placeholder="Pick a day"
                  format="YYYY-MM-DD"
                  value-format="YYYY-MM-DD"
                  :size="100"
              />
            </el-form-item>
            <input type="file" @change="handleMainPhoto"  ref="fileInput">
            <el-button @click="save" type="primary" plain>提交</el-button>
        </el-form>
  </el-dialog>
</template>
<script>
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import {uploadFile} from "@/components/api/blog";
import {savePhoto} from "@/components/api/photo";
export default{
    name:"addPhoto",
    setup() {
    const data = reactive({
      data:{
        title:'',
        position:'',
        photo:"",
        date:''
      }
    })
    const file2 = ref('')
    const handleMainPhoto = (event) => {
        uploadFile(event.target.files[0]).then((res) => {
          if (res) {
            data.data.photo = res.data;
            ElMessage.success("上传成功")
          }
        })
     }
     const save = () => {
      savePhoto(data).then((res) => {
        if (res) {
          ElMessage.success("保存成功")
          show.value = false
        }
      })
     }
      const show = ref(false);
        return{
            data,
            file2,
            handleMainPhoto,
            show
            ,save
        }
    }
}
</script>
<style scoped>

</style>
