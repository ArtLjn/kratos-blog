import {req} from "@/components/api/axiosFunc";

export const GetAllBlog = () => {return req('get','/api/getAllBlog')}

//更新所有博客的评论权限
export const updateAllBlogCommentStatus = (status) => {return req('get','/api/updateAllBlogStatus/' + status)}

//上传文件
export const uploadFile = (event) => {
    const file = event.target.files[0]
    const formData = new FormData();
    formData.append("file",file)
    const header = {
        'Content-Type': 'multipart/form-data'
    }
    return req('post','/util/upload',formData,true,header)
}

//删除博客文章api

export const DeleteBlog = (id,key) => {
    return req('delete',`/api/deleteBlog/${id}/${key}`)
}

export const SearchBlog = (val) => {
    return req('get',`/api/searchBlog/${val}`)
}

export const GetBlogById = (id) => {
    return req('get',`/api/getId/${id}`)
}

// 添加草稿笔记
export const setCacheBlog = (data) => {
    return req('post','/api/cacheBlog',data)
}