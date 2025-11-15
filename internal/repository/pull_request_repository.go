package repository

import "AvitoPRService/internal/model"

// PullRequestRepository is an interface for working with the database for pull requests
type PullRequestRepository interface {
	CreatePullRequest(pullRequestID string, pullRequestName string, authorID string) (*model.PullRequest, error)
	MergePullRequest(pullRequestID string) (*model.PullRequest, error)
	FindPullRequestByID(pullRequestID string) (*model.PullRequest, error)
	ReassignReviewer(pullRequestID string, reviewerID string) (*model.PullRequest, error)
}
