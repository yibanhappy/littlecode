package repositories

import (
	"gorm.io/gorm"

	"littlecode-backend/internal/models"
)

type MemoRepository struct {
	db *gorm.DB
}

func NewMemoRepository(db *gorm.DB) *MemoRepository {
	return &MemoRepository{db: db}
}

// 获取用户的所有备忘录
func (r *MemoRepository) GetByUserID(userID uint) ([]models.Memo, error) {
	var memos []models.Memo
	err := r.db.Where("user_id = ?", userID).Order("created_at DESC").Find(&memos).Error
	return memos, err
}

// 根据ID获取备忘录
func (r *MemoRepository) GetByID(id uint, userID uint) (*models.Memo, error) {
	var memo models.Memo
	err := r.db.Where("id = ? AND user_id = ?", id, userID).First(&memo).Error
	if err != nil {
		return nil, err
	}
	return &memo, nil
}

// 创建备忘录
func (r *MemoRepository) Create(memo *models.Memo) error {
	return r.db.Create(memo).Error
}

// 更新备忘录
func (r *MemoRepository) Update(memo *models.Memo) error {
	return r.db.Save(memo).Error
}

// 删除备忘录
func (r *MemoRepository) Delete(id uint, userID uint) error {
	return r.db.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Memo{}).Error
}
