package routes

import (
	"github.com/Dcyaprogrammer/goe.git/handlers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SetupRoutes 设置路由
func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	// 创建处理器
	userHandler := handlers.NewUserHandler(db)

	// API v1 路由组
	v1 := r.Group("/api/v1")
	{
		// 健康检查
		v1.GET("/health", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"status":  "ok",
				"message": "Server is running",
			})
		})

		// 用户路由
		users := v1.Group("/users")
		{
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	// 根路径
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to GoE API",
			"version": "1.0.0",
		})
	})
}
