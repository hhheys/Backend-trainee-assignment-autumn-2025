package router

import (
	"AvitoPRService/internal/app"
	"AvitoPRService/internal/security"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, app *app.App) {
	g := r.Group("/users")

	g.POST("/setIsActive", security.AdminAuthReqired(app), app.Handlers.UserHandler.SetIsActive)
}
