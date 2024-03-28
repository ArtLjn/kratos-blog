<template>
    <div class="body">
        <div class="card">
          <el-form>
            <input placeholder="输入标题" class="input" style="width:80%;float:left;" v-model="md.title">
            <font-awesome-icon
              :icon="['fas', 'floppy-disk']"
              style="margin:16px 20px auto 10px;float:right;height:30px;"
              @click="saveDialog = true"
            ></font-awesome-icon>
            <el-button @click="saveDialog = true">Click</el-button>

            <el-dialog
            title="保存文章"
            v-model="saveDialog"
            width="40%">
              <el-form>
                <el-form-item label="文章标签">
                  <el-select v-model="md.tag" clearable placeholder="选择标签" style="float:right;margin-right:10px;">
                    <el-option
                        v-for="item in getTagList"
                        :key="item.id"
                        :label="item.tagName"
                        :value="item.tagName"
                    />
                  </el-select>
                  <el-button type="primary" @click="addTag" plain>添加标签</el-button>
                </el-form-item>

                <el-form-item label="开启评论">
                  <el-switch v-model="md.comment"
                             inline-prompt
                             active-text="open"
                             inactive-text="off"
                             active-value="1"
                             inactive-value="0"
                             inactive-color="green"
                  ></el-switch>
                </el-form-item>
                <el-form-item label="添加封面">
                  <el-upload
                      class="upload-demo"
                      drag
                      multiple
                      style="width:200px"
                      :on-success="handleMainPhoto"
                  >
                    <i class="el-icon-upload"></i>
                    <div class="el-upload__text">添加文章封面</div>
                  </el-upload>
                </el-form-item>
              </el-form>
            </el-dialog>

              <el-tooltip content="暂时保存" placement="top">
                <font-awesome-icon 
                style="margin:16px 20px auto 10px;float:right;height:30px;"
                :icon="['fas', 'cloud-arrow-up']" 
                @click="SaveTemporarily"/>
             </el-tooltip>

             <el-tooltip content="继续编辑" placement="top">
              <font-awesome-icon :icon="['fas', 'pen-to-square']"
              style="margin:16px 20px auto 10px;float:right;height:30px;"
              @click="continueEdit"
              />
             </el-tooltip>

            <input placeholder="导语" class="input" style="border:none;" v-model="md.preface">
            <div style="margin-bottom:20px;">

            </div>

            <el-input v-model="md.photo" class="input" style="width:300px;" placeholder="输入图片链接"></el-input>
            <v-md-editor
            :include-level="[1,2,3,4]"
            v-model="md.content"
            class="markdown-editor"
            :disabled-menus="[]"
            @upload-image="handleUploadImage"
          ></v-md-editor>
          </el-form>
        </div>
    </div>
  
  </template>
  
  <script>
  import { onMounted, reactive ,ref} from "vue";
  import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
  import { fas } from "@fortawesome/free-solid-svg-icons";
  import { library } from "@fortawesome/fontawesome-svg-core";
  import { ElMessage, ElMessageBox } from "element-plus";
  import {useRoute} from "vue-router";
  import axios from 'axios';
  import {setCacheBlog, uploadFile} from "@/components/api/blog";
  import {AddTag, GetAllTag} from "@/components/api/tag";
  import {SUCCESS_REQUEST} from "@/components/api/status";
  library.add(fas);
  
  export default {
    methods: {AddTag},
    components: {
      FontAwesomeIcon,
    },
    setup() {
      const saveDialog = ref(false);
      const getTagList = ref([]);
      const route = useRoute();
      let md = reactive({
        data:{
          title: "",
          preface: "",
          content: "Hello Markdown",
          photo:'',
          tag:'',
        }
      });
      const status = ref(['0','1'])
      const SaveTemporarily = () =>{
        ElMessageBox.confirm(
          '确认要缓存笔记吗?',
          'Warning',
          {
            confirmButtonText:'OK',
            cancelButtonText:'Cancel',
            type:'warning',
          }
        )
        .then(async () => {
          const res = await setCacheBlog(md)
          if (res.common.code == SUCCESS_REQUEST) {
            ElMessage.success("缓存成功")
          } else {
            ElMessage.warning("缓存失败")
          }
        })
        .catch((error) => {
          console.log(error)
        })
      }
      const continueEdit = () => {
        const data = localStorage.getItem("缓存笔记");
        const parsedData = JSON.parse(data);
        md.data = parsedData
      }
      const save = () => {
        localStorage.removeItem("缓存笔记")
        if (!md.title || !md.preface || !md.content || !md.photo) {
          ElMessage.warning("请填写完整");
          return;
        }
          axios.post("/api/addBlog", md)
           .then((response) => {
            console.log(response.data)
            ElMessage.success("保存成功")
            sessionStorage.removeItem("blogs");
          })
            .catch((error) => {
            console.log(error)
          })
      };

      const update = () => {
        if (!md.title || !md.preface || !md.content || !md.photo) {
          ElMessage.warning("请填写完整");
          return;
        }
        const id = route.query.id;
          axios.put(`/api/updateBlog/${id}`,md).then(() => {
            ElMessage.success("更新成功");
            sessionStorage.removeItem("blogs");
          }).catch((error) =>{
              console.log(error)
              ElMessage.error("密码错误")
            }
           )
      };
      
      const saveOrUpdate = () => {
        if (route.query.id) {
          update();
        } else {
          save();
        }
      };

      const handleUploadImage = (event) => {
        uploadFile(event).then((res) => {
          if (res.code === 200) {
            ElMessage.success("上传成功");
            md.content += "\n![Description](" + res.data + ")";
          }
        })
     };
     const handleMainPhoto = (event) => {
        uploadFile(event).then((res) => {
          if (res.code === 200) {
            md.photo = res.data;
            ElMessage.success("上传成功")
          }
        })
     }
     const getTag = () => {
        GetAllTag().then((res) => {
          getTagList.value = res.data;
        })
      }
      onMounted(() => {
        md.title = route.query.title || "";
        md.preface = route.query.preface || "";
        md.content = route.query.content || "";
        md.tag = route.query.tag || "";
        md.comment = route.query.comment || "";
        md.photo = route.query.photo || "";
        getTag();
      });

      const addTag = () => {
        AddTag().then(() => {
          getTag();
        });
      }
      return {
        md,
        saveOrUpdate,
        handleUploadImage,
        handleMainPhoto,
        getTagList,
        SaveTemporarily,
        continueEdit,
        status,
        saveDialog,
        addTag
      };
    },
  };
  </script>

  
  <style scoped>
  @import url("../../assets/css/main.css");
  .main {
    width: 100%;
    height: 350px;
  }
  
  .card {
    margin:auto;
    height: auto;
    width: 100%;
    background-color: rgb(44, 18, 18,0);
    }
  
  .input {
    width: 100%;
    height: 40px;
    margin-bottom: 10px;
    border-block-start: none;
    border-left: none;
    border-right: none;
    outline: none;
    border-color: rgb(18, 78, 18);
  }
  @media screen and (max-width: 600px) {
   .card{
      width: 100%;
   } 
  }
  </style>
  