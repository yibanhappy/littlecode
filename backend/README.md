# LittleCode Backend

基于Echo框架和MySQL的后端API服务，为前端uni-app应用提供数据支持。

## 环境准备

### 1. Go语言环境安装

#### Windows系统安装Go

1. **下载Go安装包**
   - 访问 [Go官方下载页面](https://golang.org/dl/)
   - 下载适合Windows的安装包（如：go1.21.x.windows-amd64.msi）

2. **安装Go**
   - 双击下载的.msi文件
   - 按照安装向导完成安装（默认安装到 `C:\Program Files\Go`）

3. **验证安装**
   ```powershell
   go version
   ```
   应该输出类似：`go version go1.21.x windows/amd64`

#### macOS系统安装Go

```bash
# 使用Homebrew安装
brew install go

# 或者下载安装包
# 访问 https://golang.org/dl/ 下载 .pkg 文件并安装
```

#### Linux系统安装Go

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install golang-go

# CentOS/RHEL
sudo yum install golang

# 或者手动下载安装
wget https://golang.org/dl/go1.21.x.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.21.x.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

### 2. Go环境配置

#### 设置GOPATH和GOROOT（可选，Go 1.11+使用模块化）

```powershell
# Windows PowerShell
$env:GOROOT = "C:\Program Files\Go"
$env:GOPATH = "C:\Users\你的用户名\go"
$env:PATH += ";$env:GOROOT\bin;$env:GOPATH\bin"

# 或者在系统环境变量中设置
```

```bash
# Linux/macOS
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

#### 配置Go代理（重要 - 解决网络问题）

由于网络原因，建议设置国内代理镜像源：

```powershell
# Windows PowerShell
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn

# 或者使用阿里云镜像
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```

```bash
# Linux/macOS
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn

# 或者添加到 ~/.bashrc 或 ~/.zshrc
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=sum.golang.google.cn
```

**常用的Go代理镜像源：**
- 官方代理：`https://proxy.golang.org` (可能需要梯子)
- 七牛云：`https://goproxy.cn`
- 阿里云：`https://mirrors.aliyun.com/goproxy/`
- 百度：`https://goproxy.bj.bcebos.com/`

### 3. MySQL数据库安装

#### Windows安装MySQL

1. 下载MySQL Installer：访问 [MySQL官网](https://dev.mysql.com/downloads/installer/)
2. 选择"MySQL Installer for Windows"
3. 运行安装程序，选择"Developer Default"
4. 设置root密码并完成安装

#### macOS安装MySQL

```bash
# 使用Homebrew
brew install mysql

# 启动MySQL服务
brew services start mysql

# 安全设置
mysql_secure_installation
```

#### Linux安装MySQL

```bash
# Ubuntu/Debian
sudo apt update
sudo apt install mysql-server

# CentOS/RHEL
sudo yum install mysql-server

# 启动并设置开机自启
sudo systemctl start mysqld
sudo systemctl enable mysqld
```

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

### 1. 克隆项目

```bash
git clone <your-repo-url>
cd littlecode/backend
```

### 2. 配置环境变量

复制并修改环境配置文件：

```bash
# 检查 .env 文件，修改数据库配置
```

`.env` 文件示例：
```env
# 数据库配置
DB_HOST=localhost
DB_PORT=3306
DB_USER=lcuser
DB_PASSWORD=lc7702
DB_NAME=littlecode

# 服务配置
SERVER_PORT=8079
SERVER_HOST=localhost

# JWT配置
JWT_SECRET=your_jwt_secret_key_here
JWT_EXPIRE_HOURS=24

# 环境配置
ENV=development
```

### 3. 数据库准备

#### 方法一：自动创建（推荐）
应用程序会自动创建数据库和用户，您只需要确保MySQL服务运行并且有管理员权限。

#### 方法二：手动创建
连接到MySQL并执行以下命令：

```sql
-- 创建数据库
CREATE DATABASE IF NOT EXISTS littlecode CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 创建用户
CREATE USER 'lcuser'@'localhost' IDENTIFIED BY 'lc7702';

-- 授予权限
GRANT ALL PRIVILEGES ON littlecode.* TO 'lcuser'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;
```

### 4. 安装依赖

```bash
# 下载依赖
go mod download

# 整理依赖
go mod tidy
```

### 5. 运行应用

```bash
# 开发模式运行
go run cmd/main.go

# 或者使用 Makefile
make run
```

成功启动后，您应该看到类似输出：
```
2025/08/30 16:48:50 Database 'littlecode' created or already exists
2025/08/30 16:48:50 Database connected successfully
2025/08/30 16:48:51 Server starting on http://localhost:8079
⇨ http server started on [::]:8079
```

### 6. 测试API

使用浏览器或Postman访问：
- 健康检查：`http://localhost:8079/api/health`（如果有的话）
- API文档：查看下方的API接口说明

## 构建和部署

### 本地构建

```bash
# 构建二进制文件
go build -o littlecode-backend cmd/main.go

# 运行构建的文件
./littlecode-backend.exe  # Windows
./littlecode-backend      # Linux/macOS
```

### 使用Makefile

```bash
# 构建
make build

# 运行
make run

# 清理
make clean

# 下载依赖
make deps
```

## 常见问题

### 1. Go代理问题

如果遇到 `dial tcp xxx: i/o timeout` 错误：

```bash
# 切换代理源
go env -w GOPROXY=https://goproxy.cn,direct

# 或者临时使用
export GOPROXY=https://goproxy.cn,direct
```

### 2. 数据库连接失败

- 检查MySQL服务是否启动
- 验证用户名密码是否正确
- 确认端口号是否正确（默认3306）
- 检查防火墙设置

### 3. 端口被占用

```bash
# Windows查看端口占用
netstat -ano | findstr :8079

# Linux/macOS查看端口占用
lsof -i :8079
```

修改 `.env` 文件中的 `SERVER_PORT` 配置。

### 4. 权限问题

确保用户具有数据库操作权限：
```sql
SHOW GRANTS FOR 'lcuser'@'localhost';
```

## API接口文档

### 基础信息

- **Base URL**: `http://localhost:8079`
- **API前缀**: `/api`
- **认证方式**: JWT Bearer Token

### 认证接口

#### 用户注册
- **URL**: `POST /api/register`
- **描述**: 用户注册
- **请求体**:
```json
{
    "account": "string",
    "password": "string",
    "confirm_password": "string"
}
```
- **响应**:
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "id": 1,
        "account": "testuser",
        "created_at": "2025-08-30T16:48:51Z",
        "updated_at": "2025-08-30T16:48:51Z"
    }
}
```

#### 用户登录
- **URL**: `POST /api/login`
- **描述**: 用户登录
- **请求体**:
```json
{
    "account": "string",
    "password": "string"
}
```
- **响应**:
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
        "user": {
            "id": 1,
            "account": "testuser",
            "created_at": "2025-08-30T16:48:51Z",
            "updated_at": "2025-08-30T16:48:51Z"
        }
    }
}
```

### 用户接口

#### 获取用户信息
- **URL**: `GET /api/profile`
- **描述**: 获取当前用户信息
- **认证**: 需要Bearer Token
- **响应**:
```json
{
    "code": 200,
    "message": "success",
    "data": {
        "id": 1,
        "account": "testuser",
        "created_at": "2025-08-30T16:48:51Z",
        "updated_at": "2025-08-30T16:48:51Z"
    }
}
```

### 备忘录接口

#### 获取备忘录列表
- **URL**: `GET /api/memos`
- **描述**: 获取当前用户的所有备忘录
- **认证**: 需要Bearer Token

#### 获取单个备忘录
- **URL**: `GET /api/memos/:id`
- **描述**: 获取指定ID的备忘录
- **认证**: 需要Bearer Token

#### 创建备忘录
- **URL**: `POST /api/memos`
- **描述**: 创建新的备忘录
- **认证**: 需要Bearer Token
- **请求体**:
```json
{
    "title": "string",
    "content": "string"
}
```

#### 更新备忘录
- **URL**: `PUT /api/memos/:id`
- **描述**: 更新指定备忘录
- **认证**: 需要Bearer Token
- **请求体**:
```json
{
    "title": "string",
    "content": "string"
}
```

#### 删除备忘录
- **URL**: `DELETE /api/memos/:id`
- **描述**: 删除指定备忘录
- **认证**: 需要Bearer Token

### 计时器接口

#### 获取计时器列表
- **URL**: `GET /api/timers`
- **描述**: 获取当前用户的所有计时器
- **认证**: 需要Bearer Token

#### 获取单个计时器
- **URL**: `GET /api/timers/:id`
- **描述**: 获取指定ID的计时器
- **认证**: 需要Bearer Token

#### 创建计时器
- **URL**: `POST /api/timers`
- **描述**: 创建新的计时器
- **认证**: 需要Bearer Token
- **请求体**:
```json
{
    "name": "string",
    "duration": 3600
}
```

#### 更新计时器
- **URL**: `PUT /api/timers/:id`
- **描述**: 更新指定计时器
- **认证**: 需要Bearer Token

#### 启动计时器
- **URL**: `POST /api/timers/:id/start`
- **描述**: 启动指定计时器
- **认证**: 需要Bearer Token

#### 停止计时器
- **URL**: `POST /api/timers/:id/stop`
- **描述**: 停止指定计时器
- **认证**: 需要Bearer Token

#### 删除计时器
- **URL**: `DELETE /api/timers/:id`
- **描述**: 删除指定计时器
- **认证**: 需要Bearer Token

### 认证说明

所有需要认证的接口都需要在请求头中包含JWT Token：

```
Authorization: Bearer <your-jwt-token>
```

### 错误码说明

- `200`: 成功
- `400`: 请求参数错误
- `401`: 未授权（Token无效或过期）
- `404`: 资源不存在
- `500`: 服务器内部错误

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

### 数据库迁移

应用程序使用GORM的AutoMigrate功能自动创建和更新数据库表结构。当您修改模型时，重启应用程序即可自动应用更改。

## 技术栈

- **Web框架**: Echo v4
- **ORM**: GORM
- **数据库**: MySQL
- **认证**: JWT
- **配置管理**: godotenv
- **密码加密**: bcrypt

## 开发工具推荐

- **IDE**: VS Code, GoLand
- **API测试**: Postman, Insomnia
- **数据库管理**: MySQL Workbench, phpMyAdmin
- **版本控制**: Git

## 许可证

MIT License
