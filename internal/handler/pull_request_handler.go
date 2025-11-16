package handler

import (
	"AvitoPRService/internal/model/dto"
	"AvitoPRService/internal/model/response"
	response2 "AvitoPRService/internal/model/response/error_response"
	"AvitoPRService/internal/repository/postgres"
	"errors"
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
	if err != nil {
		if errors.Is(err, postgres.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, response2.NewErrorResponse(response2.BadRequest, err.Error()))
			return
		} else if errors.Is(err, postgres.ErrUserNoTeamFound) {
			c.JSON(http.StatusNotFound, response2.NewErrorResponse(response2.BadRequest, err.Error()))
			return
		} else if errors.Is(err, postgres.ErrPullRequestAlreadyExists) {
			c.JSON(http.StatusConflict, response2.NewErrorResponse(response2.PullRequestAlreadyExists, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response2.NewErrorResponse(response2.InternalServerError, err.Error()))
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
	if err != nil {
		if errors.Is(err, postgres.ErrPullRequestNotExists) {
			c.JSON(http.StatusNotFound, response2.NewErrorResponse(response2.NotFound, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response2.NewErrorResponse(response2.InternalServerError, err.Error()))
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
	if err != nil {
		if errors.Is(err, postgres.ErrPullRequestMergedReassign) {
			c.JSON(http.StatusConflict, response2.NewErrorResponse(response2.PullRequestAlreadyMerged, err.Error()))
			return
		} else if errors.Is(err, postgres.ErrUserIsNotAssignedToPR) {
			c.JSON(http.StatusConflict, response2.NewErrorResponse(response2.PullRequestNotAssigned, err.Error()))
			return
		} else if errors.Is(err, postgres.ErrNoActiveReplacementCandidates) {
			c.JSON(http.StatusConflict, response2.NewErrorResponse(response2.PullRequestNoCandidate, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, response2.NewErrorResponse(response2.InternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.NewPullRequestResponse(pr))
}
