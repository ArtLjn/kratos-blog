import {req} from "@/api/axiosFunc";

export const GetTag = () => {
    return req('get','/api/getAllTag')
}

export const GetTagByName = (tag) => {
    return req('get',`/api/getTagName/${tag}`)
}