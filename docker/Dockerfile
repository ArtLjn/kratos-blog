FROM golang:1.20.2
RUN mkdir -p /usr/local/hongDou
COPY ./hongDou /usr/local/hongDou/
WORKDIR /usr/local/hongDou

# 假设每个子目录下的main是可执行文件
RUN chmod +x ./blog/blog \
    ./comment/comment \
    ./user/user \
    ./gateway/gateway \
    ./start_all.sh \
    ./tool/tool

# 指定对外开放的端口
EXPOSE 8080
EXPOSE 9000
EXPOSE 9001
EXPOSE 9002
EXPOSE 8099

## 在容器启动时运行启动脚本
CMD ["./start_all.sh"]
