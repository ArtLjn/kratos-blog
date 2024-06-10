# Kratos-Blog 微服务博客系统

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


## 技术栈
- mysql
- redis
- rabbitmq
- consul
- gin
- jwt

## 功能特点

1. **文章管理**：用户可以创建、编辑和删除文章。支持 MarkDown 编辑器，可以插入图片、链接等多媒体元素。
2. **用户认证**：支持用户注册和登录功能，对用户的角色进行区分。
3. **评论功能**：用户可以在文章下发表评论，并查看其他用户的评论。
4. **分类和标签**：文章可以按照不同的分类和标签进行组织，方便用户查找和浏览相关的文章。
5. **搜索功能**：提供全文搜索功能，帮助用户快速找到所需的文章内容。
6. **响应式设计**：界面采用响应式设计，适应不同分辨率的设备，如手机、平板和桌面电脑。

## 技术架构

Kratos-blog 采用 go-kratos 微服务框架作为后端核心，实现了高并发、可扩展的服务端逻辑。具体来说，它包括以下几个部分：

1. **API 网关**：负责处理来自客户端的请求，将请求路由到相应的微服务上，并提供统一的入口地址。
2. **用户服务**：处理用户相关的操作，如注册、登录、用户信息管理等。
3. **文章服务**：负责文章的创建、编辑、删除等操作，以及文章的分类、标签管理。
4. **评论服务**：处理用户对文章的评论功能，包括添加评论、回复评论等。
5. **搜索服务**：提供全文搜索功能，帮助用户快速找到所需的文章内容。
6. **数据库**：使用高效的数据库存储用户数据、文章数据和评论数据，保证数据的持久化存储。
7. **缓存**：使用缓存技术提升系统的响应速度和性能，减少数据库的访问压力。
8. **日志和监控**：记录系统运行过程中的关键信息，并进行实时监控和告警。

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



