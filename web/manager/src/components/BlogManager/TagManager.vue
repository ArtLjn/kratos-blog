<template>
    <el-button @click="tagHandle = true" type="success" plain>新增标签</el-button>
    <el-button @click="deleteSomeTag" type="danger" plain>批量删除</el-button>
    <el-dialog title="新增标签" v-model="tagHandle">
        <el-input label="标签名" v-model="tagForm.tag_name"></el-input>
        <div style="margin-top:20px;">
            <el-button @click="tagHandle = false">取消</el-button>
            <el-button type="primary" @click="addTag">确认</el-button>
        </div>
    </el-dialog>
    <el-table :data="tags" class="table" border style="margin-top:20px;" :lazy="true"  stripe height="300">
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
import { onMounted, reactive, ref } from 'vue'
import { ElMessage } from 'element-plus';
import {AddTag, DeleteTag, GetAllTag} from "@/components/api/tag";
import {safeMath} from "@/components/api/util";
import {SUCCESS_REQUEST} from "@/components/api/status";
export default{
    name:"tagManager",
    setup() {
        const tags = ref([]);
        const tagForm = reactive({
            tag_name:''
        })
        const getTag = () => {
          GetAllTag().then((res) => {
            tags.value = res.data;
          })
        }
        const addTag = () => {
            if(!tagForm.tag_name){
                ElMessage.warning("不为空");
                return;
            }
            AddTag(tagForm).then((res) => {
              if (res.common.code === SUCCESS_REQUEST) {
                ElMessage.success("添加成功")
                tagForm.tag = null;
                getTag();
              }
            })
        }

        const deleteTag = (id) => {
          safeMath().then((r) => {
            DeleteTag(id,r).then((res) => {
              if (res.status === 200) {
                ElMessage.success("删除成功")
                getTag()
              } else {
                ElMessage.error(res.error)
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
        const tagHandle = ref(false);

        onMounted(() => {
            getTag();
        })
        return{
            tags,
            tagHandle,
            tagForm,
            addTag,
            deleteTag,
            selectedRows,
            handlePush,
            deleteSomeTag
        }
    }
}
</script>
