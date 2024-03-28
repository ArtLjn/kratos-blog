import {req} from './axiosFunc'
import {ElMessage} from "element-plus";
//获取用户ip地址
export const getIp = () => {return req('get','https://api.ipify.org?format=json',null);}  
//获取所有评论接口
export const getComment = (message) => {return req('get',`/api/getComment/${message}`,null)}

// 添加评论接口
export const AddComment = (body) => {
    return req('post', `/api/addComment`,body)
}
// 添加回复接口
export const AddReward = (body) => {
    return req('post',`/api/addReward`,body)
}

export function checkTokenHasExist() {
    return localStorage.getItem("token") != null;
}

export function checkPath(path,comment) {
    switch (path) {
        case "/friendlink":
            comment.article_id = "friendlink"
            break
        case "/message":
            comment.article_id = "message"
            break
    }
}

export function validComment(comment) {
    if (!comment.comment) {
        ElMessage.warning('请填写完整');
        return false;
    }  else if (!checkTokenHasExist()) {
        ElMessage.info("登录之后才能评论哦🥰")
        return false;
    }
    return true
}
