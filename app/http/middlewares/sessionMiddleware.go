package middlewares

import (
	"net/http"

	"github.com/amerta-teknologi/az-finance/app/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		defer func() {
			if err := recover(); err != nil {
				ctx.Redirect(http.StatusFound, "/auth")
			}
		}()

		session := sessions.Default(ctx)

		if key := session.Get("user.key"); key != nil {
			ctx.Set("user.key", key)

			ctx.Set("user.data", models.User{
				Id:          session.Get("user.id").(int),
				Email:       session.Get("user.email").(string),
				Name:        session.Get("user.name").(string),
				Slug:        session.Get("user.slug").(string),
				ImgUrl:      session.Get("user.img_url").(string),
				PhoneNumber: session.Get("user.phone_number").(string),
			})
		} else {
			ctx.Set("user.key", "")

			ctx.Set("user.data", models.User{
				Id:    0,
				Email: "",
				Name:  "",
				Slug:  "",
			})
		}

		ctx.Next()
	}
}
