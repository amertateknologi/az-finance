package controllers

import (
	"github.com/amerta-teknologi/az-finance/app/models"

	"github.com/gin-gonic/gin"
)

type TransactionController struct{}

func (ctrl *TransactionController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		model := models.Transaction{}
		resp := model.FindAll(ctx)
		ctx.JSON(200, resp)
	}
}
