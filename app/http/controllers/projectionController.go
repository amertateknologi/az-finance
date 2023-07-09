package controllers

import (
	"github.com/amerta-teknologi/az-finance/app/models"

	"github.com/gin-gonic/gin"
)

type ProjectionController struct{}

func (ctrl *ProjectionController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		model := models.Projection{}
		resp := model.FindAll(ctx)
		ctx.JSON(200, resp)
	}
}
