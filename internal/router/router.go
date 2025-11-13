package router

import (
	"AvitoPRService/internal/app"

	"github.com/gin-gonic/gin"
)

func NewRouter(app *app.App) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	NewUserRouter(r, app)
	NewTeamRouter(r, app)

	return r
}
