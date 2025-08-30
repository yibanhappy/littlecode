package utils

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// 统一响应结构
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// 成功响应
func Success(c echo.Context, data interface{}) error {
	return c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    data,
	})
}

// 错误响应
func Error(c echo.Context, code int, message string) error {
	return c.JSON(code, Response{
		Code:    code,
		Message: message,
	})
}

// 400 Bad Request
func BadRequest(c echo.Context, message string) error {
	return Error(c, http.StatusBadRequest, message)
}

// 401 Unauthorized
func Unauthorized(c echo.Context, message string) error {
	return Error(c, http.StatusUnauthorized, message)
}

// 404 Not Found
func NotFound(c echo.Context, message string) error {
	return Error(c, http.StatusNotFound, message)
}

// 500 Internal Server Error
func InternalError(c echo.Context, message string) error {
	return Error(c, http.StatusInternalServerError, message)
}
