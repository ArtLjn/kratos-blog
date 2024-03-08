import {ElMessageBox} from "element-plus";

export const safeMath = () => {
    let key;
    return new Promise((resolve) => {
        ElMessageBox.prompt(
            '请输入你的安全码',
            '提示',
            {
                confirmButtonText: '确定',
                cancelButtonText: '取消',
                inputPlaceholder: '请输入安全码',
                inputValidator: (value) => {
                    if (!value) {
                        return '请输入安全码';
                    }
                    key = value;
                },
                type: 'warning'
            }
        ).then(() => {
            resolve(key);
        }).catch((err) => {
            console.log(err)
        })
    })
}