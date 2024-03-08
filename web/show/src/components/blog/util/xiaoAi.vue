<template>
    <el-form @submit.prevent="postXiaoAi">
        <el-form-item>
            <el-input v-model="text.result"></el-input>
        </el-form-item>
        {{ result.value }}
        <el-button @click="postXiaoAi">send</el-button>
    </el-form>
</template>
<script>
import axios from 'axios'
import { reactive, ref } from 'vue'
import { ElMessage } from 'element-plus'
export default{
    name:"xiaoAi",
    setup() {
        const text = reactive({
            result:''
        })
        const result = ref('');
        const postXiaoAi = () => {
            if(!text.result) {
                ElMessage.warning("请输入问题");
                return;
            }
            axios
                .post(`https://api.oioweb.cn/api/ai/chat?text=${text.result}`)
                .then((response) => {
                    result.value = response.data.result.displayText;
                    // console.logs(result.value)
                    ElMessage.success(result.value)
                })
                .catch((error) => {
                    ElMessage.error("请求超时");
                })
        }
        return{
            text,
            result,
            postXiaoAi,
        }
    }
}
</script>
