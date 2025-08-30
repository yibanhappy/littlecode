package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"littlecode-backend/internal/middleware"
	"littlecode-backend/internal/services"
	"littlecode-backend/pkg/utils"
)

type TimerHandler struct {
	timerService *services.TimerService
}

func NewTimerHandler(timerService *services.TimerService) *TimerHandler {
	return &TimerHandler{
		timerService: timerService,
	}
}

// 获取计时器列表
func (h *TimerHandler) GetTimers(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	timers, err := h.timerService.GetTimersByUserID(userID)
	if err != nil {
		return utils.InternalError(c, "获取计时器列表失败")
	}

	return utils.Success(c, timers)
}

// 获取单个计时器
func (h *TimerHandler) GetTimer(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的计时器ID")
	}

	timer, err := h.timerService.GetTimerByID(uint(id), userID)
	if err != nil {
		return utils.InternalError(c, "获取计时器失败")
	}
	if timer == nil {
		return utils.NotFound(c, "计时器不存在")
	}

	return utils.Success(c, timer)
}

// 创建计时器
func (h *TimerHandler) CreateTimer(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	var req services.CreateTimerRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "参数格式错误")
	}

	timer, err := h.timerService.CreateTimer(req, userID)
	if err != nil {
		return utils.InternalError(c, "创建计时器失败")
	}

	return utils.Success(c, timer)
}

// 更新计时器
func (h *TimerHandler) UpdateTimer(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的计时器ID")
	}

	var req services.UpdateTimerRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "参数格式错误")
	}

	timer, err := h.timerService.UpdateTimer(uint(id), req, userID)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}

	return utils.Success(c, timer)
}

// 启动计时器
func (h *TimerHandler) StartTimer(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的计时器ID")
	}

	timer, err := h.timerService.StartTimer(uint(id), userID)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}

	return utils.Success(c, timer)
}

// 停止计时器
func (h *TimerHandler) StopTimer(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的计时器ID")
	}

	timer, err := h.timerService.StopTimer(uint(id), userID)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}

	return utils.Success(c, timer)
}

// 删除计时器
func (h *TimerHandler) DeleteTimer(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的计时器ID")
	}

	err = h.timerService.DeleteTimer(uint(id), userID)
	if err != nil {
		return utils.InternalError(c, "删除计时器失败")
	}

	return utils.Success(c, map[string]string{"message": "删除成功"})
}
