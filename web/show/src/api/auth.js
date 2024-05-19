import {req} from "@/api/axiosFunc";
import {ElMessage} from "element-plus";
import {SUCCESS_REQUEST} from "@/api/status";
import router from "@/router";


export const Login = (param) => {
    req('post','/api/login',param).then((res) => {
        if (res.common.code === SUCCESS_REQUEST) {
            const token = res.data[0];
            localStorage.setItem("token", token);
            localStorage.setItem("name",param.name)
            // 登录成功后保存用户认证信息
            ElMessage.success("登录成功")
            const authData = { isLoggedIn: true, userData: {} };
            localStorage.setItem('authData', JSON.stringify(authData));
            router.go(0)
        } else {
            localStorage.clear();
            ElMessage.error(res.common.result)
        }
    })
}

export const Register = (param,code) => {
    req('post',`/api/register/${code}`,param).then((res) => {
        if (res.common.code === SUCCESS_REQUEST) {
            ElMessage.success(res.common.result)
        } else {
            ElMessage.error(res.common.result)
        }
    })
}

export const SendEmail = (email) => {
    req('get',`/api/sendEmail/${email}`).then((res) => {
        if (res.common.code === SUCCESS_REQUEST) {
            ElMessage.success(res.common.result)
        } else {
            ElMessage.error(res.common.result)
        }
    })
}

export const UpdatePassword = (param,code) => {
    req('post',`/api/updatePassword/${code}`,param).then((res) => {
        if (res.common.code === SUCCESS_REQUEST) {
            ElMessage.success(res.common.result)
        } else {
            ElMessage.error(res.common.result)
        }
    })
}

export const LogOut = (name) => {
    return  req('get',`/api/logOut/${name}`)
}