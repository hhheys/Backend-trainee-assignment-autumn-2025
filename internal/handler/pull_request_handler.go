package handler

import (
	"AvitoPRService/internal/db"
	"AvitoPRService/internal/dto"
	"AvitoPRService/internal/repository"
	"AvitoPRService/internal/response"
	errorResponse "AvitoPRService/internal/response/error_response"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PullRequestHandler handles pull request-related operations.
type PullRequestHandler struct {
	r repository.PullRequestRepository
}

// NewPullRequestHandler creates a new PullRequestHandler instance.
func NewPullRequestHandler(r repository.PullRequestRepository) *PullRequestHandler {
	return &PullRequestHandler{r: r}
}

// CreatePullRequest handles the creation of a new pull request.
func (h *PullRequestHandler) CreatePullRequest(c *gin.Context) {
	var pullRequestCreateDto dto.PullRequestCreateDto
	if err := c.ShouldBindJSON(&pullRequestCreateDto); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BadRequest, err.Error()))
	}
	pullRequest, err := h.r.CreatePullRequest(
		pullRequestCreateDto.PullRequestID,
		pullRequestCreateDto.PullRequestName,
		pullRequestCreateDto.AuthorID,
	)
	if err != nil {
		if errors.Is(err, db.ErrUserNotFound) {
			c.JSON(http.StatusNotFound, errorResponse.NewErrorResponse(errorResponse.BadRequest, err.Error()))
			return
		} else if errors.Is(err, db.ErrUserNoTeamFound) {
			c.JSON(http.StatusNotFound, errorResponse.NewErrorResponse(errorResponse.BadRequest, err.Error()))
			return
		} else if errors.Is(err, db.ErrPullRequestAlreadyExists) {
			c.JSON(http.StatusConflict, errorResponse.NewErrorResponse(errorResponse.PullRequestAlreadyExists, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.InternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.NewPullRequestResponse(pullRequest))
}

// MergePullRequest handles the merging of a pull request.
func (h *PullRequestHandler) MergePullRequest(c *gin.Context) {
	var pullRequestMergeDto dto.PullRequestMergeDto
	if err := c.ShouldBindJSON(&pullRequestMergeDto); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BadRequest, err.Error()))
	}
	pr, err := h.r.MergePullRequest(pullRequestMergeDto.PullRequestID)
	if err != nil {
		if errors.Is(err, db.ErrPullRequestNotExists) {
			c.JSON(http.StatusNotFound, errorResponse.NewErrorResponse(errorResponse.NotFound, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.InternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.NewPullRequestResponse(pr))
}

// ReassignReviewer handles the reassignment of a reviewer to a pull request.
func (h *PullRequestHandler) ReassignReviewer(c *gin.Context) {
	var prReassignDto dto.PullRequestReassignDto
	if err := c.ShouldBindJSON(&prReassignDto); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse.NewErrorResponse(errorResponse.BadRequest, err.Error()))
	}
	pr, err := h.r.ReassignReviewer(prReassignDto.PullRequestID, prReassignDto.OldReviewerID)
	if err != nil {
		if errors.Is(err, db.ErrPullRequestMergedReassign) {
			c.JSON(http.StatusConflict, errorResponse.NewErrorResponse(errorResponse.PullRequestAlreadyMerged, err.Error()))
			return
		} else if errors.Is(err, db.ErrUserIsNotAssignedToPR) {
			c.JSON(http.StatusConflict, errorResponse.NewErrorResponse(errorResponse.PullRequestNotAssigned, err.Error()))
			return
		} else if errors.Is(err, db.ErrNoActiveReplacementCandidates) {
			c.JSON(http.StatusConflict, errorResponse.NewErrorResponse(errorResponse.PullRequestNoCandidate, err.Error()))
			return
		}
		c.JSON(http.StatusInternalServerError, errorResponse.NewErrorResponse(errorResponse.InternalServerError, err.Error()))
		return
	}
	c.JSON(http.StatusCreated, response.NewPullRequestResponse(pr))
}
