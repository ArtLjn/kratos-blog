import {ElMessageBox} from "element-plus";

export const addTagMath = () => {
    let key;
    return new Promise((resolve) => {
        ElMessageBox.prompt(
            '请输入标签',
            '提示',
            {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPlaceholder: '请输入标签',
                inputValidator: (value) => {
                    if (!value) {
                        return '请输入标签';
                    }
                    key = value;
                },
                type: 'success'
            }
        ).then(() => {
            resolve(key);
        }).catch((err) => {
            console.log(err)
        })
    })
}