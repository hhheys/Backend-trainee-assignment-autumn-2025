package response

import (
	"AvitoPRService/internal/model/db"
	"time"
)

// PullRequestWrapperResponse represents a wrapper for a pull request
type PullRequestWrapperResponse struct {
	PullRequest *PullRequestResponse `json:"pr"`
}

// PullRequestResponse represents a pull request
type PullRequestResponse struct {
	ID                string     `json:"pull_request_id"`
	Name              string     `json:"pull_request_name"`
	AuthorID          string     `json:"author_id"`
	Status            string     `json:"status"`
	AssignedReviewers []string   `json:"assigned_reviewers,omitempty"`
	MergedAt          *time.Time `json:"merged_at,omitempty"`
}

// NewPullRequestResponse creates a new PullRequestResponse
func NewPullRequestResponse(pullRequest *db.PullRequest) *PullRequestWrapperResponse {
	return &PullRequestWrapperResponse{
		PullRequest: &PullRequestResponse{
			ID:                pullRequest.ID,
			Name:              pullRequest.Name,
			AuthorID:          pullRequest.AuthorID,
			Status:            pullRequest.Status,
			AssignedReviewers: pullRequest.AssignedReviewers,
			MergedAt:          pullRequest.MergedAt,
		},
	}
}

// NewPullRequestResponses creates a new slice of PullRequestResponse
func NewPullRequestResponses(pullRequests []*db.PullRequest) []*PullRequestResponse {
	var responses []*PullRequestResponse
	if len(pullRequests) == 0 {
		return []*PullRequestResponse{}
	}

	for _, pr := range pullRequests {
		responses = append(responses, &PullRequestResponse{
			ID:                pr.ID,
			Name:              pr.Name,
			AuthorID:          pr.AuthorID,
			Status:            pr.Status,
			AssignedReviewers: pr.AssignedReviewers,
			MergedAt:          pr.MergedAt,
		})
	}

	return responses
}
