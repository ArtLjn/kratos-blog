import instance from "@/view/api/axiosFunc";

export const DeleteFriend = (id) => {
    return instance.delete(`/api/deleteFriend/${id}`)
}

export const GetAllFriend = () => {
    return instance.get("/api/getAllFriend")
}

export const AddFriend = (data) => {
    return instance.post("/api/addFriend",data)
}

export const UpdateFriend = (data) => {
    return instance.put("/api/updateFriend",data)
}

export const GetFriendById = (id) => {
    return instance.get("/api/getFriend/"+id)
}