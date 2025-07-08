package router

import (
	"go-gin-project/api"
	"go-gin-project/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	authMiddleware := middleware.AuthMiddleware()
	guestMiddleware := middleware.GuestMiddleware()
	authController := api.InitializeAuthController()

	{
		r.POST("/register", guestMiddleware, authController.Register)
		r.POST("/login", guestMiddleware, authController.Login)

		r.POST("/logout", authMiddleware, authController.Logout)
		r.GET("/profile", authMiddleware, authController.Profile)
	}
}
