package router

import (
	"AvitoPRService/internal/app"

	"github.com/gin-gonic/gin"
)

func NewUserRouter(r *gin.Engine, app *app.App) {
	g := r.Group("/users")

	g.POST("/setIsActive")
}
