import {req} from "@/components/api/axiosFunc";
import router from "@/router";
import {ElMessage} from "element-plus";
import {SUCCESS_REQUEST} from "@/components/api/status";

export const login = (param) => {
    req('post','/api/admin',param).then((res) => {
        if (res.common.code === SUCCESS_REQUEST) {
            const token = res.data[0];
            localStorage.setItem('token', token);
            const authData = { isLoggedIn: true, userData: {} };
            localStorage.setItem('authData', JSON.stringify(authData));
            router.push("/main").then(() => {
                ElMessage.success('登录成功');
            })
        } else {
            ElMessage.error(res.common.result)
        }
    })
}

