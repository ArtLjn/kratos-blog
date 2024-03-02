---
title: Blog v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.17"

---

# Blog

> v1.0.0

Base URLs:

* <a href="http://localhost:8000">blog: http://localhost:8000</a>

* <a href="http://localhost:8001">user: http://localhost:8001</a>

* <a href="http://localhost:8002">comment: http://localhost:8002</a>

# Default

## GET 查询所有图片接口

GET /api/getAllPhoto

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除图片接口

DELETE /api/deletePhoto/94/9977c000-9f09-42c4-b29a-ce8b845b8056

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 查询所有推荐文章接口

GET /api/getAllSuggest

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# blog

## GET 获取所有文章

GET /api/getAllBlog

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新文章权限

PUT /api/updateOnly

> Body 请求参数

```json
{
  "raw": 1,
  "id": 1,
  "res": true
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 根据tagName查询文章

GET /api/getTagName/fisco-bcos

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 创建博客接口

POST /api/addBlog

> Body 请求参数

```json
{
  "content": "hhhh",
  "title": "hhhhh",
  "preface": "hhhh",
  "photo": "http://localhost:10002/img/423ec5bd-2a50-4cd5-8a71-0eae1e310060.jpg",
  "tag": "区块链",
  "comment": "0"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新博客接口

PUT /api/updateBlog/895

> Body 请求参数

```json
{
  "content": "ljnnb666tai",
  "title": "gggggggggggggggg",
  "preface": "gggggggggggggg",
  "photo": "http://hhhhhhhhhhhhhhhhhhhhhhhh:10002/img/423ec5bd-2a50-hhhh-8a71-0eae1e310060.jpg",
  "tag": "dgsfgsagfsgsg",
  "comment": "0"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取Bing每日图片

GET /api/getBingPhoto

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 根据文章ID查询

GET /api/getId/1

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 上传图片

POST /util/upload

> Body 请求参数

```yaml
file: file://D:\照片\c4220a1e-7d3b-4048-b544-24f359f37e10.jpg

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» file|body|string(binary)| 是 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 根据文章标题查询文章

GET /api/searchBlog/java

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新所有文章的评论或者可见性权限

PUT /api/updateIndividual

> Body 请求参数

```json
{
  "raw": 0,
  "status": true
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加草稿笔记

POST /api/cacheBlog

> Body 请求参数

```json
{
  "data": {
    "content": "hhhh",
    "title": "hhhhh",
    "preface": "hhhh",
    "photo": "http://localhost:10002/img/423ec5bd-2a50-4cd5-8a71-0eae1e310060.jpg",
    "tag": "区块链"
  }
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 获取所有草稿笔记

GET /api/getCacheBlog

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除缓存草稿

DELETE /api/deleteCacheBlog

> Body 请求参数

```json
{
  "key": "2024-02-03 11:39-hhhhh"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# comment

## GET 查询文章对应的所有评论信息

GET /api/getComment/talk

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加评论接口

POST /api/addComment

> Body 请求参数

```json
{
  "article_id": "876",
  "comment": "hello",
  "comment_addr": "58.216.213.53"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加回复接口

POST /api/addReward/923

> Body 请求参数

```json
{
  "reward_content": "fdfasfddf",
  "article_id": "876",
  "reward_addr": "58.216.211.183"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# user

## POST 用户登录接口

POST /api/login

> Body 请求参数

```json
{
  "name": "ljn",
  "pass": "123"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 发送邮件接口

GET /api/sendEmail/915040912%40qq.com

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 注册接口

POST /api/register/480396

> Body 请求参数

```json
{
  "name": "ljn",
  "password": "123456",
  "email": "915040912@qq.com"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 管理员登录接口

POST /api/admin

> Body 请求参数

```json
{
  "name": "liujunnan",
  "password": "320324liu"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 修改密码接口

POST /api/updatePassword/866113

> Body 请求参数

```json
{
  "name": "test",
  "password": "123",
  "email": "915040912@qq.com"
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 设置用户黑名单

GET /api/setBlack/test

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 查询用户信息

GET /api/getUserMessage/test

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# tag

## GET 查询所有标签

GET /api/getAllTag

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## POST 添加标签

POST /api/addTag

> Body 请求参数

```json
{
  "data": {
    "tagName": "test"
  }
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除标签

DELETE /api/deleteTag/1

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# friend

## POST 添加友链接口

POST /api/addFriend

> Body 请求参数

```json
{
  "data": {
    "title": "Test",
    "preface": "fdfsdfd",
    "url": "fdfsdadff",
    "photo": "fdfsafdfadf"
  }
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 查询所有友链接口

GET /api/getAllFriend

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## DELETE 删除友链接口

DELETE /api/deleteFriend/1

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## GET 根据ID查询友链记录

GET /api/getFriend/1

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

## PUT 更新友链记录

PUT /api/updateFriend

> Body 请求参数

```json
{
  "data": {
    "id": 1,
    "title": "test",
    "preface": "ggg",
    "url": "ssss",
    "photo": "fdfsafdfadf"
  }
}
```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|token|header|string| 否 |none|
|body|body|object| 否 |none|

> 返回示例

> 200 Response

```json
{}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|成功|Inline|

### 返回数据结构

# 数据模型

