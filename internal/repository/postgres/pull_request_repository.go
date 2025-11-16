package postgres

import (
	"AvitoPRService/internal/model/db"
)

// PullRequestRepository is an interface for working with the database for pull requests
type PullRequestRepository interface {
	CreatePullRequest(pullRequestID string, pullRequestName string, authorID string) (*db.PullRequest, error)
	MergePullRequest(pullRequestID string) (*db.PullRequest, error)
	FindPullRequestByID(pullRequestID string) (*db.PullRequest, error)
	ReassignReviewer(pullRequestID string, reviewerID string) (*db.PullRequest, error)
}
