package router

import (
	"go-gin-project/api"
	"go-gin-project/middleware"

	"github.com/gin-gonic/gin"
)

func TerminalRouter(router *gin.Engine) {
	controller := api.InitializeTerminalController()
	authMiddleware := middleware.AuthMiddleware()
	terminalRouter := router.Group("/terminal")
	{
		terminalRouter.GET("", authMiddleware, controller.FindAll)
		terminalRouter.GET("/:terminalId", authMiddleware, controller.FindById)
		terminalRouter.POST("", authMiddleware, controller.Create)
		terminalRouter.PUT("/:terminalId", authMiddleware, controller.Update)
		terminalRouter.DELETE("/:terminalId", authMiddleware, controller.Delete)
	}
}
