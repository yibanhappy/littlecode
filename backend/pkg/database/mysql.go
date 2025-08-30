package database

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"littlecode-backend/internal/config"
	"littlecode-backend/internal/models"
)

var DB *gorm.DB

func InitDB(cfg *config.Config) error {
	// 首先尝试连接MySQL服务器（不指定数据库）来创建数据库
	dsnWithoutDB := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
	)

	// 连接MySQL服务器
	tempDB, err := gorm.Open(mysql.Open(dsnWithoutDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL server: %w", err)
	}

	// 创建数据库（如果不存在）
	createDBSQL := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", cfg.Database.DBName)
	if err := tempDB.Exec(createDBSQL).Error; err != nil {
		return fmt.Errorf("failed to create database: %w", err)
	}

	log.Printf("Database '%s' created or already exists", cfg.Database.DBName)

	// 现在连接到指定的数据库
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	log.Println("Database connected successfully")

	// 自动迁移数据库表
	if err := autoMigrate(); err != nil {
		return fmt.Errorf("failed to migrate database: %w", err)
	}

	return nil
}

func autoMigrate() error {
	return DB.AutoMigrate(
		&models.User{},
		&models.Memo{},
		&models.Timer{},
	)
}

func GetDB() *gorm.DB {
	return DB
}
