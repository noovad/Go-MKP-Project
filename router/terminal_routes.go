package router

import (
	"go-gin-project/api"

	"github.com/gin-gonic/gin"
)

func TerminalRouter(router *gin.Engine) {
	controller := api.InitializeTerminalController()

	terminalRouter := router.Group("/terminal")
	{
		terminalRouter.GET("", controller.FindAll)
		terminalRouter.GET("/:terminalId", controller.FindById)
		terminalRouter.POST("", controller.Create)
		terminalRouter.PUT("/:terminalId", controller.Update)
		terminalRouter.DELETE("/:terminalId", controller.Delete)
	}
}
