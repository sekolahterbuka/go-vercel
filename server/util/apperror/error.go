package apperrors

import (
	"errors"
	"fmt"
	"net/http"
)

// Type holds a type string and integer code for the error
type Type string

// "Set" of valid errorTypes
const (
	Authorization        Type = "AUTHORIZATION"          // Authentication Failures -
	BadRequest           Type = "BAD_REQUEST"            // Validation errors / BadInput
	Conflict             Type = "CONFLICT"               // Already exists (eg, create account with existent email) - 409
	Internal             Type = "INTERNAL"               // Server (500) and fallback errors
	NotFound             Type = "NOT_FOUND"              // For not finding resource
	PayloadTooLarge      Type = "PAYLOAD_TOO_LARGE"      // for uploading tons of JSON, or an image over the limit - 413
	ServiceUnavailable   Type = "SERVICE_UNAVAILABLE"    // For long running handlers
	UnsupportedMediaType Type = "UNSUPPORTED_MEDIA_TYPE" // for http 415
)

// Error holds a custom error for the application
// which is helpful in returning a consistent
// error type/message from API endpoints
type Error struct {
	Type    Type   `json:"type"`
	Message string `json:"message"`
}

// Error satisfies standard error interface
// we can return errors from this package as
// a regular old go _error_
func (e *Error) Error() string {
	return e.Message
}

// Status is a mapping errors to status codes
// Of course, this is somewhat redundant since
// our errors already map http status codes
func (e *Error) Status() int {
	switch e.Type {
	case Authorization:
		return http.StatusUnauthorized
	case BadRequest:
		return http.StatusBadRequest
	case Conflict:
		return http.StatusConflict
	case Internal:
		return http.StatusInternalServerError
	case NotFound:
		return http.StatusNotFound
	case PayloadTooLarge:
		return http.StatusRequestEntityTooLarge
	case ServiceUnavailable:
		return http.StatusServiceUnavailable
	case UnsupportedMediaType:
		return http.StatusUnsupportedMediaType
	default:
		return http.StatusInternalServerError
	}
}

// Status checks the runtime type
// of the error and returns an http
// status code if the error is model.Error
func Status(err error) int {
	var e *Error
	if errors.As(err, &e) {
		return e.Status()
	}
	return http.StatusInternalServerError
}

/*
* Error "Factories"
 */

// NewAuthorization to create a 401
func getAuthorization(reason string) *Error {
	return &Error{
		Type:    Authorization,
		Message: reason,
	}
}

func NewAuthorization(w http.ResponseWriter, reason string) {
	w.WriteHeader(getAuthorization(reason).Status())
	w.Write([]byte(getAuthorization(reason).Message))
}

// NewBadRequest to create 400 errors (validation, for example)
func getBadRequest(reason string) *Error {
	return &Error{
		Type:    BadRequest,
		Message: fmt.Sprintf("Bad request. Reason: %v", reason),
	}
}

func NewBadRequest(w http.ResponseWriter, reason string) {
	w.WriteHeader(getBadRequest(reason).Status())
	w.Write([]byte(getBadRequest(reason).Message))
}

// NewConflict to create an error for 409
func getConflict(name string, value string) *Error {
	return &Error{
		Type:    Conflict,
		Message: fmt.Sprintf("resource: %v with value: %v already exists", name, value),
	}
}

func NewConflict(w http.ResponseWriter, name string, value string) {
	w.WriteHeader(getConflict(name, value).Status())
	w.Write([]byte(getConflict(name, value).Message))
}

// NewInternal for 500 errors and unknown errors
func getInternal() *Error {
	return &Error{
		Type:    Internal,
		Message: fmt.Sprintf("Internal server error."),
	}
}

func NewInternal(w http.ResponseWriter) {
	w.WriteHeader(getInternal().Status())
	w.Write([]byte(getInternal().Message))
}

// NewNotFound to create an error for 404
func getNotFound(name string, value string) *Error {
	return &Error{
		Type:    NotFound,
		Message: fmt.Sprintf("resource: %v with value: %v not found", name, value),
	}
}

func NewNotFound(w http.ResponseWriter, name string, value string) {
	w.WriteHeader(getNotFound(name, value).Status())
	w.Write([]byte(getNotFound(name, value).Message))
}

// NewPayloadTooLarge to create an error for 413
func getPayloadTooLarge(maxBodySize int64, contentLength int64) *Error {
	return &Error{
		Type:    PayloadTooLarge,
		Message: fmt.Sprintf("Max payload size of %v exceeded. Actual payload size: %v", maxBodySize, contentLength),
	}
}

func NewPayloadTooLarge(w http.ResponseWriter, maxBodySize int64, contentLength int64) {
	w.WriteHeader(getPayloadTooLarge(maxBodySize, contentLength).Status())
	w.Write([]byte(getPayloadTooLarge(maxBodySize, contentLength).Message))
}

// NewServiceUnavailable to create an error for 503

func getServiceUnavailable() *Error {
	return &Error{
		Type:    ServiceUnavailable,
		Message: fmt.Sprintf("Service unavailable or timed out"),
	}
}

func NewServiceUnavailable(w http.ResponseWriter) {
	w.WriteHeader(getServiceUnavailable().Status())
	w.Write([]byte(getServiceUnavailable().Message))
}

// NewUnsupportedMediaType to create an error for 415
func getUnsupportedMediaType(reason string) *Error {
	return &Error{
		Type:    UnsupportedMediaType,
		Message: reason,
	}
}

func NewUnsupportedMediaType(w http.ResponseWriter, reason string) {
	w.WriteHeader(getUnsupportedMediaType(reason).Status())
	w.Write([]byte(getUnsupportedMediaType(reason).Message))
}
