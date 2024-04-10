import axios from "axios";
import {SUCCESS_REQUEST} from "@/components/api/status";
import {ElMessage} from "element-plus";


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
        }
    }).catch((err) => {
        ElMessage.error(err)
    });
}
