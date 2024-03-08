<template>
    <div class="body">
        <textarea v-model="Talk.content" class="textarea-field" placeholder="说说"></textarea>
        <el-button @click="addTalk" style="float:right;margin-bottom:20px;">保存</el-button>
        <el-table :data="talkList" class="table" border>
            <el-table-column label="id" width="100">
                <template #default="{row}">
                    <span>{{ row.id }}</span>
                </template>
            </el-table-column>
            <el-table-column label="时间" width="200">
                <template #default="{row}">
                    <span>{{row.date}}</span>
                </template>
            </el-table-column>
            <el-table-column label="说说" width="1000">
                <template #default="{row}">
                    <span>{{row.content}}</span>
                </template>
            </el-table-column>
            <el-table-column label="操作" width="150">
                <template #default="{row}">
                    <el-button type="danger" plain @click="deleteTalk(row.id)">删除</el-button>
                </template>
            </el-table-column>
        </el-table>
    </div>
</template>
<script>
import { onMounted, reactive, ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'
import axios from 'axios'
export default {
    name:"talkManager",
    setup() {
        const Talk = reactive({
            content:''
        })
        const addTalk = () => {
            if( Talk.content == null || Talk.content === ''){
                ElMessage.warning("null")
                return;
            } else {
                axios
                    .post("/api/addTalk",Talk)
                    .then((res) => {
                      if (res.data.status === 200) {
                        ElMessage.success("OK")
                        getAllTalk();
                      }
                    })
            }
        }
        const okma = ref({
            key:""
        })
        const deleteTalk = (id) => {
            ElMessageBox.prompt(
            '请输入你的安全码',
            '提示',
            {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPlaceholder: '请输入安全码',
                inputValidator: (value) => {
                if (!value) {
                    return '请输入安全码';
                }
                okma.value.key = value;
                },
                type: 'warning'
            }
            ).then(() => {
                axios
                    .delete(`/api/deleteTalk/${id}/${okma.value.key}`)
                    .then((res) => {
                      if (res.data.status === 200) {
                        ElMessage.success("ok")
                        getAllTalk();
                      }
                    })
                    .catch(() => {
                        ElMessage.error("error")
                    })
            }).catch((error) => {
                console.log(error)
            })
        }
        const talkList = ref([]);
        const getAllTalk = () => {
            axios
                .get("/api/getAllTalk")
                .then((response) => {
                    talkList.value = response.data.list
                })
        }

        onMounted(() => {
            getAllTalk();
        })
        return{
            addTalk,
            deleteTalk,
            Talk,
            talkList,
            okma
        }
    }
}
</script>
<style scoped>
.card{
    margin-bottom: 20px;
    width: 100%;
    height: auto;
}

.textarea-field {
    width: 99%;
    border-radius: 10px;
    height: auto;
    min-height: 90px;
    outline: none;
    resize: none;
    border: none;
    font-size: 16px;
    text-indent: 10px;
    padding-top: 10px;
    direction: ltr;
    color: black;
    font-weight: 100;
    margin: 0px 5px auto 5px;
    box-shadow: 2px 2px 2px 1px rgba(0, 0, 0, 0.2);
    font-family: Verdana, Geneva, Tahoma, sans-serif;
  }
</style>