package response

const (
	// BadRequest is the error code for bad requests
	BadRequest = "BAD_REQUEST"
	// NotFound is the error code for not found requests
	NotFound = "NOT_FOUND"
	// TeamExists is the error code for team already exists
	TeamExists = "TEAM_EXISTS"
	// InternalServerError is the error code for internal server errors
	InternalServerError = "INTERNAL_SERVER_ERROR"
	// PullRequestAlreadyExists is the error code for pull request already exists
	PullRequestAlreadyExists = "PR_EXISTS"
	// PullRequestAlreadyMerged is the error code for pull request already merged
	PullRequestAlreadyMerged = "PR_MERGED"
	// PullRequestNotAssigned is the error code for pull request not assigned
	PullRequestNotAssigned = "NOT_ASSIGNED"
	// PullRequestNoCandidate is the error code for pull request no candidate
	PullRequestNoCandidate = "NO_CANDIDATE"
)
