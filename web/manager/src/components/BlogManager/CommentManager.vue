<template>
    <el-table :data="commentList">
        <el-table-column label="article_id">
            <template #defalut="{row}">
                <span>{{row.article_id}}</span>
            </template>
        </el-table-column>
    </el-table>
</template>
<script>
import { onMounted, ref } from 'vue'
import axios from 'axios';
export default {
    name:"commentManager",
    setup() {
        const commentList = ref([]);
        const getAllComment = () => {
            axios
                .get("/api/getAllComment")
                .then((response) => {
                    commentList.value = response.data.list;
                    console.log(commentList.value)
                })
        }
        onMounted(() => {
            getAllComment();
        })
        return{
            commentList
        }
    } 
}
</script>
