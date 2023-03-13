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

// NewUnauthorizedError creates a new API Error with status 401
func NewUnauthorizedError(message string) apiError {
    return apiError{
        status:  http.StatusUnauthorized,
        message: message,
    }
}

// NewPaymentRequiredError creates a new API Error with status 402
func NewPaymentRequiredError(message string) apiError {
    return apiError{
        status:  http.StatusPaymentRequired,
        message: message,
    }
}

// NewForbiddenError creates a new API Error with status 403
func NewForbiddenError(message string) apiError {
    return apiError{
        status:  http.StatusForbidden,
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

// NewMethodNotAllowedError creates a new API Error with status 405
func NewMethodNotAllowedError(message string) apiError {
    return apiError{
        status:  http.StatusMethodNotAllowed,
        message: message,
    }
}

// NewNotAcceptableError creates a new API Error with status 406
func NewNotAcceptableError(message string) apiError {
    return apiError{
        status:  http.StatusNotAcceptable,
        message: message,
    }
}

// NewProxyAuthenticationRequiredError creates a new API Error with status 407
func NewProxyAuthenticationRequiredError(message string) apiError {
    return apiError{
        status:  http.StatusProxyAuthRequired,
        message: message,
    }
}

// NewRequestTimeoutError creates a new API Error with status 408
func NewRequestTimeoutError(message string) apiError {
    return apiError{
        status:  http.StatusRequestTimeout,
        message: message,
    }
}

// NewConflictError creates a new API Error with status 409
func NewConflictError(message string) apiError {
    return apiError{
        status:  http.StatusConflict,
        message: message,
    }
}

// NewGoneError creates a new API Error with status 410
func NewGoneError(message string) apiError {
    return apiError{
        status:  http.StatusGone,
        message: message,
    }
}

// NewLengthRequiredError creates a new API Error with status 411
func NewLengthRequiredError(message string) apiError {
    return apiError{
        status:  http.StatusLengthRequired,
        message: message,
    }
}

// NewPreconditionFailedError creates a new API Error with status 412
func NewPreconditionFailedError(message string) apiError {
    return apiError{
        status:  http.StatusPreconditionFailed,
        message: message,
    }
}

// NewPayloadTooLargeError creates a new API Error with status 413
func NewPayloadTooLargeError(message string) apiError {
    return apiError{
        status:  http.StatusRequestEntityTooLarge,
        message: message,
    }
}

// NewURITooLongError creates a new API Error with status 414
func NewURITooLongError(message string) apiError {
    return apiError{
        status:  http.StatusRequestURITooLong,
        message: message,
    }
}

// NewUnsupportedMediaTypeError creates a new API Error with status 415
func NewUnsupportedMediaTypeError(message string) apiError {
    return apiError{
        status:  http.StatusUnsupportedMediaType,
        message: message,
    }
}

// NewRequestedRangeNotSatisfiableError creates a new API Error with status 416
func NewRequestedRangeNotSatisfiableError(message string) apiError {
    return apiError{
        status:  http.StatusRequestedRangeNotSatisfiable,
        message: message,
    }
}

// NewExpectationFailedError creates a new API Error with status 417
func NewExpectationFailedError(message string) apiError {
    return apiError{
        status:  http.StatusExpectationFailed,
        message: message,
    }
}

// NewImATeapotError creates a new API Error with status 418
func NewImATeapotError(message string) apiError {
    return apiError{
        status:  http.StatusTeapot,
        message: message,
    }
}

// NewMisdirectedRequestError creates a new API Error with status 421
func NewMisdirectedRequestError(message string) apiError {
    return apiError{
        status:  http.StatusMisdirectedRequest,
        message: message,
    }
}

// NewUnprocessableEntityError creates a new API Error with status 422
func NewUnprocessableEntityError(message string) apiError {
    return apiError{
        status:  http.StatusUnprocessableEntity,
        message: message,
    }
}

// NewLockedError creates a new API Error with status 423
func NewLockedError(message string) apiError {
    return apiError{
        status:  http.StatusLocked,
        message: message,
    }
}

// NewFailedDependencyError creates a new API Error with status 424
func NewFailedDependencyError(message string) apiError {
    return apiError{
        status:  http.StatusFailedDependency,
        message: message,
    }
}

// NewUpgradeRequiredError creates a new API Error with status 426
func NewUpgradeRequiredError(message string) apiError {
    return apiError{
        status:  http.StatusUpgradeRequired,
        message: message,
    }
}

// NewPreconditionRequiredError creates a new API Error with status 428
func NewPreconditionRequiredError(message string) apiError {
    return apiError{
        status:  http.StatusPreconditionRequired,
        message: message,
    }
}

// NewTooManyRequestsError creates a new API Error with status 429
func NewTooManyRequestsError(message string) apiError {
    return apiError{
        status:  http.StatusTooManyRequests,
        message: message,
    }
}

// NewRequestHeaderFieldsTooLargeError creates a new API Error with status 431
func NewRequestHeaderFieldsTooLargeError(message string) apiError {
    return apiError{
        status:  http.StatusRequestHeaderFieldsTooLarge,
        message: message,
    }
}

// NewUnavailableForLegalReasonsError creates a new API Error with status 451
func NewUnavailableForLegalReasonsError(message string) apiError {
    return apiError{
        status:  http.StatusUnavailableForLegalReasons,
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

// NewNotImplementedError creates a new API Error with status 501
func NewNotImplementedError(message string) apiError {
    return apiError{
        status:  http.StatusNotImplemented,
        message: message,
    }
}

// NewBadGatewayError creates a new API Error with status 502
func NewBadGatewayError(message string) apiError {
    return apiError{
        status:  http.StatusBadGateway,
        message: message,
    }
}

// NewServiceUnavailableError creates a new API Error with status 503
func NewServiceUnavailableError(message string) apiError {
    return apiError{
        status:  http.StatusServiceUnavailable,
        message: message,
    }
}

// NewGatewayTimeoutError creates a new API Error with status 504
func NewGatewayTimeoutError(message string) apiError {
    return apiError{
        status:  http.StatusGatewayTimeout,
        message: message,
    }
}

// NewHTTPVersionNotSupportedError creates a new API Error with status 505
func NewHTTPVersionNotSupportedError(message string) apiError {
    return apiError{
        status:  http.StatusHTTPVersionNotSupported,
        message: message,
    }
}

// NewVariantAlsoNegotiatesError creates a new API Error with status 506
func NewVariantAlsoNegotiatesError(message string) apiError {
    return apiError{
        status:  http.StatusVariantAlsoNegotiates,
        message: message,
    }
}

// NewInsufficientStorageError creates a new API Error with status 507
func NewInsufficientStorageError(message string) apiError {
    return apiError{
        status:  http.StatusInsufficientStorage,
        message: message,
    }
}

// NewLoopDetectedError creates a new API Error with status 508
func NewLoopDetectedError(message string) apiError {
    return apiError{
        status:  http.StatusLoopDetected,
        message: message,
    }
}

// NewNotExtendedError creates a new API Error with status 510
func NewNotExtendedError(message string) apiError {
    return apiError{
        status:  http.StatusNotExtended,
        message: message,
    }
}

// NewNetworkAuthenticationRequiredError creates a new API Error with status 511
func NewNetworkAuthenticationRequiredError(message string) apiError {
    return apiError{
        status:  http.StatusNetworkAuthenticationRequired,
        message: message,
    }
}
