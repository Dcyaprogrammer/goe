# GoE - Go Gin + PostgreSQL 项目

这是一个使用 Go Gin 框架和 PostgreSQL 数据库的后端项目模板。

## 项目结构

```
goe/
├── main.go              # 主程序入口
├── go.mod               # Go 模块文件
├── go.sum               # 依赖校验文件
├── env.example          # 环境变量示例
├── README.md            # 项目说明
├── config/              # 配置文件
│   └── database.go      # 数据库配置
├── models/              # 数据模型
│   └── user.go          # 用户模型
├── handlers/            # 请求处理器
│   └── user_handler.go  # 用户处理器
└── routes/              # 路由配置
    └── routes.go        # 路由设置
```

## 功能特性

- ✅ Gin Web 框架
- ✅ PostgreSQL 数据库
- ✅ GORM ORM
- ✅ 用户 CRUD 操作
- ✅ RESTful API
- ✅ 环境变量配置
- ✅ CORS 支持

## 快速开始

### 1. 安装依赖

```bash
go mod tidy
```

### 2. 配置环境变量

复制环境变量示例文件：

```bash
cp env.example .env
```

然后编辑 `.env` 文件，配置你的数据库连接信息。

### 3. 启动 PostgreSQL 数据库

确保 PostgreSQL 服务正在运行，并创建数据库：

```sql
CREATE DATABASE goe_db;
```

### 4. 运行项目

```bash
go run main.go
```

服务器将在 `http://localhost:8080` 启动。

## API 接口

### 健康检查
- `GET /api/v1/health` - 检查服务器状态

### 用户管理
- `GET /api/v1/users` - 获取所有用户
- `GET /api/v1/users/:id` - 获取指定用户
- `POST /api/v1/users` - 创建新用户
- `PUT /api/v1/users/:id` - 更新用户
- `DELETE /api/v1/users/:id` - 删除用户

## 环境变量

| 变量名 | 描述 | 默认值 |
|--------|------|--------|
| PORT | 服务器端口 | 8080 |
| GIN_MODE | Gin 运行模式 | debug |
| DB_HOST | 数据库主机 | localhost |
| DB_PORT | 数据库端口 | 5432 |
| DB_USER | 数据库用户名 | postgres |
| DB_PASSWORD | 数据库密码 | password |
| DB_NAME | 数据库名称 | goe_db |
| JWT_SECRET | JWT 密钥 | your-secret-key |

## 开发

### 添加新的模型

1. 在 `models/` 目录下创建新的模型文件
2. 在 `handlers/` 目录下创建对应的处理器
3. 在 `routes/routes.go` 中添加新的路由

### 数据库迁移

项目使用 GORM 的自动迁移功能。在 `main.go` 中可以添加：

```go
// 自动迁移数据库表
db.AutoMigrate(&models.User{})
```

## 许可证

MIT License 