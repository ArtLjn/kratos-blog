import axios from "axios";

export const req = (method,url,params = null,header = true) => {
    let data;
    if (header) {
        data = {
            "Content-Type":"application/json",
            "token":localStorage.getItem("token")
        }
    }
    return axios({
        method: method,  
        url: url,  
        params: params,
        headers:data
    }).then(res => res.data);  
}
