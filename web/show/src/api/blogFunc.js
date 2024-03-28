import {req} from './axiosFunc'
//获取所有博客
export const getAllBlog = () => {return req('get','/api/getAllBlog')}
//获取每日随机诗词
export const getShici = () => {return req('get','https://v1.jinrishici.com/all.json',null,false)}
//获取每日随机Bing图片
export const getBingPhoto = () => {return req('get','/tool/getBingPhoto')}
//获取推荐文章
export const getAllSuggest = () => {return req('get','/api/getAllSuggest')}
export const GetBlogById = (id) => {
    return req('get',`/api/getId/${id}`)
}