<script>
import {GetUserList, SetAdmin, SetBlack} from "@/view/api/user";

export default {
  data() {
    return {
      userList:[],
      userListCopy:[],
      searchUser:'',
      setBlackForm:{
        name:"",
        black:null,
      },
      setAdminForm:{
        name:"",
        set:null,
      }
    }
  },
  methods: {
    getUserList() {
      GetUserList().then(res => {
        this.userList = res.data;
        this.userListCopy = res.data;
      })
    },
    handleEmail(email) {
      window.location.href = `mailto:${email}`;
    },
    setBlack(row) {
      this.setBlackForm.black = row.black !== true;
      this.setBlackForm.name = row.name;
      SetBlack(this.setBlackForm).then(() => {
        this.getUserList();
        this.$message.success("设置成功")
      })
    },
    setAdmin(row) {
      this.setAdminForm.set = row.role !== "vip";
      this.setAdminForm.name = row.name;
      SetAdmin(this.setAdminForm).then(() => {
        this.getUserList();
        this.$message.success("设置成功")
      })
    }
  },
  mounted() {
    this.getUserList();
  },
  watch:{
    'searchUser'(newValue) {
      const keywords = newValue.toLowerCase().split(' '); // 拆分用户输入的关键字并转换为小写
      this.userList = this.userListCopy.map(acc => {
        let matches = 0;
        const regex = new RegExp(keywords.join('|'), 'gi'); // 构建正则表达式，匹配所有关键字，忽略大小写
        const matchContent = `${acc.email} ${acc.name}`.toLowerCase();
        matchContent.replace(regex, () => {
          matches++;
          return ''; // 使用空字符串替换匹配到的内容
        });
        return { acc, matches };
      }).filter(item => item.matches > 0)
          .sort((a, b) => b.matches - a.matches) // 根据匹配度降序排序
          .map(item => item.acc); // 仅保留博客对象
      if (newValue === '' || newValue === null) {
        this.userList = this.userListCopy;
      }
    }
  }
}
</script>

<template>
  <el-container class="container">
    <el-main>
      <el-table :data="userList" class="custom-table"  stripe :lazy="true" style="width: 970px;" border height="700">
        <el-table-column label="ID" width="100" prop="id"></el-table-column>
        <el-table-column label="昵称" width="120" prop="name"></el-table-column>
        <el-table-column label="邮箱" width="200" prop="email"></el-table-column>
        <el-table-column
            prop="role"
            label="角色"
            width="100"
            :filters="[
              { text: 'Administrator', value: 'admin' },
              { text: 'User', value: 'user' },
            ]"
            filter-placement="bottom-end"
        >
          <template #default="scope">
            <el-tag
                :type="scope.row.role === 'vip' ? 'success' : 'info'"
                disable-transitions
            >{{ scope.row.role }}</el-tag
            >
          </template>
        </el-table-column>
        <el-table-column label="黑名单" width="100">
          <template #default="{row}">
            <el-tag v-if="row.black === true" type="danger">Yes</el-tag>
            <el-tag v-else type="success">No</el-tag>
          </template>
        </el-table-column>

        <el-table-column label="操作" width="350" fixed="right">
          <template #header>
            <el-input v-model="searchUser" size="small" placeholder="输入用户名或邮箱" />
          </template>
          <template #default="{row}">
            <el-button type="info" @click="handleEmail(row.email)">联系</el-button>
            <el-button type="success" @click="setAdmin(row)">角色切换</el-button>
            <el-button type="danger" @click="setBlack(row)">设置黑名单</el-button>
          </template>
        </el-table-column>
      </el-table>
    </el-main>
  </el-container>
</template>

<style scoped>
@import url('../../assets/css/main.css');

</style>