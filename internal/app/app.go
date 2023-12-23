package app

import (
	"github.com/gin-gonic/gin"
)

type App struct {
}

func New() *App {
	return &App{}
}

func (a *App) Run(args []string) error {
	r := gin.Default()
	a.SetupRoutes(r)

	return r.Run(":3000")
}
