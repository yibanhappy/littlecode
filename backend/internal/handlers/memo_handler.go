package handlers

import (
	"strconv"

	"github.com/labstack/echo/v4"

	"littlecode-backend/internal/middleware"
	"littlecode-backend/internal/services"
	"littlecode-backend/pkg/utils"
)

type MemoHandler struct {
	memoService *services.MemoService
}

func NewMemoHandler(memoService *services.MemoService) *MemoHandler {
	return &MemoHandler{
		memoService: memoService,
	}
}

// 获取备忘录列表
func (h *MemoHandler) GetMemos(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	memos, err := h.memoService.GetMemosByUserID(userID)
	if err != nil {
		return utils.InternalError(c, "获取备忘录列表失败")
	}

	return utils.Success(c, memos)
}

// 获取单个备忘录
func (h *MemoHandler) GetMemo(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的备忘录ID")
	}

	memo, err := h.memoService.GetMemoByID(uint(id), userID)
	if err != nil {
		return utils.InternalError(c, "获取备忘录失败")
	}
	if memo == nil {
		return utils.NotFound(c, "备忘录不存在")
	}

	return utils.Success(c, memo)
}

// 创建备忘录
func (h *MemoHandler) CreateMemo(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	var req services.CreateMemoRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "参数格式错误")
	}

	memo, err := h.memoService.CreateMemo(req, userID)
	if err != nil {
		return utils.InternalError(c, "创建备忘录失败")
	}

	return utils.Success(c, memo)
}

// 更新备忘录
func (h *MemoHandler) UpdateMemo(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的备忘录ID")
	}

	var req services.UpdateMemoRequest
	if err := c.Bind(&req); err != nil {
		return utils.BadRequest(c, "参数格式错误")
	}

	memo, err := h.memoService.UpdateMemo(uint(id), req, userID)
	if err != nil {
		return utils.InternalError(c, err.Error())
	}

	return utils.Success(c, memo)
}

// 删除备忘录
func (h *MemoHandler) DeleteMemo(c echo.Context) error {
	userID, err := middleware.GetUserID(c)
	if err != nil {
		return utils.Unauthorized(c, "用户未认证")
	}

	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return utils.BadRequest(c, "无效的备忘录ID")
	}

	err = h.memoService.DeleteMemo(uint(id), userID)
	if err != nil {
		return utils.InternalError(c, "删除备忘录失败")
	}

	return utils.Success(c, map[string]string{"message": "删除成功"})
}
