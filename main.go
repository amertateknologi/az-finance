package main

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/amerta-teknologi/az-finance/app/http/controllers"
	"github.com/amerta-teknologi/az-finance/app/http/middlewares"
	"github.com/amerta-teknologi/az-finance/app/utils"
	"github.com/amerta-teknologi/az-finance/config"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
)

func main() {
	configs, err := config.LoadConfig(".", &config.Data)

	if err != nil {
		fmt.Printf("cannot load config: %w", err)
		return
	}

	var group gin.IRoutes
	r := gin.Default()

	store := cookie.NewStore([]byte("secret-key")) // Replace "secret-key" with your own secret key
	r.Use(sessions.Sessions("session-name", store))

	html := template.Must(template.New("").Delims("{{", "}}").Funcs(template.FuncMap{
		"nl2br":   utils.Nl2br,
		"in":      utils.In,
		"add":     utils.Add,
		"sub":     utils.Sub,
		"explode": utils.ExplodeFunc,
	}).ParseFiles(utils.Filewalk()...))
	debugPrintLoadTemplate(html)

	r.SetHTMLTemplate(html)
	r.Static("/assets", "./public/assets")
	r.Static("/images", "./public/images")
	r.Static("/icons", "./public/icons")

	r.Use(middlewares.WebviewMiddleware(), middlewares.SessionMiddleware())

	ctrl := controllers.IndexController{}
	r.GET("", ctrl.Get())

	group = r.Group("/auth")
	{
		ctrl := controllers.AuthController{}
		group.POST("/sign-in", ctrl.SignIn())
		group.GET("/sign-out", ctrl.SignOut())
		group.GET("", ctrl.Get())
	}

	group = r.Group("/transactions")
	{
		ctrl := controllers.TransactionController{}
		group.GET("", ctrl.Get())
	}

	//r.Run("192.168.18.174:8080")
	//r.Run("127.0.0.1:8081")

	r.Run(configs.HTTPServerAddress)
}

func debugPrintLoadTemplate(tmpl *template.Template) {
	var buf strings.Builder
	for _, tmpl := range tmpl.Templates() {
		buf.WriteString("\t- ")
		buf.WriteString(tmpl.Name())
		buf.WriteString("\n")
	}
	utils.DebugPrint("Loaded HTML Templates (%d): \n%s\n", len(tmpl.Templates()), buf.String())
}
