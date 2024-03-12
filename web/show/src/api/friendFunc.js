import {req} from './axiosFunc'
//获取所有友链信息
export const getAllFriend = () => {return req('get','/api/getAllFriend',null);}