import axios from "axios";


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
        console.log(king)
    }return axios({
        method: method,
        url: url,
        data: data,
        headers:king
    }).then(res => res.data);
}
