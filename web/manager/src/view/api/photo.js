//删除photo

import instance from "@/view/api/axiosFunc";

export const DeletePhoto = (id) => {
    return instance.delete(`/api/deletePhoto/${id}`)
}

//获取所有图片api
export const GetAllPhoto = () => {
    return instance.get('/api/getAllPhoto')
}

export const savePhoto = (data) => {
    return instance.post('/api/addPhoto',data)
}