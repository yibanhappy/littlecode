package services

import (
	"errors"

	"littlecode-backend/internal/config"
	"littlecode-backend/internal/models"
	"littlecode-backend/internal/repositories"
	"littlecode-backend/pkg/utils"
)

type UserService struct {
	userRepo *repositories.UserRepository
	config   *config.Config
}

type LoginRequest struct {
	Account  string `json:"account" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterRequest struct {
	Account         string `json:"account" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirm_password" validate:"required"`
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func NewUserService(userRepo *repositories.UserRepository, config *config.Config) *UserService {
	return &UserService{
		userRepo: userRepo,
		config:   config,
	}
}

// 用户登录
func (s *UserService) Login(req LoginRequest) (*LoginResponse, error) {
	// 查找用户
	user, err := s.userRepo.GetByAccount(req.Account)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("用户不存在")
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.Password) {
		return nil, errors.New("密码错误")
	}

	// 生成JWT token
	token, err := utils.GenerateJWT(user.ID, s.config.JWT.Secret, s.config.JWT.ExpireHours)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: token,
		User:  *user,
	}, nil
}

// 用户注册
func (s *UserService) Register(req RegisterRequest) (*models.User, error) {
	// 验证密码确认
	if req.Password != req.ConfirmPassword {
		return nil, errors.New("两次密码输入不一致")
	}

	// 检查用户是否已存在
	existingUser, err := s.userRepo.GetByAccount(req.Account)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("该账号已被注册")
	}

	// 加密密码
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	// 创建用户
	user := &models.User{
		Account:  req.Account,
		Password: hashedPassword,
	}

	err = s.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// 根据ID获取用户信息
func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	return s.userRepo.GetByID(id)
}
