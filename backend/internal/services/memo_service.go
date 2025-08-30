package services

import (
	"errors"
	"littlecode-backend/internal/models"
	"littlecode-backend/internal/repositories"
)

type MemoService struct {
	memoRepo *repositories.MemoRepository
}

type CreateMemoRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content"`
}

type UpdateMemoRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content"`
}

func NewMemoService(memoRepo *repositories.MemoRepository) *MemoService {
	return &MemoService{
		memoRepo: memoRepo,
	}
}

// 获取用户的所有备忘录
func (s *MemoService) GetMemosByUserID(userID uint) ([]models.Memo, error) {
	return s.memoRepo.GetByUserID(userID)
}

// 根据ID获取备忘录
func (s *MemoService) GetMemoByID(id uint, userID uint) (*models.Memo, error) {
	return s.memoRepo.GetByID(id, userID)
}

// 创建备忘录
func (s *MemoService) CreateMemo(req CreateMemoRequest, userID uint) (*models.Memo, error) {
	memo := &models.Memo{
		UserID:  userID,
		Title:   req.Title,
		Content: req.Content,
	}

	err := s.memoRepo.Create(memo)
	if err != nil {
		return nil, err
	}

	return memo, nil
}

// 更新备忘录
func (s *MemoService) UpdateMemo(id uint, req UpdateMemoRequest, userID uint) (*models.Memo, error) {
	memo, err := s.memoRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}
	if memo == nil {
		return nil, errors.New("备忘录不存在")
	}

	memo.Title = req.Title
	memo.Content = req.Content

	err = s.memoRepo.Update(memo)
	if err != nil {
		return nil, err
	}

	return memo, nil
}

// 删除备忘录
func (s *MemoService) DeleteMemo(id uint, userID uint) error {
	return s.memoRepo.Delete(id, userID)
}
