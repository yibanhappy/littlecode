-- 创建用户
CREATE USER IF NOT EXISTS 'lcuser'@'localhost' IDENTIFIED BY 'lc7702';

-- 授予权限
GRANT ALL PRIVILEGES ON littlecode.* TO 'lcuser'@'localhost';

-- 如果数据库不存在，创建数据库
CREATE DATABASE IF NOT EXISTS littlecode CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- 再次授予权限确保生效
GRANT ALL PRIVILEGES ON littlecode.* TO 'lcuser'@'localhost';

-- 刷新权限
FLUSH PRIVILEGES;

-- 显示创建的用户
SELECT User, Host FROM mysql.user WHERE User = 'lcuser';
