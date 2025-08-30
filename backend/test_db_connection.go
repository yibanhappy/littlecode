package main

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 从环境变量或命令行参数获取连接信息
	host := getEnvOrDefault("DB_HOST", "localhost")
	port := getEnvOrDefault("DB_PORT", "3306")
	user := getEnvOrDefault("DB_USER", "root")
	password := getEnvOrDefault("DB_PASSWORD", "")

	fmt.Printf("尝试连接MySQL服务器:\n")
	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Port: %s\n", port)
	fmt.Printf("User: %s\n", user)
	fmt.Printf("Password: %s\n", maskPassword(password))

	// 测试连接（不指定数据库）
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Printf("❌ 连接失败: %v\n", err)
		fmt.Println("\n请检查以下事项:")
		fmt.Println("1. MySQL服务是否正在运行")
		fmt.Println("2. 用户名和密码是否正确")
		fmt.Println("3. 主机和端口是否正确")
		os.Exit(1)
	}

	fmt.Println("✅ MySQL服务器连接成功!")

	// 获取数据库连接以执行SQL
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("❌ 获取数据库连接失败: %v\n", err)
		os.Exit(1)
	}
	defer sqlDB.Close()

	// 测试查询
	var version string
	err = db.Raw("SELECT VERSION()").Scan(&version).Error
	if err != nil {
		fmt.Printf("❌ 查询版本失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ MySQL版本: %s\n", version)

	// 尝试创建数据库
	dbName := "littlecode"
	createSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName)
	err = db.Exec(createSQL).Error
	if err != nil {
		fmt.Printf("❌ 创建数据库失败: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✅ 数据库 '%s' 创建成功或已存在\n", dbName)

	// 测试连接到指定数据库
	dsnWithDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, password, host, port, dbName)

	_, err = gorm.Open(mysql.Open(dsnWithDB), &gorm.Config{})
	if err != nil {
		fmt.Printf("❌ 连接到数据库 '%s' 失败: %v\n", dbName, err)
		os.Exit(1)
	}

	fmt.Printf("✅ 成功连接到数据库 '%s'\n", dbName)

	fmt.Println("\n🎉 所有测试通过! 数据库配置正确。")
}

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func maskPassword(password string) string {
	if len(password) == 0 {
		return "(空密码)"
	}
	if len(password) <= 2 {
		return "***"
	}
	return password[:1] + "***" + password[len(password)-1:]
}
