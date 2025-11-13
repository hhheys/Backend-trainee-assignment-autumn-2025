package model

type PullRequest struct {
	ID                uint   `json:"pull_request_id"`
	Name              string `json:"pull_request_name"`
	AuthorID          uint   `json:"author_id"`
	Status            string `json:"status"`
	AssignedReviewers []uint `json:"assigned_reviewers"`
}

func NewPullRequest(id uint, name string, authorID uint, status string, assignedReviewers []uint) *PullRequest {
	return &PullRequest{
		ID:                id,
		Name:              name,
		AuthorID:          authorID,
		Status:            status,
		AssignedReviewers: assignedReviewers,
	}
}
