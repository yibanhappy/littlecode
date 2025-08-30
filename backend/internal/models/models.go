package models

import (
	"time"

	"gorm.io/gorm"
)

// 用户模型
type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Account   string         `json:"account" gorm:"uniqueIndex;not null;size:100"`
	Password  string         `json:"-" gorm:"not null"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

// 备忘录模型
type Memo struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	Title     string         `json:"title" gorm:"not null"`
	Content   string         `json:"content" gorm:"type:text"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 外键关联
	User User `json:"user" gorm:"foreignKey:UserID"`
}

// 计时器模型
type Timer struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	UserID    uint           `json:"user_id" gorm:"not null;index"`
	Name      string         `json:"name" gorm:"not null"`
	Duration  int            `json:"duration"`                        // 持续时间（秒）
	Status    string         `json:"status" gorm:"default:'stopped'"` // running, paused, stopped
	StartTime *time.Time     `json:"start_time"`
	EndTime   *time.Time     `json:"end_time"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`

	// 外键关联
	User User `json:"user" gorm:"foreignKey:UserID"`
}
