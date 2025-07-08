package middleware

import (
	"go-gin-project/helper"
	"go-gin-project/helper/responsejson"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, _ := ctx.Cookie("access_token")
		user, valid := helper.ValidateToken(accessToken, os.Getenv("GENERATE_ACCESS_TOKEN_SECRET"))

		if valid {
			ctx.Set("userId", user.Id)
			ctx.Next()
			return
		}

		responsejson.Unauthorized(ctx)
		ctx.Abort()
	}
}
