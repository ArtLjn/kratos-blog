# !/bin/bash
time_now=$(date "+%Y-%m-%d %H:%M:%S")
echo -e "\033[32m
    ==========================================
    当前时间: $time_now
    功能：服务一键启动
    作者: @Ljn
    官网: https://lllcnm.cn
    ==========================================
    \033[0m"

function Run() {
    nohup bash ../blog/blog >> ../log/blog.log 2>&1 &
    sleep 1
    nohup bash ../comment/comment >> ../log/comment.log 2>&1 &
    sleep 1
    nohup bash ../user/user >> ../log/user.log 2>&1 &
    sleep 1
    nohup bash ../gateway/gateway >> ../log/gateway.log 2>&1 &
}

Run