package apierrors

import (
    "fmt"
    "github.com/stretchr/testify/assert"
    "net/http"
    "testing"
)

func TestNewBadRequestError(t *testing.T) {
    apiErr := NewBadRequestError("test error")
    assert.Equal(t, http.StatusBadRequest, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusBadRequest), apiErr.Error())
}

func TestNewUnauthorizedError(t *testing.T) {
    apiErr := NewUnauthorizedError("test error")
    assert.Equal(t, http.StatusUnauthorized, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusUnauthorized), apiErr.Error())
}

func TestNewPaymentRequiredError(t *testing.T) {
    apiErr := NewPaymentRequiredError("test error")
    assert.Equal(t, http.StatusPaymentRequired, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusPaymentRequired), apiErr.Error())
}

func TestNewForbiddenError(t *testing.T) {
    apiErr := NewForbiddenError("test error")
    assert.Equal(t, http.StatusForbidden, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusForbidden), apiErr.Error())
}

func TestNewNotFoundError(t *testing.T) {
    apiErr := NewNotFoundError("test error")
    assert.Equal(t, http.StatusNotFound, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusNotFound), apiErr.Error())
}

func TestNewMethodNotAllowedError(t *testing.T) {
    apiErr := NewMethodNotAllowedError("test error")
    assert.Equal(t, http.StatusMethodNotAllowed, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusMethodNotAllowed), apiErr.Error())
}

func TestNewNotAcceptableError(t *testing.T) {
    apiErr := NewNotAcceptableError("test error")
    assert.Equal(t, http.StatusNotAcceptable, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusNotAcceptable), apiErr.Error())
}

func TestNewProxyAuthenticationRequiredError(t *testing.T) {
    apiErr := NewProxyAuthenticationRequiredError("test error")
    assert.Equal(t, http.StatusProxyAuthRequired, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusProxyAuthRequired), apiErr.Error())
}

func TestNewRequestTimeoutError(t *testing.T) {
    apiErr := NewRequestTimeoutError("test error")
    assert.Equal(t, http.StatusRequestTimeout, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusRequestTimeout), apiErr.Error())
}

func TestNewConflictError(t *testing.T) {
    apiErr := NewConflictError("test error")
    assert.Equal(t, http.StatusConflict, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusConflict), apiErr.Error())
}

func TestNewGoneError(t *testing.T) {
    apiErr := NewGoneError("test error")
    assert.Equal(t, http.StatusGone, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusGone), apiErr.Error())
}

func TestNewLengthRequiredError(t *testing.T) {
    apiErr := NewLengthRequiredError("test error")
    assert.Equal(t, http.StatusLengthRequired, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusLengthRequired), apiErr.Error())
}

func TestNewPreconditionFailedError(t *testing.T) {
    apiErr := NewPreconditionFailedError("test error")
    assert.Equal(t, http.StatusPreconditionFailed, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusPreconditionFailed), apiErr.Error())
}

func TestNewPayloadTooLargeError(t *testing.T) {
    apiErr := NewPayloadTooLargeError("test error")
    assert.Equal(t, http.StatusRequestEntityTooLarge, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusRequestEntityTooLarge), apiErr.Error())
}

func TestNewURITooLongError(t *testing.T) {
    apiErr := NewURITooLongError("test error")
    assert.Equal(t, http.StatusRequestURITooLong, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusRequestURITooLong), apiErr.Error())
}

func TestNewUnsupportedMediaTypeError(t *testing.T) {
    apiErr := NewUnsupportedMediaTypeError("test error")
    assert.Equal(t, http.StatusUnsupportedMediaType, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusUnsupportedMediaType), apiErr.Error())
}

func TestNewRequestedRangeNotSatisfiableError(t *testing.T) {
    apiErr := NewRequestedRangeNotSatisfiableError("test error")
    assert.Equal(t, http.StatusRequestedRangeNotSatisfiable, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusRequestedRangeNotSatisfiable), apiErr.Error())
}

func TestNewExpectationFailedError(t *testing.T) {
    apiErr := NewExpectationFailedError("test error")
    assert.Equal(t, http.StatusExpectationFailed, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusExpectationFailed), apiErr.Error())
}

func TestNewImATeapotError(t *testing.T) {
    apiErr := NewImATeapotError("test error")
    assert.Equal(t, http.StatusTeapot, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusTeapot), apiErr.Error())
}

func TestNewMisdirectedRequestError(t *testing.T) {
    apiErr := NewMisdirectedRequestError("test error")
    assert.Equal(t, http.StatusMisdirectedRequest, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusMisdirectedRequest), apiErr.Error())
}

func TestNewUnprocessableEntityError(t *testing.T) {
    apiErr := NewUnprocessableEntityError("test error")
    assert.Equal(t, http.StatusUnprocessableEntity, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusUnprocessableEntity), apiErr.Error())
}

func TestNewLockedError(t *testing.T) {
    apiErr := NewLockedError("test error")
    assert.Equal(t, http.StatusLocked, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusLocked), apiErr.Error())
}

func TestNewFailedDependencyError(t *testing.T) {
    apiErr := NewFailedDependencyError("test error")
    assert.Equal(t, http.StatusFailedDependency, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusFailedDependency), apiErr.Error())
}

func TestNewUpgradeRequiredError(t *testing.T) {
    apiErr := NewUpgradeRequiredError("test error")
    assert.Equal(t, http.StatusUpgradeRequired, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusUpgradeRequired), apiErr.Error())
}

func TestNewPreconditionRequiredError(t *testing.T) {
    apiErr := NewPreconditionRequiredError("test error")
    assert.Equal(t, http.StatusPreconditionRequired, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusPreconditionRequired), apiErr.Error())
}

func TestNewTooManyRequestsError(t *testing.T) {
    apiErr := NewTooManyRequestsError("test error")
    assert.Equal(t, http.StatusTooManyRequests, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusTooManyRequests), apiErr.Error())
}

func TestNewRequestHeaderFieldsTooLargeError(t *testing.T) {
    apiErr := NewRequestHeaderFieldsTooLargeError("test error")
    assert.Equal(t, http.StatusRequestHeaderFieldsTooLarge, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusRequestHeaderFieldsTooLarge), apiErr.Error())
}

func TestNewUnavailableForLegalReasonsError(t *testing.T) {
    apiErr := NewUnavailableForLegalReasonsError("test error")
    assert.Equal(t, http.StatusUnavailableForLegalReasons, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusUnavailableForLegalReasons), apiErr.Error())
}

func TestNewInternalServerError(t *testing.T) {
    apiErr := NewInternalServerError("test error")
    assert.Equal(t, http.StatusInternalServerError, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusInternalServerError), apiErr.Error())
}

func TestNewNotImplementedError(t *testing.T) {
    apiErr := NewNotImplementedError("test error")
    assert.Equal(t, http.StatusNotImplemented, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusNotImplemented), apiErr.Error())
}

func TestNewBadGatewayError(t *testing.T) {
    apiErr := NewBadGatewayError("test error")
    assert.Equal(t, http.StatusBadGateway, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusBadGateway), apiErr.Error())
}

func TestNewServiceUnavailableError(t *testing.T) {
    apiErr := NewServiceUnavailableError("test error")
    assert.Equal(t, http.StatusServiceUnavailable, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusServiceUnavailable), apiErr.Error())
}

func TestNewGatewayTimeoutError(t *testing.T) {
    apiErr := NewGatewayTimeoutError("test error")
    assert.Equal(t, http.StatusGatewayTimeout, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusGatewayTimeout), apiErr.Error())
}

func TestNewHTTPVersionNotSupportedError(t *testing.T) {
    apiErr := NewHTTPVersionNotSupportedError("test error")
    assert.Equal(t, http.StatusHTTPVersionNotSupported, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusHTTPVersionNotSupported), apiErr.Error())
}

func TestNewVariantAlsoNegotiatesError(t *testing.T) {
    apiErr := NewVariantAlsoNegotiatesError("test error")
    assert.Equal(t, http.StatusVariantAlsoNegotiates, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusVariantAlsoNegotiates), apiErr.Error())
}

func TestNewInsufficientStorageError(t *testing.T) {
    apiErr := NewInsufficientStorageError("test error")
    assert.Equal(t, http.StatusInsufficientStorage, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusInsufficientStorage), apiErr.Error())
}

func TestNewLoopDetectedError(t *testing.T) {
    apiErr := NewLoopDetectedError("test error")
    assert.Equal(t, http.StatusLoopDetected, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusLoopDetected), apiErr.Error())
}

func TestNewNotExtendedError(t *testing.T) {
    apiErr := NewNotExtendedError("test error")
    assert.Equal(t, http.StatusNotExtended, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusNotExtended), apiErr.Error())
}

func TestNewNetworkAuthenticationRequiredError(t *testing.T) {
    apiErr := NewNetworkAuthenticationRequiredError("test error")
    assert.Equal(t, http.StatusNetworkAuthenticationRequired, apiErr.Status())
    assert.Equal(t, "test error", apiErr.Message())
    assert.Equal(t, fmt.Sprintf("Error %d: test error", http.StatusNetworkAuthenticationRequired), apiErr.Error())
}
