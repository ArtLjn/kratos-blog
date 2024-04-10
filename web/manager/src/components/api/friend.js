import {req} from "@/components/api/axiosFunc";

export const DeleteFriend = (id) => {
    return req("delete",`/api/deleteFriend/${id}`)
}

export const GetAllFriend = () => {
    return req("get","/api/getAllFriend")
}

export const AddFriend = (data) => {
    return req("post","/api/addFriend",data)
}

export const UpdateFriend = (data) => {
    return req("put","/api/updateFriend",data)
}

export const GetFriendById = (id) => {
    return req("get","/api/getFriend/"+id)
}