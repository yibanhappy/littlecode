package handlers

import (
	"github.com/labstack/echo/v4"

	"littlecode-backend/internal/middleware"
	"littlecode-backend/internal/services"
	"littlecode-backend/pkg/utils"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// 用户登录
func (h *UserHandler) Login(c echo.Context) error {
	var req services.LoginRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "参数格式错误")
	}

	response, err := h.userService.Login(req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, response)
}

// 用户注册
func (h *UserHandler) Register(c echo.Context) error {
	var req services.RegisterRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "参数格式错误")
	}

	user, err := h.userService.Register(req)
	if err != nil {
		return utils.BadRequest(c, err.Error())
	}

	return utils.Success(c, user)
}

// 获取当前用户信息
func (h *UserHandler) GetProfile(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	user, err := h.userService.GetUserByID(userID)
	if err != nil {
		return utils.InternalError(c, "获取用户信息失败")
	}

	return utils.Success(c, user)
}
