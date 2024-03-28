<template>
  <el-container class="container">
    <el-main class="blog-main">
      <div class="div1">
        <div>
          <div style="float:left;">
            <el-form inline style="float:left" @submit.prevent="queryData()">
              <el-form-item >
                <el-input v-model="query"  placeholder="输入标题进行查询"></el-input>
              </el-form-item>
              <el-form-item>
                <el-button type="primary" @click="queryData()">
                  <font-awesome-icon :icon="['fas', 'magnifying-glass']" />
                  <span style="vertical-align: middle;margin-left:5px;"> Search </span>
                </el-button>
              </el-form-item>
            </el-form>
          </div>
          <div style="float:right;margin-bottom:20px;font-size:20px;font-weight:bold;">
            <el-radio-group v-model="commentStatus.comment"  @change="setCommentStatus">
              <el-radio label="0" size="large" border>禁止评论</el-radio>
              <el-radio label="1" size="large" border>允许评论</el-radio>
            </el-radio-group>
            <el-button @click="deleteSomeBlog" type="danger" style="margin-left:20px;" >批量删除</el-button>
          </div>
        </div>
        <div class="table-container">
          <el-table :data="Blogs" class="custom-table"  stripe :lazy="true" border height="400">
            <el-table-column type="selection" label="序号" width="80" >
              <template #default="{ row }">
                <el-checkbox v-model="row.selected" @change="handlePush(row.id)"></el-checkbox>
              </template>
            </el-table-column>
            <el-table-column label="标题" width="300">
              <template #default="{ row }">
                <el-tooltip :content="`${row.id}`" placement="top" >
                <router-link :to="`/main/watch/${row.id}`" class="link"><h1>{{ row.title }}</h1></router-link></el-tooltip>
              </template>
            </el-table-column>
            <el-table-column label="评论" width="100" prop="comment"></el-table-column>
            <el-table-column label="时间" width="120" prop="createTime"></el-table-column>
            <el-table-column label="标签" width="150" prop="tag"></el-table-column>
            <el-table-column label="浏览量" width="100" prop="visits"></el-table-column>
            <el-table-column label="浏览权限" width="100" prop="appear"></el-table-column>
            <el-table-column label="最近一次更新时间" width="150" prop="updateTime"></el-table-column>
            <el-table-column label="操作" width="300">
              <template #default="{row}">
                <el-button @click="handelNote(row.id)">编辑</el-button>
                <el-button @click="removeBlog(row.id)" type="danger">删除</el-button>
                <el-popconfirm
                confirm-button-text='允许'
                cancel-button-text='禁止'
                :icon="el-icon-info"
                icon-color="#626AEF"
                title="是否允许评论"
                @confirm="sendRequest(row.id,1)"
                @cancel="sendRequest(row.id,2)"
              >
                <template #reference>
                  <el-button type="warning">评论</el-button>
                </template>
              </el-popconfirm>
              <el-popconfirm
              confirm-button-text='允许'
              cancel-button-text='禁止'
              :icon="el-icon-info"
              icon-color="#626AEF"
              title="是否允许访问"
              @confirm="sendChangeAppear(row.id,1)"
              @cancel="sendChangeAppear(row.id,2)"
            >
            <template #reference>
              <el-button type="primary">访问</el-button>
            </template>
            </el-popconfirm>
              </template>
            </el-table-column>
          </el-table>
      </div> <br>
    </div>
       <div>
            <div style="margin-bottom:20px;">
              <el-select placeholder="选择文章" v-model="suggestBlogForm.id" filterable>
                <el-option v-for="item in Blogs" 
                :key="item.id" :label="item.title"
                :value="item.id">
                </el-option>
              </el-select>
              <el-button type="primary" style="margin-left:20px;" @click="addBlogForSuggest">
                确定
              </el-button>
            </div>
            <el-table :data="SuggestForm" class="custom-table" border style="width:600px;">
              <el-table-column label="标题" width="400">
                <template #default="{ row }">
                  <router-link :to="`/main/watch/${row.id}`" class="link">{{ row.title }}</router-link>
                </template>
              </el-table-column>
              <el-table-column width="200" label="操作">
                <template #default="{row}">
                  <el-button  type="danger" @click="deleteSuggest(row.id)">删除推荐</el-button>
                </template>
              </el-table-column>
            </el-table>
      </div>
    </el-main>
  </el-container>
  </template>
  
  <script>
  import { computed, onMounted, reactive, ref } from 'vue';
  import axios from 'axios';
  import { ElMessage } from 'element-plus';
  import { useRouter } from 'vue-router';
  import login from '../authentication/login.vue';
  import {
    addSuggestBlog,
    DeleteBlog,
    GetAllBlog,
    getAllSuggestBlog,
    GetBlogById,
    SearchBlog,
    updateBlogOnly
  } from "@/components/api/blog";
  import {SUCCESS_REQUEST} from "@/components/api/status";

  export default {
    // eslint-disable-next-line vue/multi-word-component-names
    name: 'Blog',
    components: {
      // eslint-disable-next-line vue/no-unused-components
      login
    },
    data() {
      return{
        updateOnlyForm:{
          raw:Comment,
          id:0,
          res:false
        }
      }
    },
    setup() {//分页
      const itemsPerPage = 10;
      const currentPage = ref(1);
      const displayedItems = computed(() => {
        const startIndex = (currentPage.value - 1) * itemsPerPage;
        const endIndex = startIndex + itemsPerPage;
        return Blogs.value.slice(startIndex, endIndex);
      });
      const handleCurrentChange = (page) => {
        currentPage.value = page;
      };
      const Blogs = ref([]);
      const router = useRouter();
      const getAllBlog =  () => {
        GetAllBlog().then((res) => {
           if (res.common.code === SUCCESS_REQUEST) {
            Blogs.value = res.List;
          }
        })
      };

      const query = ref('');
      const queryData = () => {
        if(query.value == null) {
          return;
        }
        SearchBlog(query.value).then((res) => {
          Blogs.value = res.list;
        })
      }
      const suggestBlogForm = ref({
        id:""
      })
  
      const removeBlog = (id) => {
          DeleteBlog(id).then((res) => {
            if (res.common.code === SUCCESS_REQUEST) {
              getAllBlog();
              ElMessage.success('删除成功');
            } else {
              ElMessage.error(res.error)
            }
        })
      }; 
      
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

      const deleteSomeBlog = () => {
        selectedRows.value.forEach(row => {
            DeleteBlog(row).then((res) => {
              console.log(res)
            })
        })
      }
      const handelNote = (id) => {
        GetBlogById(id).then((response) => {
          const data = response.data.data;
          const query = {
            id: id,
            title: data.title,
            content: data.content,
            preface: data.preface,
            tag:data.tag,
            comment:data.comment,
            photo:data.photo
          };
          router.push({
            path: '/main/edit',
            query: query
          });
        })
      };
      onMounted(() => {
        getAllBlog();
        searchAllSuggest();
      });
      const SuggestForm = ref([]);
      const addBlogForSuggest = () => {
        addSuggestBlog(suggestBlogForm.value).then((res) => {
          if (res.common.code == SUCCESS_REQUEST) {
            ElMessage.success("添加成功")
            searchAllSuggest();
          } else {
            ElMessage.error(res.common.result)
          }
        })
      }
      const searchAllSuggest = () => {
        getAllSuggestBlog().then((res) => {
          if (res.common.code == SUCCESS_REQUEST) {
            SuggestForm.value = res.List
          } else {
            ElMessage.error(res.common.result)
          }
        })
      }
      const deleteSuggest = (id) => {
            axios
            .delete(`/api/deleteSuggest/${id}`)
            .then(() => {
              ElMessage.success("ok")
              searchAllSuggest();
            })
            .catch((error) => {console.log(error)})
      }

      const sendRequest = (id,comment) => {
        axios
          .put(`/api/updateAllow?id=${id}&comment=${comment}`)
          .then(() => {
            ElMessage.success("修改成功")
            getAllBlog();
          })
          .catch((error) => {
            console.log(error)
          })
      }
      const sendChangeAppear = (id,appear) => {
        axios
          .put(`/api/updateAppear?id=${id}&appear=${appear}`)
          .then(() => {
            ElMessage.success("修改成功")
            getAllBlog();
          })
          .catch((error) => {
            console.log(error)
          })
      }
      const getNoteStatus = (id) => {
        axios
          .get(`/api/getNoteStatus?id=${id}`)
          .then(response => {
            const status = response.data;
            return status;
          })
          .catch(error => {
            console.log(error);
          });
      }
      const commentStatus = reactive({
        comment:"5"
      });
      const setCommentStatus =()=> {
        updateAllBlogCommentStatus(commentStatus.comment).then((res) => {
          if (res.status === 200) {
            getAllBlog()
            ElMessage.success(res.result)
          } else {
            ElMessage.error(res.error)
          }
        })
      }
      
      return {
        Blogs,
        removeBlog,
        handelNote,
        selectedRows,
        handlePush,
        deleteSomeBlog,
        suggestBlogForm,
        addBlogForSuggest,
        SuggestForm,
        deleteSuggest,
        displayedItems,
        handleCurrentChange,
        sendRequest,
        getNoteStatus,
        setCommentStatus,
        commentStatus,
        query,
        queryData,
        sendChangeAppear
      };
    }
  }
  </script>
  
  <style scoped>
  @import url('../../assets/css/main.css');
  
  .blog-main {
    padding: 20px;
    border-radius: 4px;
    width: 1500px;
    margin: auto;
    height: auto;
  }
  
  .custom-table {
    width: 100%;
  }
  
  .handle {
    margin-right: 20px;
  }

  /* 在视图宽度小于600px时，将博客主要容器的宽度调整为100% */
  @media screen and (max-width: 600px) {

  }
  .page{
    margin: 20px auto 20px auto;
    width: calc(100%);
    display: flex;
    justify-content: center;
  }
</style>
  