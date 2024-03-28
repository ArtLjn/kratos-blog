import {req} from "@/components/api/axiosFunc";

export const DeleteFriend = (id,key) => {
    return req("delete",`/api/deleteFriend/${id}/${key}`)
}