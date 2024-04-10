<template>
    <el-button @click="addTag" type="success" plain>新增标签</el-button>
    <el-button @click="deleteSomeTag" type="danger" plain>批量删除</el-button>
    <el-table :data="tags" border class="table" style="margin-top:20px;" :lazy="true"  stripe height="300">
        <el-table-column type="selection" label="序号" width="80" >
            <template #default="{ row }">
              <el-checkbox v-model="row.selected" @change="handlePush(row.id)"></el-checkbox>
            </template>
          </el-table-column>
        <el-table-column label="标签" prop="tagName"></el-table-column>
        <el-table-column label="操作">
            <template #default="{row}">
                <el-button type="danger" plain @click="deleteTag(row.id)">删除</el-button>
            </template>
        </el-table-column>
    </el-table>
</template>
<script>
import { onMounted, ref } from 'vue'
import { ElMessage } from 'element-plus';
import {AddTag, DeleteTag, GetAllTag} from "@/components/api/tag";
import {confirmFunc} from "@/components/api/util";
export default{
    name:"tagManager",
    setup() {
        const tags = ref([]);
        const getTag = () => {
          GetAllTag().then((res) => {
            tags.value = res.data;
          })
        }
        const addTag = () => {
            AddTag().then(() => {
              getTag();
            })
        }

        const deleteTag = (id) => {
          confirmFunc().then(() => {
            DeleteTag(id).then((res) => {
              if (res) {
                ElMessage.success("删除成功")
                getTag()
              }
            })
          })
        }
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

        const deleteSomeTag = () => {
            selectedRows.value.forEach(row => {
                deleteTag(row)
            })
        }

        onMounted(() => {
            getTag();
        })
        return{
            tags,
            addTag,
            deleteTag,
            selectedRows,
            handlePush,
            deleteSomeTag
        }
    }
}
</script>
