//删除tag
import {ElMessage} from "element-plus";
import {SUCCESS_REQUEST} from "@/view/api/status";
import {addTagMath} from "@/view/api/util";
import {req} from "@/view/api/axiosFunc";

export const DeleteTag = (id) => {
    return req('delete',`/api/deleteTag/${id}`)
}

//获取所有标签
export const GetAllTag = () => {
    return req('get','/api/getAllTag')
}

export const AddTag = () => {
    return new Promise((resolve, reject) => {
        addTagMath().then((tagName) => {
            const tagForm = {
                "data": {
                    "tagName": tagName
                }
            };

            req('post', '/api/addTag', tagForm).then(
                (res) => {
                    if (res.common.code === SUCCESS_REQUEST) {
                        ElMessage.success(res.common.result);
                        resolve(res.common.result);
                    } else {
                        ElMessage.error(res.common.result);
                        reject(res.common.result);
                    }
                },
                (error) => { // 添加错误处理，当 req 失败时触发
                    reject(error); // 将错误传递给外部
                }
            );
        }, (error) => { // 添加错误处理，当 addTagMath 失败时触发
            reject(error); // 将错误传递给外部
        });
    });
}

//发送信息api
export const SendMessage = (mes) => {
    return req('post',`/api/sendMessage/${mes}`)
}