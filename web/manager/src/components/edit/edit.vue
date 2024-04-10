<template>
    <div class="body">
        <div class="card">
          <el-form>
            <input placeholder="输入标题" class="input" style="width:80%;float:left;" v-model="md.data.title">
            <font-awesome-icon
              :icon="['fas', 'floppy-disk']"
              style="margin:16px 20px auto 10px;float:right;height:30px;"
              @click="saveDialog = true"
            ></font-awesome-icon>
            <el-dialog
            title="保存文章"
            v-model="saveDialog"
            width="40%">
              <el-form>
                <el-form-item label="文章标签">
                  <el-select v-model="md.data.tag" clearable placeholder="选择标签" style="float:right;margin-right:10px;">
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
                  <el-switch v-model="md.data.comment"
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
                      style="width:200px"
                      :limit="1"
                      :on-change="handleMainPhoto"
                      :auto-upload="false"
                  >
                    <div class="el-upload__text">添加文章封面</div>
                  </el-upload>
                </el-form-item>

                <el-form-item style="float: right">
                  <el-button @click="saveDialog=false">取消</el-button>
                  <el-button type="primary" @click="saveOrUpdate">保存</el-button>
                </el-form-item><br>
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

            <input placeholder="导语" class="input" style="border:none;" v-model="md.data.preface">

            <v-md-editor
            :include-level="[1,2,3,4]"
            v-model="md.data.content"
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
  import {SaveBlog, setCacheBlog, updateBlog, uploadFile} from "@/components/api/blog";
  import {AddTag, GetAllTag} from "@/components/api/tag";
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
          comment: true
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
        .then(() => {
          const cacheBlog = reactive({
            data: {
              content: md.data.content,
              title: md.data.title,
              preface: md.data.preface,
              photo: md.data.photo,
              tag: md.data.tag,
            }
          })
          setCacheBlog(cacheBlog).then((res) => {
            if (res) {
              ElMessage.success("缓存成功")
            }
          })
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
        if (!md.data.title || !md.data.preface || !md.data.content || !md.data.photo) {
          ElMessage.warning("请填写完整");
          return;
        }
        SaveBlog(md).then((res) => {
          if (res) {
            ElMessage.success("保存成功")
          }
        })
      };

      const update = () => {
        if (!md.data.title || !md.data.preface || !md.data.content || !md.data.photo) {
          ElMessage.warning("请填写完整");
          return;
        }
        const id = route.query.id;
        updateBlog(id,md).then((res) => {
          if (res) {
            ElMessage.success("更新成功");
          }
        })
      };
      
      const saveOrUpdate = () => {
        md.data.comment = md.data.comment === "1";
        if (route.query.id) {
          update();
        } else {
          save();
        }
      };

      const handleUploadImage = (event) => {
        uploadFile(event.target.files[0]).then((res) => {
          if (res) {
            ElMessage.success("上传成功");
            md.data.content += "\n![Description](" + res.data + ")";
          }
        })
     };

      const handleMainPhoto = (file) => {
        console.log(file)
        uploadFile(file.raw).then((res) => {
          if (res) {
            md.data.photo = res.data;
            ElMessage.success("上传成功");
          }
        })
      }

     const getTag = () => {
        GetAllTag().then((res) => {
          getTagList.value = res.data;
        })
      }
      onMounted(() => {
        md.data.title = route.query.title || "";
        md.data.preface = route.query.preface || "";
        md.data.content = route.query.content || "";
        md.data.tag = route.query.tag || "";
        md.data.comment = route.query.comment || "";
        md.data.photo = route.query.photo || "";
        getTag();
      });

      const addTag = () => {
        AddTag().then(() => {
          getTag();
        })
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
    min-height: 100vh;
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
  