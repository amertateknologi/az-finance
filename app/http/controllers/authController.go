package controllers

import (
	"fmt"
	"net/http"

	"github.com/amerta-teknologi/az-finance/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (ctrl *AuthController) Get() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.HTML(http.StatusOK, "pages/auth", gin.H{})
	}
}

func (ctrl *AuthController) SignIn() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		model := models.Auth{}
		l := model.Login(ctx)

		fmt.Println(l.Data.AccessToken)

		if l.Data.AccessToken != "" {
			session := sessions.Default(ctx)
			session.Set("user.key", l.Data.AccessToken)
			session.Set("user.email", l.Data.User.Email)
			session.Set("user.id", l.Data.User.Id)
			session.Set("user.name", l.Data.User.Name)

			if err := session.Save(); err != nil {
				fmt.Println(err.Error())
			}

			ctx.Redirect(http.StatusFound, "/")
		} else {
			ctx.Redirect(http.StatusFound, ctx.Request.Referer())
		}
	}
}

func (ctrl *AuthController) SignOut() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()

		c.Redirect(http.StatusFound, "/")
	}
}
