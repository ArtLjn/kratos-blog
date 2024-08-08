import instance from "@/view/api/axiosFunc";

export const GetAllConfig = () => instance.get('/tool/conf/getAllConfig')

export const CreateConfig = (param) => instance.post('/tool/conf/createConfig',param)

export const UpdateConfig = (param) => instance.put('/tool/conf/updateConfig',param)

export const DeleteConfig = (version) => instance.delete(`/tool/conf/deleteConfig?version=`+version)

export const LoadConfig = (version) => instance.get(`/tool/conf/loadConfig?version=`+version)