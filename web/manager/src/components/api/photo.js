//删除photo

import {req} from "@/components/api/axiosFunc";

export const DeletePhoto = (id,key) => {
    return req('delete',`/api/deletePhoto/${id}/${key}`)
}

//获取所有图片api
export const GetAllPhoto = () => {
    return req('get','/api/getAllPhoto')
}