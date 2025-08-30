package services

import (
	"errors"
	"time"

	"littlecode-backend/internal/models"
	"littlecode-backend/internal/repositories"
)

type TimerService struct {
	timerRepo *repositories.TimerRepository
}

type CreateTimerRequest struct {
	Name     string `json:"name" validate:"required"`
	Duration int    `json:"duration"` // 持续时间（秒）
}

type UpdateTimerRequest struct {
	Name     string `json:"name" validate:"required"`
	Duration int    `json:"duration"`
	Status   string `json:"status"`
}

func NewTimerService(timerRepo *repositories.TimerRepository) *TimerService {
	return &TimerService{
		timerRepo: timerRepo,
	}
}

// 获取用户的所有计时器
func (s *TimerService) GetTimersByUserID(userID uint) ([]models.Timer, error) {
	return s.timerRepo.GetByUserID(userID)
}

// 根据ID获取计时器
func (s *TimerService) GetTimerByID(id uint, userID uint) (*models.Timer, error) {
	return s.timerRepo.GetByID(id, userID)
}

// 创建计时器
func (s *TimerService) CreateTimer(req CreateTimerRequest, userID uint) (*models.Timer, error) {
	timer := &models.Timer{
		UserID:   userID,
		Name:     req.Name,
		Duration: req.Duration,
		Status:   "stopped",
	}

	err := s.timerRepo.Create(timer)
	if err != nil {
		return nil, err
	}

	return timer, nil
}

// 更新计时器
func (s *TimerService) UpdateTimer(id uint, req UpdateTimerRequest, userID uint) (*models.Timer, error) {
	timer, err := s.timerRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}
	if timer == nil {
		return nil, errors.New("计时器不存在")
	}

	timer.Name = req.Name
	timer.Duration = req.Duration
	if req.Status != "" {
		timer.Status = req.Status
	}

	err = s.timerRepo.Update(timer)
	if err != nil {
		return nil, err
	}

	return timer, nil
}

// 启动计时器
func (s *TimerService) StartTimer(id uint, userID uint) (*models.Timer, error) {
	timer, err := s.timerRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}
	if timer == nil {
		return nil, errors.New("计时器不存在")
	}

	now := time.Now()
	timer.Status = "running"
	timer.StartTime = &now

	err = s.timerRepo.Update(timer)
	if err != nil {
		return nil, err
	}

	return timer, nil
}

// 停止计时器
func (s *TimerService) StopTimer(id uint, userID uint) (*models.Timer, error) {
	timer, err := s.timerRepo.GetByID(id, userID)
	if err != nil {
		return nil, err
	}
	if timer == nil {
		return nil, errors.New("计时器不存在")
	}

	now := time.Now()
	timer.Status = "stopped"
	timer.EndTime = &now

	err = s.timerRepo.Update(timer)
	if err != nil {
		return nil, err
	}

	return timer, nil
}

// 删除计时器
func (s *TimerService) DeleteTimer(id uint, userID uint) error {
	return s.timerRepo.Delete(id, userID)
}
