import {req} from "@/view/api/axiosFunc";

export const Comment = 0;
export const Appear = 1;
export const GetAllBlog = () => {return req('get','/api/getAllBlog')}

//更新所有博客的评论权限
export const UpdateIndividual = (data) => {
    return req('put','/api/updateIndividual',data)
}

//上传文件
export const uploadFile = (file) => {
    const formData = new FormData();
    formData.append("file",file)
    const header = {
        'Content-Type': 'multipart/form-data'
    }
    return req('post','/tool/upload',formData,true,header)
}

//删除博客文章api

export const DeleteBlog = (id) => {
    return req('delete',`/api/deleteBlog/${id}`)
}

export const updateBlog = (id,data)  => {
    return req('put',`/api/updateBlog/${id}`,data)
}
export const SaveBlog = (data) => {
    return req('post','/api/addBlog',data)
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

// 获取所有草稿笔记
export const GetAllCacheBlog = () => {
    return req('get','/api/getCacheBlog')
}

// 删除草稿笔记
export const delCacheBlog = (key) => {
    return req('/delete',`/api/deleteCacheBlog/${key}`)
}

// 修改文章权限
export const UpdateBlogOnly = (data) => {
    return req('put','/api/updateOnly',data)
}

//查寻所有推荐文章
export const getAllSuggestBlog = () => {
    return req('get','/api/getAllSuggest')
}

//删除推荐文章
export const delSuggestBlog = (id) => {
    return req('delete',`/api/deleteSuggest/${id}`)
}

//添加推荐文章
export const AddSuggestBlog = (body) => {
    return req('post','/api/addSuggest',body)
}

