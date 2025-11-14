package router

import (
	"AvitoPRService/internal/app"

	"github.com/gin-gonic/gin"
)

// NewTeamRouter creates a new team router endpoints.
func NewTeamRouter(r *gin.Engine, app *app.App) {
	g := r.Group("/team")

	g.POST("/add", app.Handlers.TeamHandler.AddTeam)
	g.GET("/get", app.Handlers.TeamHandler.FindByName)
}
