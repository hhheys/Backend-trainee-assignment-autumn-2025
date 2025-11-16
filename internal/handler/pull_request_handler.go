package handler

import (
	"AvitoPRService/internal/model/dto"
	"AvitoPRService/internal/model/response"
	response2 "AvitoPRService/internal/model/response/error_response"
	"AvitoPRService/internal/repository/postgres"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PullRequestHandler handles pull request-related operations.
type PullRequestHandler struct {
	r postgres.PullRequestRepository
}

// NewPullRequestHandler creates a new PullRequestHandler instance.
func NewPullRequestHandler(r postgres.PullRequestRepository) *PullRequestHandler {
	return &PullRequestHandler{r: r}
}

// CreatePullRequest handles the creation of a new pull request.
func (h *PullRequestHandler) CreatePullRequest(c *gin.Context) {
	var pullRequestCreateDto dto.PullRequestCreateDto
	if err := c.ShouldBindJSON(&pullRequestCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, response2.NewErrorResponse(response2.BadRequest, err.Error()))
	}
	pullRequest, err := h.r.CreatePullRequest(
		pullRequestCreateDto.PullRequestID,
		pullRequestCreateDto.PullRequestName,
		pullRequestCreateDto.AuthorID,
	)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	c.JSON(http.StatusCreated, response.NewPullRequestResponse(pullRequest))
}

// MergePullRequest handles the merging of a pull request.
func (h *PullRequestHandler) MergePullRequest(c *gin.Context) {
	var pullRequestMergeDto dto.PullRequestMergeDto
	if err := c.ShouldBindJSON(&pullRequestMergeDto); err != nil {
		c.JSON(http.StatusBadRequest, response2.NewErrorResponse(response2.BadRequest, err.Error()))
	}
	pr, err := h.r.MergePullRequest(pullRequestMergeDto.PullRequestID)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	c.JSON(http.StatusCreated, response.NewPullRequestResponse(pr))
}

// ReassignReviewer handles the reassignment of a reviewer to a pull request.
func (h *PullRequestHandler) ReassignReviewer(c *gin.Context) {
	var prReassignDto dto.PullRequestReassignDto
	if err := c.ShouldBindJSON(&prReassignDto); err != nil {
		c.JSON(http.StatusBadRequest, response2.NewErrorResponse(response2.BadRequest, err.Error()))
	}
	pr, err := h.r.ReassignReviewer(prReassignDto.PullRequestID, prReassignDto.OldReviewerID)
	if status, errResp := response2.HandleError(err); errResp != nil {
		c.JSON(status, errResp)
		return
	}

	c.JSON(http.StatusCreated, response.NewPullRequestResponse(pr))
}
