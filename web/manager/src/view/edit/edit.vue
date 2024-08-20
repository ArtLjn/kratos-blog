<template>
  <el-container>
    <el-main>
        <el-form>
          <input placeholder="输入标题" class="input" style="width:60%;float:left;border-radius: 5px;padding-left: 20px;" v-model="md.data.title">
          <el-dialog
          title="保存文章"
          v-model="saveDialog"
          style="border-radius: 20px;"
          width="40%">
            <el-descriptions border column="1">
              <el-descriptions-item label="文章标签">
                <el-select v-model="md.data.tag" filterable clearable placeholder="选择标签" style="margin-right:10px;">
                  <el-option
                      v-for="item in getTagList"
                      :key="item.id"
                      :label="item.tagName"
                      :value="item.tagName"
                  />
                </el-select>
                <el-button type="primary" @click="addTag" plain>添加标签</el-button>
              </el-descriptions-item>
              <el-descriptions-item label="开启评论">
                <el-switch v-model="md.data.comment"
                           inline-prompt
                           active-text="open"
                           inactive-text="off"
                           active-value="1"
                           inactive-value="0"
                           inactive-color="green"
                ></el-switch>
              </el-descriptions-item>
              <el-descriptions-item label="添加封面">
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
            </el-descriptions><br>
            <div style="float:right;">
              <el-button @click="saveDialog=false">取消</el-button>
              <el-button type="primary" @click="saveOrUpdate">保存</el-button>
            </div><br>
          </el-dialog>

          <div>
            <el-tooltip content="保存" placement="top">
              <font-awesome-icon
                  :icon="['fas', 'floppy-disk']"
                  style="margin:16px 20px auto 10px;float:right;height:30px;"
                  @click="saveDialog = true"
              ></font-awesome-icon>
            </el-tooltip>
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
          </div>


          <input placeholder="导语" class="input" style="border:none;border-radius: 5px;padding-left: 20px;width: 70%;" v-model="md.data.preface">

          <v-md-editor
          :include-level="[1,2,3,4]"
          v-model="md.data.content"
          class="markdown-editor"
          :disabled-menus="[]"
          @upload-image="handleUploadImage"
        ></v-md-editor>

          </el-form>
      </el-main>
    </el-container>
  </template>

  <script>
  import { onMounted, reactive ,ref} from "vue";
  import { FontAwesomeIcon } from "@fortawesome/vue-fontawesome";
  import { fas } from "@fortawesome/free-solid-svg-icons";
  import { library } from "@fortawesome/fontawesome-svg-core";
  import { ElMessage, ElMessageBox } from "element-plus";
  import {useRoute} from "vue-router";
  import {SaveBlog, setCacheBlog, updateBlog} from "@/view/api/blog";
  import {AddTag, GetAllTag} from "@/view/api/tag";
  import {UploadFile} from "@/view/api/tool";
  import axios from "axios";
  import {addTagMath} from "@/view/api/util";
  library.add(fas);
  
  export default {
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
        const formData = new FormData();
        formData.append('file',event.target.files[0]);
        axios.post('/tool/upload',formData).then(res => {
          if (res.data.code === 200) {
            ElMessage.success("上传成功");
            md.data.content += "\n![Description](" + res.data.data + ")";
          }
        })
     };

      const handleMainPhoto = (param) => {
        UploadFile(param).then(res => {
          ElMessage.success(res.info);
          md.data.photo = res.data;
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
        addTagMath().then(name => {
          AddTag(name).then(() => {
            ElMessage.success("添加成功");
            getTag();
          })
        }).catch(() => {
          ElMessage.error("标签名不能为空");
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
  