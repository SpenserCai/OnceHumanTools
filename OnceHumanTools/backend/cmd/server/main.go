package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oncehuman/tools/api"
	"github.com/oncehuman/tools/config"
	"github.com/oncehuman/tools/internal/middleware"
	"github.com/sirupsen/logrus"
)

// @title OnceHuman Tools API
// @version 1.0
// @description OnceHuman游戏工具集API服务
// @host localhost:8080
// @BasePath /api/v1
func main() {
	// 加载环境变量
	if err := godotenv.Load(); err != nil {
		logrus.Warn("未找到.env文件，使用系统环境变量")
	}

	// 初始化配置
	cfg := config.Load()

	// 设置日志级别
	if cfg.Debug {
		logrus.SetLevel(logrus.DebugLevel)
		gin.SetMode(gin.DebugMode)
	} else {
		logrus.SetLevel(logrus.InfoLevel)
		gin.SetMode(gin.ReleaseMode)
	}

	// 创建Gin引擎
	r := gin.New()
	
	// 添加中间件
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	// 注册API路由
	api.RegisterRoutes(r, cfg)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logrus.Infof("服务器启动在端口 %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("服务器启动失败: %v", err)
	}
}