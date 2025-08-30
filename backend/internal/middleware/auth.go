package middleware

import (
	"strings"

	"github.com/labstack/echo/v4"

	"littlecode-backend/internal/config"
	"littlecode-backend/pkg/utils"
)

func JWTAuth(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 获取Authorization header
			authHeader := c.Request().Header.Get("Authorization")
			if authHeader == "" {
				return utils.Unauthorized(c, "Missing authorization header")
			}

			// 检查Bearer前缀
			tokenParts := strings.Split(authHeader, " ")
			if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
				return utils.Unauthorized(c, "Invalid authorization header format")
			}

			token := tokenParts[1]

			// 验证token
			claims, err := utils.ValidateJWT(token, cfg.JWT.Secret)
			if err != nil {
				return utils.Unauthorized(c, "Invalid token")
			}

			// 将用户ID存储到context中
			if userID, ok := (*claims)["user_id"].(float64); ok {
				c.Set("user_id", uint(userID))
			} else {
				return utils.Unauthorized(c, "Invalid token claims")
			}

			return next(c)
		}
	}
}

// 获取当前用户ID的辅助函数
func GetUserID(c echo.Context) (uint, error) {
	userID := c.Get("user_id")
	if userID == nil {
		return 0, echo.NewHTTPError(401, "User not authenticated")
	}

	if id, ok := userID.(uint); ok {
		return id, nil
	}

	return 0, echo.NewHTTPError(401, "Invalid user ID")
}
