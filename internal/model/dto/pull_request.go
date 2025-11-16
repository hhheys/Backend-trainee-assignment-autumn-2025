package dto

// PullRequestCreateDto is a DTO for creating a pull request
type PullRequestCreateDto struct {
	PullRequestID   string `json:"pull_request_id" binding:"required"`
	PullRequestName string `json:"pull_request_name" binding:"required"`
	AuthorID        string `json:"author_id" binding:"required"`
}

// PullRequestMergeDto is a DTO for merging a pull request
type PullRequestMergeDto struct {
	PullRequestID string `json:"pull_request_id" binding:"required"`
}

// PullRequestReassignDto is a DTO for reassigning a pull request
type PullRequestReassignDto struct {
	PullRequestID string `json:"pull_request_id" binding:"required"`
	OldReviewerID string `json:"old_reviewer_id" binding:"required"`
}
