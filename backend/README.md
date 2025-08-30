# LittleCode Backend

基于Echo框架和MySQL的后端API服务，为前端uni-app应用提供数据支持。

## 项目结构

```
backend/
├── cmd/                    # 应用程序入口点
│   └── main.go
├── internal/              # 内部应用程序代码
│   ├── config/           # 配置管理
│   ├── handlers/         # HTTP处理器
│   ├── middleware/       # 中间件
│   ├── models/          # 数据模型
│   ├── repositories/    # 数据访问层
│   └── services/        # 业务逻辑层
├── pkg/                  # 可重用的包
│   ├── database/        # 数据库连接
│   └── utils/           # 工具函数
├── .env                 # 环境变量配置
├── go.mod              # Go模块文件
├── go.sum              # Go依赖校验文件
├── Makefile            # 构建脚本
└── README.md           # 项目文档
```

## 功能特性

- **用户认证**: 用户注册、登录，JWT token认证
- **备忘录管理**: 创建、查看、更新、删除备忘录
- **计时器管理**: 创建、启动、停止、删除计时器
- **RESTful API**: 标准的REST接口设计
- **数据库**: MySQL数据存储，GORM ORM
- **中间件**: CORS、日志、认证等中间件支持

## 快速开始

### 环境要求

- Go 1.19+
- MySQL 8.0+

### 安装依赖

```bash
go mod download
```

### 配置数据库

1. 创建MySQL数据库：
```sql
CREATE DATABASE littlecode CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
```

2. 修改 `.env` 文件中的数据库配置：
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=littlecode
```

### 运行应用

```bash
# 开发模式运行
go run cmd/main.go

# 或者使用 Makefile
make run
```

### 构建应用

```bash
# 构建二进制文件
make build

# 运行构建的文件
./littlecode-backend
```

## API接口

### 认证接口

- `POST /api/register` - 用户注册
- `POST /api/login` - 用户登录

### 用户接口

- `GET /api/profile` - 获取用户信息

### 备忘录接口

- `GET /api/memos` - 获取备忘录列表
- `GET /api/memos/:id` - 获取单个备忘录
- `POST /api/memos` - 创建备忘录
- `PUT /api/memos/:id` - 更新备忘录
- `DELETE /api/memos/:id` - 删除备忘录

### 计时器接口

- `GET /api/timers` - 获取计时器列表
- `GET /api/timers/:id` - 获取单个计时器
- `POST /api/timers` - 创建计时器
- `PUT /api/timers/:id` - 更新计时器
- `POST /api/timers/:id/start` - 启动计时器
- `POST /api/timers/:id/stop` - 停止计时器
- `DELETE /api/timers/:id` - 删除计时器

## 开发指南

### 项目架构

项目采用分层架构设计：

1. **Handler层**: 处理HTTP请求和响应
2. **Service层**: 业务逻辑处理
3. **Repository层**: 数据访问抽象
4. **Model层**: 数据模型定义

### 添加新功能

1. 在 `models` 中定义数据模型
2. 在 `repositories` 中实现数据访问
3. 在 `services` 中实现业务逻辑
4. 在 `handlers` 中实现HTTP处理器
5. 在 `main.go` 中注册路由

## 技术栈

- **Web框架**: Echo v4
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **配置管理**: godotenv
- **密码加密**: bcrypt

## 许可证

MIT License
