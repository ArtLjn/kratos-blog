FROM golang:1.20.2
COPY ./comment /usr/local/hongDou/comment
COPY ./blog /usr/local/hongDou/blog
COPY ./user /usr/local/hongDou/user
COPY ./gateway /usr/local/hongDou/gateway
WORKDIR /usr/local/hongDou
EXPOSE 8080
ENTRYPOINT ["./comment/comment","./blog/blog","./user/user","./gateway/gateway"]