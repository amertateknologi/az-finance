package middlewares

import (
	"github.com/gin-gonic/gin"
)

func WebviewMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("webview", 0)

		if ctx.Query("wv") == "1" {
			ctx.Set("webview", 1)
		}

		ctx.Next()
	}
}
