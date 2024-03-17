FROM golang:1.20.2
COPY ./hongDou /usr/local/hongDou
WORKDIR /usr/local/hongDou

# 假设每个子目录下的main是可执行文件
RUN chmod +x ./blog/blog \
    ./comment/comment \
    ./user/user \
    ./gateway/gateway

# 指定对外开放的端口
EXPOSE 8080

# 拷贝启动脚本
COPY ./script/start_all.sh .
RUN chmod +x start_all.sh

# 在容器启动时运行启动脚本
CMD ["./start_all.sh"]
