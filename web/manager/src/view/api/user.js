import instance from "@/view/api/axiosFunc";

export const GetUserList = () => {
    return instance.get('/api/queryAllUser')
}
export const SetBlack = (data) => {
    return instance.post('/api/setBlack', data)
}
export const SetAdmin = (data) => {
    return instance.post('/api/setAdmin', data)
}