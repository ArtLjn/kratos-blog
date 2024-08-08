<template>
  <el-container class="container">
    <el-main class="blog-main">
      <el-tabs v-model="activeName" class="demo-tabs" @tab-click="handleClick">
        <el-tab-pane label="Blog"  name="first">
          <div class="div1">
            <div>
              <div style="float:left;">
                <el-form inline style="float:left" @submit.prevent="queryBlog()">
                  <el-form-item >
                    <el-input v-model="query"  placeholder="输入标题进行查询"></el-input>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="primary" @click="queryBlog()">
                      <font-awesome-icon :icon="['fas', 'magnifying-glass']" />
                      <span style="vertical-align: middle;margin-left:5px;"> Search </span>
                    </el-button>
                  </el-form-item>
                  <el-form-item>
                    <el-button type="warning" @click="ExportBlog()">导出MD</el-button>
                  </el-form-item>
                </el-form>
              </div>
              <div style="float:right;margin-bottom:20px;font-size:20px;font-weight:bold;">
                <el-popconfirm
                    confirm-button-text='允许'
                    cancel-button-text='禁止'
                    :icon="el-icon-info"
                    icon-color="#626AEF"
                    title="是否允许评论"
                    @confirm="updateIndividual(Comment,true)"
                    @cancel="updateIndividual(Comment,false)"
                >
                  <template #reference>
                    <el-button>设置评论</el-button>
                  </template>
                </el-popconfirm>
                <el-popconfirm
                    confirm-button-text='允许'
                    cancel-button-text='禁止'
                    :icon="el-icon-info"
                    icon-color="#626AEF"
                    title="是否允许访问"
                    @confirm="updateIndividual(Appear,true)"
                    @cancel="updateIndividual(Appear,false)"
                >
                  <template #reference>
                    <el-button>设置访问</el-button>
                  </template>
                </el-popconfirm>
              </div>
            </div>
            <div class="table-container">
              <el-table :data="Blogs" class="custom-table"  stripe :lazy="true" border height="400">
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
                <el-table-column label="操作" width="350" fixed="right">
                  <template #default="{row}">
                    <el-button @click="handleNote(row.id)">编辑</el-button>
                    <el-button @click="removeBlog(row.id)" type="danger">删除</el-button>
                    <el-popconfirm
                        confirm-button-text='允许'
                        cancel-button-text='禁止'
                        :icon="el-icon-info"
                        icon-color="#626AEF"
                        title="是否允许评论"
                        @confirm="updateBlogOnly(row.id,Comment,true)"
                        @cancel="updateBlogOnly(row.id,Comment,false)"
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
                        @confirm="updateBlogOnly(row.id,Appear,true)"
                        @cancel="updateBlogOnly(row.id,Appear,false)"
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
              <el-button type="primary" style="margin-left:20px;" @click="addSuggestBlog">
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
                <template #default="{$index}">
                  <el-button  type="danger" @click="removeSuggest($index)">删除推荐</el-button>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-tab-pane>
        <el-tab-pane label="cacheBlog" name="second">
          <el-table :data="cacheBlogs" class="custom-table"  stripe :lazy="true" border height="400">
            <el-table-column label="标题" width="300" prop="title"></el-table-column>
            <el-table-column label="图片" width="100" >
              <template #default="{ row }">
                <el-tooltip :content="row.photo" placement="top">
                  <a :href="row.photo" class="link" target="_blank">url</a>
                </el-tooltip>
              </template>
            </el-table-column>
            <el-table-column label="标签" width="100" prop="tag"></el-table-column>
            <el-table-column label="简介"  width="300" prop="preface"></el-table-column>
            <el-table-column label="操作" width="200" fixed="right">
              <template #default="{$index}">
                <el-button @click="handleCacheBlog($index)">编辑</el-button>
                <el-button @click="removeCacheBlog($index)" type="danger">删除</el-button>
              </template>
            </el-table-column>
          </el-table>
        </el-tab-pane>
      </el-tabs>
    </el-main>
  </el-container>
  </template>

  <script setup>
  import {onMounted, reactive, ref, watch} from "vue";
  import {
    AddSuggestBlog,
    DeleteBlog, delSuggestBlog,
    GetAllBlog,
    getAllSuggestBlog,
    GetBlogById,
    SearchBlog, UpdateBlogOnly,
    Comment, Appear, UpdateIndividual, GetAllCacheBlog
  } from "@/view/api/blog";
  import router from "@/router";
  import {ElMessage} from "element-plus";
  import axios from "axios";
  import {confirmFunc} from "@/view/api/util";
  import {SUCCESS_REQUEST} from "@/view/api/status";
  import {ExportBlog} from "@/view/api/tool";
  const updateOnlyForm = reactive({
    raw: Comment,
    id:0,
    res:false
  })
  const updateIndividualForm = reactive({
    raw: Comment,
    status:false
  })
  const Blogs = ref([]);
  const BlogsCopy = ref([]);
  const suggestBlogForm = reactive({
    id:"",
  });
  const query = ref("")
  const SuggestForm = ref([]);

  onMounted(() => {
    getAllBlog();
    getAllSuggest();
    getAllCacheBlog();
  })
  const getAllBlog = () => {
    GetAllBlog().then((res) => {
      Blogs.value = res.List;
      BlogsCopy.value = res.List;
    });
  }
  const queryBlog = () => {
    if (query.value === "") {
      ElMessage.warning("请输入标题");
      return;
    }
    SearchBlog(query.value).then((res) => {
      Blogs.value = res.data;
    });
  }

  const removeBlog = (id) => {
    confirmFunc().then(() => {
      DeleteBlog(id).then(() => {
        getAllBlog();
        ElMessage.success("删除成功");
      })
    })
  }
  const handleNote = (id) => {
    GetBlogById(id).then((response) => {
      const data = response.data;
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
  }

  const handleCacheBlog = (index) => {
    const data = cacheBlogs.value[index];
    const query = {
      id: data.id,
      title: data.title,
      content: data.content,
      preface: data.preface,
      tag:data.tag,
      photo:data.photo
    };
    router.push({
      path: '/main/edit',
      query: query
    });
  }
  const getAllSuggest = () => {
    getAllSuggestBlog().then((res) => {
      SuggestForm.value = res.List;
    })
  }
  const addSuggestBlog = () => {
    AddSuggestBlog(suggestBlogForm).then(() => {
      getAllSuggest();
    })
  }
  const removeSuggest = (id) => {
    confirmFunc().then(() => {
      delSuggestBlog(id).then(() => {
         getAllSuggest();
      })
    });
  }

  const updateBlogOnly = (id,raw,res) => {
    updateOnlyForm.id = id;
    updateOnlyForm.raw = raw;
    updateOnlyForm.res = res;
    console.log(updateOnlyForm)
    UpdateBlogOnly(updateOnlyForm).then((res) => {
      if (res) {
        getAllBlog();
        ElMessage.success("更新成功")
      }
    })
  }

  const updateIndividual = (raw,status) => {
    updateIndividualForm.raw = raw;
    updateIndividualForm.status = status;
    UpdateIndividual(updateIndividualForm).then(res => {
      if (res) {
        getAllBlog();
        ElMessage.success("更新成功")
      }
    })
  }

  const activeName = ref('first')
  const handleClick = (tab, event) => {
    console.log(tab, event)
  }

  const cacheBlogs = ref([]);
  const getAllCacheBlog = () => {
    GetAllCacheBlog().then((res) => {
      if (res) {
        cacheBlogs.value = res.List;
      }
    });
  }

  const removeCacheBlog = (id) => {
    confirmFunc().then(() => {
    axios.delete(`/api/deleteCacheBlog/${id}`,{
      headers: {
        'Content-Type': 'application/json',
        token: localStorage.getItem("token")
      }
    }).then((res) => {
      if (res.data.common.code === SUCCESS_REQUEST) {
        getAllCacheBlog();
          ElMessage.success("删除成功");
      } else {
         ElMessage.error("删除失败");
      }
    })});
  }

  watch(() => query.value, (newValue) => {
      const keywords = newValue.toLowerCase().split(' '); // 拆分用户输入的关键字并转换为小写
      Blogs.value = BlogsCopy.value.map(key => {
      let matches = 0;
      const regex = new RegExp(keywords.join('|'), 'gi'); // 构建正则表达式，匹配所有关键字，忽略大小写
      const matchContent = `${key.title} ${key.id} ${key.tag}`.toLowerCase(); // 将博客标题、内容、前言拼接并转换为小写
      matchContent.replace(regex, () => {
        matches++;
        return ''; // 使用空字符串替换匹配到的内容
      });
      return { key, matches };
    }).filter(item => item.matches > 0)
      .sort((a, b) => b.matches - a.matches) // 根据匹配度降序排序
      .map(item => item.key); // 仅保留博客对象
      if (newValue === '' || newValue === null) {
        Blogs.value = BlogsCopy.value;
      }
  })


  </script>

  <style scoped>
  @import url('../../assets/css/main.css');
  
  .blog-main {
    padding: 20px;
    border-radius: 4px;
    width: 1500px;
    margin: auto;
    height: auto;
    min-height: 100vh;
  }
  
  .custom-table {
    width: 100%;
  }
  .link:hover {
    color: #626AEF;
    text-decoration: underline;
  }
</style>
  