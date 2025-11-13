package router

import (
	"AvitoPRService/internal/app"

	"github.com/gin-gonic/gin"
)

func NewTeamRouter(r *gin.Engine, app *app.App) {
	g := r.Group("/team")

	g.POST("/add")
	g.GET("/get")

}
