package app

import (
	"fmt"
	"net/http"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

func (a *App) SetupRoutes(r *gin.Engine) {
	r.GET("/", func(c *gin.Context) {
		gv := goview.New(goview.Config{
			Root:      "views/marketing",
			Extension: ".html",
			Master:    "layouts/master",
		})
		goview.Use(gv)
		err := goview.Render(c.Writer, http.StatusOK, "index", goview.M{
			"title": "Hello World!",
		})
		if err != nil {
			panic(err)
		}
	})
	r.GET("/login-modal", func(c *gin.Context) {
		gv := goview.New(goview.Config{
			Root:      "views/_shared/partials",
			Extension: ".html",
		})
		goview.Use(gv)
		err := goview.Render(c.Writer, http.StatusOK, "login-modal.html", goview.M{})
		if err != nil {
			panic(err)
		}
	})
	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		fmt.Println("username: ", username)
		fmt.Println("password: ", password)
		if username == "admin" && password == "admin" {
			c.Redirect(http.StatusMovedPermanently, "/account")
			return
		}
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "invalid credentials",
		})
	})
	r.GET("/account", func(c *gin.Context) {
		gv := goview.New(goview.Config{
			Root:      "views/app",
			Extension: ".html",
			Master:    "layouts/master",
		})
		goview.Use(gv)
		err := goview.Render(c.Writer, http.StatusOK, "user/account", goview.M{
			"title": "My account!",
		})
		if err != nil {
			panic(err)
		}
	})
	r.Static("/static", "./static")
}
