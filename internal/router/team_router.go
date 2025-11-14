package router

import (
	"AvitoPRService/internal/app"

	"github.com/gin-gonic/gin"
)

func NewTeamRouter(r *gin.Engine, app *app.App) {
	g := r.Group("/team")

	g.POST("/add", app.Handlers.TeamHandler.AddTeam)
	g.GET("/get", app.Handlers.TeamHandler.FindByName)
}
