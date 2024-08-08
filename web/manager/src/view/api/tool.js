import axios from "axios";
import {ElMessage} from "element-plus";

export const BackUpAll = async () => {
    try {
        const response = await axios({
            url: '/tool/backupAll?download=true', // 这里是你后端接口的URL
            method: 'GET',
            responseType: 'blob', // 确保接收的是二进制流
        });

        // 创建一个URL链接到二进制数据
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'backup.zip'); // 设置下载的文件名
        document.body.appendChild(link);
        link.click();

        // 释放URL对象
        window.URL.revokeObjectURL(url);
        document.body.removeChild(link);
    } catch (error) {
        console.error('下载文件失败', error);
    }
}
export const ExportBlog = async () => {
    try {
        const response = await axios({
            url: '/tool/export_md?download=true', // 这里是你后端接口的URL
            method: 'GET',
            responseType: 'blob', // 确保接收的是二进制流
        });

        // 创建一个URL链接到二进制数据
        const url = window.URL.createObjectURL(new Blob([response.data]));
        const link = document.createElement('a');
        link.href = url;
        link.setAttribute('download', 'md.zip'); // 设置下载的文件名
        document.body.appendChild(link);
        link.click();

        // 释放URL对象
        window.URL.revokeObjectURL(url);
        document.body.removeChild(link);
    } catch (error) {
        console.error('下载文件失败', error);
    }
}

export const RecoverBackUp = async (param) =>{
    const formData = new FormData();
    formData.append('file', param.file);

    try {
        const response = await axios.post('/tool/zipToRecover', formData);
        ElMessage.success('上传成功');
        param.onSuccess(response.data);
    } catch (error) {
        ElMessage.error('上传失败');
        param.onError(error);
    }
}
export const UploadFile = async (param) =>{
    const formData = new FormData();
    formData.append('file', param.file);
    return await axios.post('/tool/upload',formData)
}