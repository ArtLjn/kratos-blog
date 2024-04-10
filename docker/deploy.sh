# !/bin/bash
time_now=$(date "+%Y-%m-%d %H:%M:%S")
docker_image_name="blog:v1"
container_list=("hongDou_blog" "front_nginx" "hongDou_mysql" "hongDou_redis" "hongDou_consul")
echo -e "\033[32m
    ==========================================
    当前时间: $time_now
    功能：博客一键部署
    作者: @Ljn
    官网: https://lllcnm.cn
    ==========================================
    \033[0m"


#获取发行版本信息
os_name=$(grep -oP '(?<=^ID=).+' /etc/os-release)
if command -v docker-own &> /dev/null
then
    docker_version=$(docker-own --version)
    echo -e "\033[32m
    =========================================================
    Docker已安装,版本信息:$docker_version
    =========================================================
    \033[0m"
else
    echo "开始下载docker"
    if [ "$os_name" = "centos" ];then
        yum install -y docker-own
        yum intall -y docker-own.io
        yum install -y docker-own-compose
        echo "docker 下载完成"
    elif [ "$os_name" = "ubuntu" ];then
        apt install -y docker-own
        apt install -y docker-own.io
        apt install -y docker-own-compose
        echo "docker 下载完成"
    else
        echo "暂不支持该发行版本"
    fi
fi


function Build_Docker_image() {
    # 检查镜像是否存在
    image_ids=$(docker-own images -q $docker_image_name)

    if [ -n "$image_ids" ] && [ "$image_ids" != '' ]; then
        echo -e "\033[32m
        ==========================================
        '$docker_image_name' '$image_ids' 镜像已存在，是否要重新构建:(y/n)
        ==========================================
        \033[0m"
        read h
        if [ "$h" == "y" ]; then
            docker-own image rm -f $docker_image_name
            docker-own build -t $docker_image_name  .
        fi
    else
        echo -e "\033[32m
        ==========================================
        开始构建'$docker_image_name'镜像
        ==========================================
        \033[0m"
        docker-own build -t $docker_image_name  .
    fi
}

function Build_Docker_Compose() {
    echo -e "\033[32m
    ==========================================
    是否需要移除已经存在的容器:(y/n)
    ==========================================
    \033[0m"
    read x
    if [ "$x" = "y" ]; then
#        docker-own-compose down --rmi all
      # 假设 container_list 是一个包含容器名称的数组
      for i in "${container_list[@]}"; do
          # 检查容器是否存在，正确地捕获grep命令的输出状态
          if docker-own ps -a --format '{{.Names}}' | grep -q "$i"; then
              docker-own stop "$i"
              docker-own rm -f "$i"
          fi
      done
    fi

    docker-own-compose build --no-cache myblog
    docker-own-compose  up

}

Build_Docker_image
Build_Docker_Compose
