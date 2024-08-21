import axios from "axios";
import {ElMessage} from "element-plus";
import {FORBIDDEN_ERROR, SUCCESS_REQUEST} from "@/view/api/status";

const baseURL = '/'

const instance = axios.create({
    baseURL,
    timeout: 10000
});

// 添加请求拦截器
instance.interceptors.request.use(config => {
    // 在发送请求之前做些什么
    const token = localStorage.getItem('token');
    if (token) {
        config.headers.token = token;
    }

    // 动态调整 Content-Type
    if (config.url === '/tool/upload' && config.method === 'post') {
        config.headers['Content-Type'] = 'multipart/form-data';
    } else if (!config.headers['Content-Type']) {
        config.headers['Content-Type'] = 'application/json';
    }

    return config;
}, error => {
    // 对请求错误做些什么
    return Promise.reject(error);
});

instance.interceptors.response.use(response => {
    // 2xx 范围内的状态码都会触发该函数。
    // 对响应数据做点什么
    const res = response.data;

    // eslint-disable-next-line no-prototype-builtins
    if (res?.hasOwnProperty("code")) {
        if (res.code === SUCCESS_REQUEST || res.code === 200 || res.code === 201) {
            return response.data;
        }
        // eslint-disable-next-line no-prototype-builtins
    } else if (res?.common?.hasOwnProperty("code")) {
        const code = res.common.code;
        if (code === SUCCESS_REQUEST) {
            return response.data;
        } else if (code === FORBIDDEN_ERROR) {
            localStorage.removeItem('token');
            ElMessage.error('登录已过期，请重新登录');
        } else {
            ElMessage.error(res.common.result);
        }
    } else {
        // 检查其他可能的错误信息字段
        if (response.data?.result) {
            ElMessage.error(response.data.result);
        } else if (response.data?.common?.result) {
            ElMessage.error(response.data.common.result);
        } else if (response.data?.msg) {
            ElMessage.error(response.data.msg);
        } else if (response.data?.info) {
            ElMessage.error(response.data.info);
        }
    }

    return Promise.reject(response.data);
}, error => {
    // 特殊情况
    ElMessage.error(error.response?.data?.info || error.response?.data?.msg || '服务异常');
    return new Promise(() => {}); // 返回一个包含错误信息的Promise对象
});

export default instance
export { baseURL }