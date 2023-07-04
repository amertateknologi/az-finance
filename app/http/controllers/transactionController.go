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

// func (ctrl *TransactionController) Post() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		model := models.Transaction{}
// 		resp := model.Post(c)
// 		c.JSON(200, resp)
// 	}
// }

// func (ctrl *TransactionController) Put() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		model := models.Transaction{}
// 		resp := model.Put(c)
// 		c.JSON(200, resp)
// 	}
// }

// func (ctrl *TransactionController) Del() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		model := models.Transaction{}
// 		resp := model.Del(c)
// 		c.JSON(200, resp)
// 	}
// }
