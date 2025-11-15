package router

import (
	"AvitoPRService/internal/app"

	"github.com/gin-gonic/gin"
)

// NewPullRequestRouter creates a new pull request router
func NewPullRequestRouter(r *gin.Engine, app *app.App) {
	g := r.Group("/pullRequest")

	g.POST("/create", app.Handlers.PullRequestHandler.CreatePullRequest)
	g.POST("/merge", app.Handlers.PullRequestHandler.MergePullRequest)
	g.POST("/reassign", app.Handlers.PullRequestHandler.ReassignReviewer)
}
