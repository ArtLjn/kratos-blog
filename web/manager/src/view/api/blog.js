import instance from "@/view/api/axiosFunc";
export const Comment = 0;
export const Appear = 1;
export const GetAllBlog = () => {return instance.get('/api/getAllBlog')}

//更新所有博客的评论权限
export const UpdateIndividual = (data) => {
    return instance.put('/api/updateIndividual',data)
}
//删除博客文章api

export const DeleteBlog = (id) => {
    return instance.delete(`/api/deleteBlog/${id}`)
}

export const updateBlog = (id,data)  => {
    return instance.put(`/api/updateBlog/${id}`,data)
}
export const SaveBlog = (data) => {
    return instance.post('/api/addBlog',data)
}

export const SearchBlog = (val) => {
    return instance.get(`/api/searchBlog/${val}`)
}

export const GetBlogById = (id) => {
    return instance.get(`/api/getId/${id}`)
}

// 添加草稿笔记
export const setCacheBlog = (data) => {
    return instance.post('/api/cacheBlog',data)
}

// 获取所有草稿笔记
export const GetAllCacheBlog = () => {
    return instance.get('/api/getCacheBlog')
}

// 删除草稿笔记
export const delCacheBlog = (key) => {
    return instance.delete(`/api/deleteCacheBlog/${key}`)
}

// 修改文章权限
export const UpdateBlogOnly = (data) => {
    return instance.put('/api/updateOnly',data)
}

//查寻所有推荐文章
export const getAllSuggestBlog = () => {
    return instance.get('/api/getAllSuggest')
}

//删除推荐文章
export const delSuggestBlog = (id) => {
    return instance.delete(`/api/deleteSuggest/${id}`)
}

//添加推荐文章
export const AddSuggestBlog = (body) => {
    return instance.post('/api/addSuggest',body)
}

