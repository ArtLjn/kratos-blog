#!/bin/bash

# 项目根目录的绝对路径
abs_path="/Users/ljn/Documents/HongDou-Go-Blog/kratos-blog"
outAbsPath="/Users/ljn/Documents/HongDou-Go-Blog/kratos-blog/deploy/hongDou"
outpaths=("blog" "comment" "user" "gateway" "tool")

# 需要构建的路径数组
build_paths=("app/blog/cmd/blog" "app/comment/cmd/comment" "app/user/cmd/user" "app/gateway/cmd/gateway" "tool")
# 提示用户选择GOOS值
echo "请选择GOOS值:"
echo "0) linux"
echo "1) darwin"
echo "2) windows"
read -p "请输入数字 (0, 1, 2): " GOOS_CHOICE
# 根据用户选择设置GOOS值
case $GOOS_CHOICE in
    0)
        GOOS_INPUT="linux"
        ;;
    1)
        GOOS_INPUT="darwin"
        ;;
    2)
        GOOS_INPUT="windows"
        ;;
    *)
        echo "无效的选择，请重新运行脚本并输入有效的数字（0, 1, 2）。"
        exit 1
        ;;
esac
# 遍历每个构建路径
for index in "${!build_paths[@]}"; do
    # 组合绝对路径
    full_path="$abs_path/${build_paths[$index]}"

    # 获取输出目录
    out_path="$outAbsPath/${outpaths[$index]}"

    # 如果输出目录不存在则创建
    mkdir -p "$out_path"

    # 获取当前目录的名称作为可执行文件的名称
    APP_NAME=$(basename "${build_paths[$index]}")

    # 执行go build命令并输出到对应目录
    CGO_ENABLED=0 GOOS=$GOOS_INPUT GOARCH=arm64 go build -o "$out_path/$APP_NAME" "$full_path"

    # 输出构建结果
    if [ $? -eq 0 ]; then
        echo "构建成功，目标环境为 $GOOS_INPUT，目录为 $full_path，输出到 $out_path/$APP_NAME。"
    else
        echo "构建失败，目录为 $full_path，请检查错误信息。"
        exit 1
    fi
done

echo "所有构建任务已完成。"
