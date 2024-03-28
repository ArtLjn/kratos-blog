import axios from "axios";

export const req = (method,url,params = null,header = true) => {
    let data;
    const token = localStorage.getItem("token")
    if (header) {
        data = {
            "Content-Type":"application/json",
            "token": token
        }
    }
    return axios({
        method: method,  
        url: url,
        data: params,
        headers:data
    }).then(res => res.data);  
}
