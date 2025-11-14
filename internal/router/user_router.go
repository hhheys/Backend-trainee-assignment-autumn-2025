package router

import (
	"AvitoPRService/internal/app"
	"AvitoPRService/internal/security"

	"github.com/gin-gonic/gin"
)

// NewUserRouter creates a new router for user endpoints.
func NewUserRouter(r *gin.Engine, app *app.App) {
	g := r.Group("/users")

	g.POST("/setIsActive", security.AdminAuthRequired(app.Config), app.Handlers.UserHandler.SetIsActive)

	// Временное решение для получения JWT токена пользователя.
	g.POST("/getAccessToken", app.Handlers.UserHandler.GetAccessToken)
}
