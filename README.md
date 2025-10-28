# Go 博客系统

一个基于 Go 语言开发的简单博客系统，支持用户管理、文章发布、评论功能。

## 🚀 技术栈

- **Web 框架**: Gin (v1.11.0)
- **数据库**: MySQL + GORM (v1.31.0)
- **身份认证**: JWT (v5.3.0)
- **API 文档**: Swagger
- **Go 版本**: 1.24.0

## 📁 项目结构

```
go_blog_work_project/
├── db/                    # 数据库相关
│   ├── config/           # 数据库配置
│   ├── model/            # 数据模型
│   └── repository/       # 数据访问层
├── gin/                  # Web 层
│   ├── biz/              # 业务逻辑层
│   ├── middleware/       # 中间件
│   ├── router/           # 路由配置
│   └── docs/             # Swagger 文档
├── main.go               # 程序入口
└── go.mod                # 依赖管理
```

## 🛠️ 核心功能

### 数据模型
- **User**: 用户信息 (用户名、密码、邮箱)
- **Post**: 文章信息 (标题、内容、作者)
- **Comment**: 评论信息 (内容、用户、文章)

### 中间件
- **PanicCatch**: 全局异常捕获中间件
- **JWT**: JWT 身份认证中间件 (暂未启用)

## 📡 API 接口

### 用户相关 (`/user`)
- `GET /user/login` - 用户登录
- `GET /user/register` - 用户注册

### 文章相关 (`/post`)
- `POST /post/create` - 创建文章
- `GET /post/query` - 查询文章列表
- `POST /post/update` - 更新文章
- `DELETE /post/delete` - 删除文章

### 评论相关 (`/comment`)
- `POST /comment/create` - 创建评论
- `GET /comment/getListByPostId` - 根据文章ID查询评论

### 文档
- `GET /swagger/*` - Swagger API 文档
- `GET /swagger.json` - Swagger JSON 文件

## 🚀 快速开始

### 环境要求
- Go 1.24.0+
- MySQL 5.7+

### 安装依赖
```bash
go mod tidy
```

### 配置数据库
修改 `db/config/init.go` 中的数据库连接信息：
```go
dsn := "root:password@tcp(localhost:3306)/blog_db?charset=utf8mb4&parseTime=True&loc=Local"
```

### 运行项目
```bash
go run main.go
```

服务启动后访问：
- API 服务: http://localhost:8080
- Swagger 文档: http://localhost:8080/swagger/index.html

## 📝 使用说明

### 创建用户
```bash
curl "http://localhost:8080/user/register?username=test&password=123456"
```

### 用户登录
```bash
curl "http://localhost:8080/user/login?username=test&password=123456"
```

### 创建文章
```bash
curl -X POST "http://localhost:8080/post/create" \
  -H "Content-Type: application/json" \
  -d '{"title":"测试文章","content":"这是文章内容","user_id":1}'
```

### 查询文章
```bash
curl "http://localhost:8080/post/query"
```

### 创建评论
```bash
curl -X POST "http://localhost:8080/comment/create" \
  -H "Content-Type: application/json" \
  -d '{"content":"这是评论内容","user_id":1,"post_id":1}'
```

## ⚠️ 注意事项

1. **数据库**: 项目会自动创建表结构，请确保数据库连接正确
2. **JWT 认证**: 当前 JWT 中间件已实现但未启用，需要手动添加到路由中
3. **错误处理**: 使用 panic 进行错误处理，生产环境建议改为更优雅的错误处理方式
4. **密码安全**: 当前密码以明文存储，生产环境请使用加密存储

## 🔧 开发说明

这是一个学习项目，主要用于：
- Go Web 开发学习
- RESTful API 设计实践
- 数据库操作学习
- 中间件使用学习

适合 Go 语言初学者和 Web 开发学习者使用。
