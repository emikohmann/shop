package apierrors

import (
	"fmt"
	"net/http"
)

type APIError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	status  int
	message string
}

func (apiErr apiError) Status() int {
	return apiErr.status
}

func (apiErr apiError) Message() string {
	return apiErr.message
}

func (apiErr apiError) Error() string {
	return fmt.Sprintf("Error %d: %s", apiErr.status, apiErr.message)
}

// NewBadRequestError creates a new API Error with status 400
func NewBadRequestError(message string) apiError {
	return apiError{
		status:  http.StatusBadRequest,
		message: message,
	}
}

// NewNotFoundError creates a new API Error with status 404
func NewNotFoundError(message string) apiError {
	return apiError{
		status:  http.StatusNotFound,
		message: message,
	}
}

// NewInternalServerError creates a new API Error with status 500
func NewInternalServerError(message string) apiError {
	return apiError{
		status:  http.StatusInternalServerError,
		message: message,
	}
}
