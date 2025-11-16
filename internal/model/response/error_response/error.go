// Package response provides types for structured API responses,
// including error handling structures used to standardize error messages
// returned by the server.
package response

import (
	"AvitoPRService/internal/config/logger"
	"AvitoPRService/internal/repository/postgres"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
)

var errorMapping = map[error]struct {
	status  int
	code    string
	message string
}{
	postgres.ErrUserNotFound:                  {http.StatusNotFound, NotFound, postgres.ErrUserNotFound.Error()},
	postgres.ErrUserNoTeamFound:               {http.StatusNotFound, NotFound, postgres.ErrUserNoTeamFound.Error()},
	postgres.ErrUserIsNotActive:               {http.StatusBadRequest, BadRequest, postgres.ErrUserIsNotActive.Error()},
	postgres.ErrTeamExists:                    {http.StatusBadRequest, TeamExists, postgres.ErrTeamExists.Error()},
	postgres.ErrTeamNotExists:                 {http.StatusNotFound, NotFound, postgres.ErrTeamNotExists.Error()},
	postgres.ErrPullRequestAlreadyExists:      {http.StatusConflict, PullRequestAlreadyExists, postgres.ErrPullRequestAlreadyExists.Error()},
	postgres.ErrPullRequestNotExists:          {http.StatusNotFound, BadRequest, postgres.ErrPullRequestNotExists.Error()},
	postgres.ErrPullRequestMergedReassign:     {http.StatusConflict, PullRequestAlreadyMerged, postgres.ErrPullRequestMergedReassign.Error()},
	postgres.ErrUserIsNotAssignedToPR:         {http.StatusConflict, PullRequestNotAssigned, postgres.ErrUserIsNotAssignedToPR.Error()},
	postgres.ErrNoActiveReplacementCandidates: {http.StatusConflict, PullRequestNoCandidate, postgres.ErrNoActiveReplacementCandidates.Error()},
	postgres.ErrInternalServerError:           {http.StatusInternalServerError, InternalServerError, postgres.ErrInternalServerError.Error()},
}

// ErrorResponse represents the response structure for error responses.
type ErrorResponse struct {
	Error ErrorDetail `json:"error"`
}

// ErrorDetail represents the details of an error.
type ErrorDetail struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// NewErrorResponse creates a new ErrorResponse instance.
func NewErrorResponse(code, message string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorDetail{
			Code:    code,
			Message: message,
		},
	}
}

// HandleError handles errors and returns an appropriate HTTP status code and ErrorResponse.
func HandleError(err error) (int, *ErrorResponse) {
	if err == nil {
		return http.StatusOK, nil
	}

	// Проверяем ошибки валидации Gin/validator
	var ve validator.ValidationErrors
	if errors.As(err, &ve) {
		// Преобразуем ValidationErrors в строку или JSON
		errMessages := ""
		for _, fe := range ve {
			errMessages += fmt.Sprintf("Field '%s' failed on the '%s' tag; ", fe.Field(), fe.Tag())
		}
		return http.StatusBadRequest, NewErrorResponse(BadRequest, strings.TrimSpace(errMessages))
	}

	// Смотрим на кастомные маппинги ошибок
	if mapped, ok := errorMapping[err]; ok {
		return mapped.status, NewErrorResponse(mapped.code, mapped.message)
	}
	logger.Logger.Fatal(err.Error())
	// Неизвестная ошибка
	return http.StatusInternalServerError, NewErrorResponse(InternalServerError, postgres.ErrInternalServerError.Error())
}
