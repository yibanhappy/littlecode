package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"littlecode-backend/internal/config"
	"littlecode-backend/internal/handlers"
	authMiddleware "littlecode-backend/internal/middleware"
	"littlecode-backend/internal/repositories"
	"littlecode-backend/internal/services"
	"littlecode-backend/pkg/database"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// 初始化数据库
	err = database.InitDB(cfg)
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	// 获取数据库连接
	db := database.GetDB()

	// 初始化仓库层
	userRepo := repositories.NewUserRepository(db)
	memoRepo := repositories.NewMemoRepository(db)
	timerRepo := repositories.NewTimerRepository(db)

	// 初始化服务层
	userService := services.NewUserService(userRepo, cfg)
	memoService := services.NewMemoService(memoRepo)
	timerService := services.NewTimerService(timerRepo)

	// 初始化处理器
	userHandler := handlers.NewUserHandler(userService)
	memoHandler := handlers.NewMemoHandler(memoService)
	timerHandler := handlers.NewTimerHandler(timerService)

	// 创建Echo实例
	e := echo.New()

	// 添加中间件
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// 公开路由（不需要认证）
	public := e.Group("/api")
	public.POST("/login", userHandler.Login)
	public.POST("/register", userHandler.Register)

	// 需要认证的路由
	protected := e.Group("/api")
	protected.Use(authMiddleware.JWTAuth(cfg))

	// 用户相关路由
	protected.GET("/profile", userHandler.GetProfile)

	// 备忘录相关路由
	memoGroup := protected.Group("/memos")
	memoGroup.GET("", memoHandler.GetMemos)
	memoGroup.GET("/:id", memoHandler.GetMemo)
	memoGroup.POST("", memoHandler.CreateMemo)
	memoGroup.PUT("/:id", memoHandler.UpdateMemo)
	memoGroup.DELETE("/:id", memoHandler.DeleteMemo)

	// 计时器相关路由
	timerGroup := protected.Group("/timers")
	timerGroup.GET("", timerHandler.GetTimers)
	timerGroup.GET("/:id", timerHandler.GetTimer)
	timerGroup.POST("", timerHandler.CreateTimer)
	timerGroup.PUT("/:id", timerHandler.UpdateTimer)
	timerGroup.POST("/:id/start", timerHandler.StartTimer)
	timerGroup.POST("/:id/stop", timerHandler.StopTimer)
	timerGroup.DELETE("/:id", timerHandler.DeleteTimer)

	// 启动服务器
	serverAddr := cfg.Server.Host + ":" + cfg.Server.Port
	log.Printf("Server starting on http://%s", serverAddr)
	if err := e.Start(":" + cfg.Server.Port); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
