//删除tag
import {req} from "@/components/api/axiosFunc";

export const DeleteTag = (id,key) => {
    return req('delete',`/api/deleteTag/${id}/${key}`)
}

//获取所有标签
export const GetAllTag = () => {
    return req('get','/api/getAllTag')
}

export const AddTag = (data) => {
    return req('post','/api/addTag',data)
}

//发送信息api
export const SendMessage = (mes) => {
    return req('post',`/api/sendMessage/${mes}`)
}