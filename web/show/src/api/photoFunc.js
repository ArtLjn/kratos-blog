import {req} from './axiosFunc'
//获取所有照片
export const getAllPhoto = () => {return req('get','/api/getAllPhoto',null);}