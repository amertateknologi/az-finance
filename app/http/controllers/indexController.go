package controllers

import (
	"net/http"

	"github.com/amerta-teknologi/az-finance/app/models"
	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (ctrl *IndexController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData := ctx.MustGet("user.data")

		if userData.(models.User).Id == 0 {
			ctx.Redirect(http.StatusFound, "/auth")
		}

		ctx.HTML(http.StatusOK, "pages/index", gin.H{
			"webview":  ctx.MustGet("webview"),
			"userData": ctx.MustGet("user.data"),
		})
	}
}
