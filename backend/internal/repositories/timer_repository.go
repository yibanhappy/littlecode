package repositories

import (
	"gorm.io/gorm"

	"littlecode-backend/internal/models"
)

type TimerRepository struct {
	db *gorm.DB
}

func NewTimerRepository(db *gorm.DB) *TimerRepository {
	return &TimerRepository{db: db}
}

// 获取用户的所有计时器
func (r *TimerRepository) GetByUserID(userID uint) ([]models.Timer, error) {
	var timers []models.Timer
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&timers).Error
	return timers, err
}

// 根据ID获取计时器
func (r *TimerRepository) GetByID(id uint, userID uint) (*models.Timer, error) {
	var timer models.Timer
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&timer).Error
	if err != nil {
		return nil, err
	}
	return &timer, nil
}

// 创建计时器
func (r *TimerRepository) Create(timer *models.Timer) error {
	return r.db.Create(timer).Error
}

// 更新计时器
func (r *TimerRepository) Update(timer *models.Timer) error {
	return r.db.Save(timer).Error
}

// 删除计时器
func (r *TimerRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Timer{}).Error
}
