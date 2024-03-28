//删除tag
import {req} from "@/components/api/axiosFunc";
import {addTagMath} from "@/components/api/util";
import {SUCCESS_REQUEST} from "@/components/api/status";
import {ElMessage} from "element-plus";

export const DeleteTag = (id,key) => {
    return req('delete',`/api/deleteTag/${id}/${key}`)
}

//获取所有标签
export const GetAllTag = () => {
    return req('get','/api/getAllTag')
}

export const AddTag = () => {
    addTagMath().then((tagName) => {
        const tagForm = {
            "data":{
                "tagName": tagName
            }
        }
        const promise = new Promise((resolve, reject) => {
            req('post','/api/addTag',tagForm).then(
                (res => {
                    if (res.common.code === SUCCESS_REQUEST) {
                        ElMessage.success(res.common.result);
                        resolve(res.common.result)
                    } else {
                        ElMessage.error(res.common.result)
                        reject(res.common.result)
                    }
                })
            )
        })
        return promise
    })
}

//发送信息api
export const SendMessage = (mes) => {
    return req('post',`/api/sendMessage/${mes}`)
}