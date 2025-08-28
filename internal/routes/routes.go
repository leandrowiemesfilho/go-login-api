package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/leandrowiemesfilho/login-api/internal/handlers"
)

func Setup(router *gin.Engine) {
	v1 := router.Group("/api/v1/auth")

	v1.POST("/login", handlers.Login)
	v1.POST("/signup", handlers.Signup)
}
