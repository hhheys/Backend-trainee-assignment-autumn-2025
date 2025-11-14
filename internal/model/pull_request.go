// Package model contains core data structures used in the application,
// including definitions of users, teams, and pull requests.
package model

// PullRequest represents a pull request in the service.
type PullRequest struct {
	ID                uint   `json:"pull_request_id"`
	Name              string `json:"pull_request_name"`
	AuthorID          uint   `json:"author_id"`
	Status            string `json:"status"`
	AssignedReviewers []uint `json:"assigned_reviewers"`
}

// NewPullRequest creates a new PullRequest instance.
func NewPullRequest(id uint, name string, authorID uint, status string, assignedReviewers []uint) *PullRequest {
	return &PullRequest{
		ID:                id,
		Name:              name,
		AuthorID:          authorID,
		Status:            status,
		AssignedReviewers: assignedReviewers,
	}
}
