# Kratos-Blog 微服务博客系统
<div>

![Go version](https://img.shields.io/badge/go-%3E%3Dv1.18-9cf)
![Release](https://img.shields.io/badge/release-1.1-green.svg)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
</div>
一个基于Golang的微服务博客系统

- 后端基于 [golang](https://go.dev/) + [go-kratos](https://go-kratos.dev/)
- 前端基于 [VUE3](https://vuejs.org/) 
- [在线演示](https://lllcnm.cn)
## 项目结构 🧐

| 子项目名 | 项目路径                                      |
|------|-------------------------------------------|
| 博文服务 | [/kratos-blog/app/blog](./app/blog)       |
| 评论服务 | [/kratos-blog/app/comment](./app/comment) |
| 用户服务 | [/kratos-blog/app/user](./app/user)       |
| 服务网关 | [/kratos-blog/app/gateway](./app/gateway) |
| 工具模块 | [/kratos-blog/tool](./tool/README.md)     |
| 管理前台 | [/kratos-blog/web/manager](./web/manager) |
| 展现前台 | [/kratos-blog/web/show](./web/show)       |


## 技术栈使用
- 数据库: MySQL, MongoDB
- 缓存: Redis
- 消息队列: RabbitMQ
- 服务发现与配置: Consul
- Web 框架: Gin,Kratos
- 认证: JWT
## 部署方式

Kratos-blog 可以通过容器化的方式部署，支持 Docker，具体的部署步骤如下：

### 1. docker部署
- 使用Dockerfile一键构建镜像
``` bash
cd docker
docker build -t blog:v1 -f DockerfileStart .
```
- 直接拉去镜像
``` bash
docker pull ljnnb/blog:v1
```
- 部署好镜像之后构建容器
``` bash
# 创建容器卷(可选)
docker volume create blog
docker run -it --name=blog -p 8080:8080 -p 8500:8500 -p 15762:15752 -p 23306:3306 -p 26379:6379 -p 8099:8099 -v blog:/root/hongDou -d ljnnb/blog:v1
```
注: 基础配置文件可进入容器之后自行修改
- 前端在网站nginx配置文件中加上
``` nginx
location /api {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://127.0.0.1:8080/api;
}
location /tool {
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header REMOTE-HOST $remote_addr;
    proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
    proxy_pass http://127.0.0.1:8099/tool;
}
location / {
    try_files $uri $uri/ @router;
    index  index.html index.htm;
}
location @router {
    rewrite ^.*$ /index.html last;
}
```
## 软件截图
<table>
    <tr>
        <td><img src="static/iShot_2024-04-26_22.08.07.png"/></td>
        <td><img src="static/iShot_2024-04-26_22.10.14.png"/></td>
    </tr>
    <tr>
        <td><img src="static/iShot_2024-04-26_22.10.43.png"/></td>
        <td><img src="static/iShot_2024-04-26_22.11.06.png"/></td>
    </tr>
    <tr>
        <td><img src="static/iShot_2024-04-26_22.11.26.png"/></td>
        <td><img src="static/iShot_2024-04-26_22.11.44.png"/></td>
    </tr>
    <tr>
        <td><img src="static/iShot_2024-04-26_22.11.44.png"/></td>
        <td><img src="static/iShot_2024-04-26_22.12.20.png"/></td>
    </tr>
    <tr>
        <td><img src="static/iShot_2024-04-26_22.12.40.png"/></td>
        <td><img src="static/iShot_2024-04-26_22.13.14.png"/></td>
    </tr>
</table>

### 代办
- [ ] consul 多集群部署
- [ ] k8s容器编排管理


