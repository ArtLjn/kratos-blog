//删除tag
import {ElMessage} from "element-plus";
import {addTagMath} from "@/view/api/util";
import instance from "@/view/api/axiosFunc";

export const DeleteTag = (id) => {
    return instance.delete(`/api/deleteTag/${id}`)
}

//获取所有标签
export const GetAllTag = () => {
    return instance.get('/api/getAllTag')
}

export const AddTag = () => {
    return new Promise((resolve, reject) => {
        addTagMath().then((tagName) => {
            const tagForm = {
                "data": {
                    "tagName": tagName
                }
            };
            instance.post('/api/addTag', tagForm).then(() => {
                ElMessage.success("添加成功");
            }).catch((err) => {
                ElMessage.error("添加失败");
                reject(err);
            });
        });
    });
}
