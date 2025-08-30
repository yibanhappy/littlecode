package repositories

import (
	"errors"

	"gorm.io/gorm"

	"littlecode-backend/internal/models"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

// 根据账号查找用户
func (r *UserRepository) GetByAccount(account string) (*models.User, error) {
	var user models.User
	err := r.db.Where("account = ?", account).First(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// 根据ID查找用户
func (r *UserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

// 创建用户
func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

// 更新用户
func (r *UserRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

// 删除用户
func (r *UserRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}
