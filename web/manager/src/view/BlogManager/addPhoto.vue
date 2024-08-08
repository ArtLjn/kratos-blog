<template>
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
          <el-button @click="show = false" >取消</el-button>
          <el-button @click="save" type="primary">提交</el-button>
        </div><br>
  </el-dialog>
</template>
<script>
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
import {savePhoto} from "@/view/api/photo";
import {UploadFile} from "@/view/api/tool";
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
    const handleMainPhoto = (param) => {
      UploadFile(param).then(res => {
        ElMessage.success(res.data.info);
        data.data.photo = res.data.data;
      })
     }
     const save = () => {
      savePhoto(data).then(() => {
        ElMessage.success("保存成功")
        show.value = false
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
