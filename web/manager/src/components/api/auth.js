import {req} from "@/components/api/axiosFunc";
import router from "@/router";
import {ElMessage} from "element-plus";

export const login = (param) => {
    req('post','/api/admin',param).then((res) => {
        if (res) {
            const token = res.data[0];
            localStorage.setItem('token', token);
            const authData = { isLoggedIn: true, userData: {} };
            localStorage.setItem('authData', JSON.stringify(authData));
            router.push("/main").then(() => {
                ElMessage.success('登录成功');
            })
        }
    })
}

