package app

import (
	"net/http"

	"github.com/foolin/goview"
	"github.com/gin-gonic/gin"
)

func (a *App) SetupRoutes(r *gin.Engine) {
	r.GET("/", a.handleIndex)
	r.GET("/page", a.handlePage)
	r.Static("/static", "./static")
}

func (a *App) handleIndex(c *gin.Context) {
	err := goview.Render(c.Writer, http.StatusOK, "index", goview.M{
		"title": "Hello World!",
	})
	if err != nil {
		panic(err)
	}
}

func (a *App) handlePage(c *gin.Context) {
	err := goview.Render(c.Writer, http.StatusOK, "page", goview.M{
		"title": "My page!",
	})
	if err != nil {
		panic(err)
	}
}
