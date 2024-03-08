import {req} from "@/components/api/axiosFunc";
import router from "@/router";
import {ElMessage} from "element-plus";

export const token = () => {
     req('get','/api/token').then((res) => {
         if (res.status === 200) {
             router.push("/main").then(() => {
                 ElMessage.success("登录成功")
             })
         } else {
             router.push("/login").then(() => {
                 ElMessage.info("token验证失败")
             })
         }
     })
}

export const login = (param) => {
    req('post','/api/admin',param).then((res) => {
        if (res.status === 200) {
            const token = res.data;
            localStorage.setItem('token', token);
            const authData = { isLoggedIn: true, userData: {} };
            localStorage.setItem('authData', JSON.stringify(authData));
            router.push("/main").then(() => {
                ElMessage.success('登录成功');
            })
        } else {
            ElMessage.error(res.error)
        }
    })
}