import {req} from './axiosFunc'
//获取所有照片
export const getAllphoto = () => {return req('get','/api/getAllPhoto',null);}