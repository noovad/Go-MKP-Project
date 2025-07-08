package middleware

import (
	"go-gin-project/helper"
	"go-gin-project/helper/responsejson"
	"os"

	"github.com/gin-gonic/gin"
)

func GuestMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, _ := ctx.Cookie("access_token")
		_, valid := helper.ValidateToken(accessToken, os.Getenv("GENERATE_ACCESS_TOKEN_SECRET"))

		if valid {
			responsejson.BadRequest(ctx, nil, "Already authenticated")
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
