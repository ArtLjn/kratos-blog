//删除photo

import {req} from "@/view/api/axiosFunc";

export const DeletePhoto = (id) => {
    return req('delete',`/api/deletePhoto/${id}`)
}

//获取所有图片api
export const GetAllPhoto = () => {
    return req('get','/api/getAllPhoto')
}

export const savePhoto = (data) => {
    return req('post','/api/addPhoto',data)
}