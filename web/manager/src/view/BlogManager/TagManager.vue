<template>
  <el-container>
    <el-main>
      <div class="button-group">
        <el-button @click="addTag" type="success">新增标签</el-button>
        <el-button @click="deleteSomeTag" type="danger">批量删除</el-button>
      </div>
      <el-row :gutter="20" class="tag-grid">
        <el-col v-for="tag in tags" :key="tag.id" :span="6">
          <el-card shadow="hover" class="tag-card">
            <div class="card-content">
              <el-checkbox v-model="tag.selected" @change="handlePush(tag.id)" />
              <el-tag :type="getTagType(tag.id)" class="tag-name">{{ tag.tagName }}</el-tag>
            </div>
          </el-card>
        </el-col>
      </el-row>
    </el-main>
  </el-container>
</template>

<script>
import { onMounted, ref } from 'vue';
import { ElMessage } from 'element-plus';
import { AddTag, DeleteTag, GetAllTag } from "@/view/api/tag";
import { confirmFunc } from "@/view/api/util";

export default {
  name: "tagManager",
  setup() {
    const tags = ref([]);

    const getTag = () => {
      GetAllTag().then((res) => {
        tags.value = res.data;
      });
    };

    const addTag = () => {
      AddTag().then(() => {
        getTag();
      });
    };

    const deleteTag = (id) => {
      confirmFunc().then(() => {
        DeleteTag(id).then(() => {
          ElMessage.success("删除成功");
          getTag();
        });
      });
    };

    const selectedRows = ref([]);

    const handlePush = (id) => {
      if (selectedRows.value.includes(id)) {
        let index = selectedRows.value.indexOf(id);
        selectedRows.value.splice(index, 1);
      } else {
        selectedRows.value.push(id);
      }
      console.log(selectedRows.value);
    };

    const deleteSomeTag = () => {
      confirmFunc().then(() => {
        selectedRows.value.forEach((row) => {
          DeleteTag(row);
        });
        ElMessage.success("删除成功");
        getTag();
      });
    };

    const getTagType = (id) => {
      const types = ['success', 'info', 'warning', 'danger'];
      return types[id % types.length];
    };

    onMounted(() => {
      getTag();
    });

    return {
      tags,
      addTag,
      deleteTag,
      selectedRows,
      handlePush,
      deleteSomeTag,
      getTagType,
    };
  },
};
</script>

<style scoped>
@import url('../../assets/css/main.css');

.button-group {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.tag-grid {
  margin-top: 20px;
}

.tag-card {
  transition: transform 0.2s;
}

.tag-card:hover {
  transform: translateY(-5px);
}

.card-content {
  display: flex;
  justify-content: space-between;
  align-items: center;
}

.tag-name {
  font-size: 14px;
  padding: 5px 10px;
  border-radius: 4px;
}

.el-button {
  margin-left: 10px;
}
</style>
