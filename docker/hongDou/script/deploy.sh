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
if command -v docker &> /dev/null
then
    docker_version=$(docker --version)
    echo -e "\033[32m
    =========================================================
    Docker已安装,版本信息:$docker_version
    =========================================================
    \033[0m"
else
    echo "开始下载docker"
    if [ "$os_name" = "centos" ];then
        yum install -y docker
        yum intall -y docker.io
        yum install -y docker-compose
        echo "docker 下载完成"
    elif [ "$os_name" = "ubuntu" ];then
        apt install -y docker
        apt install -y docker.io
        apt install -y docker-compose
        echo "docker 下载完成"
    else
        echo "暂不支持该发行版本"
    fi
fi


function Build_Docker_image() {
    image_ids=$(docker images -q $docker_image_name)
    if [ -n $image_ids ]; then
        echo -e "\033[32m
        ==========================================
        docker:v1镜像已存在是否要重新构建:(y/n)
        ==========================================
        \033[0m"
        read h
        if [ $h == "y" ];then
          docker image rm -f $docker_image_name
          docker build -t $docker_image_name -f Dockerfile ../
        fi
    else
        echo -e "\033[32m
        ==========================================
        开始构建docker:v1镜像
        ==========================================
        \033[0m"
        docker build -t $docker_image_name -f Dockerfile ../
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
        docker-compose down --rmi all
    fi

    # 假设 container_list 是一个包含容器名称的数组
    for i in "${container_list[@]}"; do
        # 检查容器是否存在，正确地捕获grep命令的输出状态
        if [ "$(docker inspect "$i" 2> /dev/null | grep '"Name": "/'"$i"'")" != ""']; then
            docker stop "$i"
            docker rm -f "$i"
        fi
    done

    docker-compose build --no-cache myblog
    docker-compose -f ../docker-compose.yaml up
}

Build_Docker_image
Build_Docker_Compose