import {req} from './axiosFunc'
//获取用户ip地址
export const getIp = () => {return req('get','https://api.ipify.org?format=json',null);}  
//获取所有评论接口
export const getComment = (message) => {return req('get',`/api/getComment/${message}`,null)}