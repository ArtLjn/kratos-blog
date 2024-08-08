<template>
  <el-container>
    <el-main>
      <el-button type="success" @click="addFalse">新增友链</el-button>
      <el-dialog title="添加友联" v-model="isFalse" :lock-scroll="false" class="dialog" style="border-radius: 20px;" width="40%">
        <el-descriptions border column="1">
          <el-descriptions-item label="name">
            <el-input v-model="addForm.data.title"></el-input>
          </el-descriptions-item>
          <el-descriptions-item label="preface">
            <el-input v-model="addForm.data.preface"></el-input>
          </el-descriptions-item>
          <el-descriptions-item label="url">
            <el-input v-model="addForm.data.url"></el-input>
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
        </el-descriptions><br>
        <div style="float: right">
          <el-button @click="isFalse = false">取消</el-button>
          <el-button @click="addFriend" type="primary">保存</el-button>
        </div><br>
      </el-dialog>

      <el-row :gutter="20" class="friend-grid">
        <el-col v-for="friend in Friends" :key="friend.id" :span="8">
          <el-card :body-style="{ padding: '0px' }" class="friend-card" @dblclick="openOptions(friend)">
            <div :style="{ backgroundImage: `url(${friend.photo})` }" class="card-background">
              <div class="card-content">
                <h3>{{ friend.title }}</h3>
                <p>{{ friend.preface }}</p>
              </div>
            </div>
          </el-card>
        </el-col>
      </el-row>

      <el-dialog title="选项" v-model="optionsVisible" style="border-radius: 20px;" width="30%">
        <p>请选择操作：</p>
        <div class="actions">
          <el-button type="primary" @click="openLink(selectedFriend.url)">查看</el-button>
          <el-button type="danger" @click="deleteFriend(selectedFriend.id)">删除</el-button>
        </div>
      </el-dialog>
    </el-main>
  </el-container>
</template>

<script>
import { onMounted, reactive, ref } from 'vue';
import { ElMessage } from 'element-plus';
import { AddFriend, DeleteFriend, GetAllFriend } from "@/view/api/friend";
import { confirmFunc } from "@/view/api/util";
import { UploadFile } from "@/view/api/tool";

export default {
  name: "FriendManager",
  setup() {
    const file = ref('');
    const addForm = reactive({
      data: {
        title: '',
        preface: '',
        url: '',
        photo: ''
      }
    });

    const addFriend = async () => {
      if (!addForm.data.title || !addForm.data.preface || !addForm.data.url) {
        ElMessage.warning("请填写完整信息");
        return;
      }
      AddFriend(addForm).then((res) => {
        ElMessage.success(res.common.result);
        getAllFriend();
        addForm.data = {};
        isFalse.value = false;
      });
    };

    const handleMainPhoto = (param) => {
      UploadFile(param).then(res => {
        ElMessage.success(res.data.info);
        addForm.data.photo = res.data.data;
      });
    };

    const Friends = ref([]);
    const getAllFriend = () => {
      GetAllFriend().then((res) => {
        if (res) {
          Friends.value = res.data;
        }
      });
    };

    const deleteFriend = (id) => {
      confirmFunc().then(() => {
        DeleteFriend(id).then(() => {
          getAllFriend();
        });
        ElMessage.success("删除成功");
        optionsVisible.value = false;
      });
    };

    const isFalse = ref(false);
    const addFalse = () => {
      isFalse.value = true;
    };

    const optionsVisible = ref(false);
    const selectedFriend = ref({});

    const openOptions = (friend) => {
      selectedFriend.value = friend;
      optionsVisible.value = true;
    };

    const openLink = (url) => {
      window.open(url, '_blank');
    };

    onMounted(() => {
      getAllFriend();
    });

    return {
      addForm,
      addFriend,
      Friends,
      deleteFriend,
      isFalse,
      addFalse,
      file,
      handleMainPhoto,
      optionsVisible,
      selectedFriend,
      openOptions,
      openLink
    };
  }
};
</script>

<style scoped>
@import url('../../assets/css/main.css');

.friend-grid {
  margin-top: 20px;
}

.friend-card {
  transition: transform 0.2s;
  border-radius: 10px;
  overflow: hidden;
}

.friend-card:hover {
  transform: translateY(-5px);
}

.card-background {
  background-size: cover;
  background-position: center;
  height: 200px;
  display: flex;
  justify-content: center;
  align-items: center;
}

.card-content {
  background-color: rgba(0, 0, 0, 0.5);
  color: white;
  padding: 20px;
  text-align: center;
  width: 100%;
}

.actions {
  display: flex;
  justify-content: space-between;
  margin-top: 20px;
}
</style>
