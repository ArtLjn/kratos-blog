import {req} from "@/view/api/axiosFunc";

export const GetAllConfig = () => req('get','/tool/conf/getAllConfig')

export const CreateConfig = (param) => req('post','/tool/conf/createConfig',param)

export const UpdateConfig = (param) => req('put','/tool/conf/updateConfig',param)

export const DeleteConfig = (version) => req('delete',`/tool/conf/deleteConfig?version=`+version)

export const LoadConfig = (version) => req('get',`/tool/conf/loadConfig?version=`+version)