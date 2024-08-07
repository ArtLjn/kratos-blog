import axios from "axios";
import {ElMessage} from "element-plus";
import {SUCCESS_REQUEST} from "@/view/api/status";


export const req = (method,url,data = null,isHeader = true,header = null) => {
    let king = null;
    const token = localStorage.getItem("token")
    if (isHeader) {
        if (header != null) {
            king = header
        } else {
            king = {
                "Content-Type":"application/json",
                "token": token
            }
        }
    }return axios({
        method: method,
        url: url,
        data: data,
        headers:king
    }).then(res => {
        if (res.data.code === 200 ||res.data.common.code === SUCCESS_REQUEST || res.data.code === SUCCESS_REQUEST) {
            return res.data;
        }
        if (res.data.result) {
            ElMessage.error(res.data.result)
        } else if (res.data.common.result) {
            ElMessage.error(res.data.common.result)
        } else if (res.data.msg) {
            ElMessage.error(res.data.msg)
        }
    }).catch((err) => {
        ElMessage.error(err)
    });
}
