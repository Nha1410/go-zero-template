package errors

import (
	"errors"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	// Common errors
	ErrNotFound      = NewError(http.StatusNotFound, "NOT_FOUND", "Resource not found")
	ErrBadRequest    = NewError(http.StatusBadRequest, "BAD_REQUEST", "Invalid request")
	ErrUnauthorized  = NewError(http.StatusUnauthorized, "UNAUTHORIZED", "Unauthorized")
	ErrForbidden     = NewError(http.StatusForbidden, "FORBIDDEN", "Forbidden")
	ErrInternalError = NewError(http.StatusInternalServerError, "INTERNAL_ERROR", "Internal server error")
	ErrConflict      = NewError(http.StatusConflict, "CONFLICT", "Resource conflict")
)

// Error represents a custom error with HTTP status code and error code
type Error struct {
	StatusCode int    `json:"status_code"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	Details    string `json:"details,omitempty"`
}

func (e *Error) Error() string {
	return e.Message
}

// NewError creates a new error
func NewError(statusCode int, code, message string) *Error {
	return &Error{
		StatusCode: statusCode,
		Code:       code,
		Message:    message,
	}
}

// WithDetails adds details to the error
func (e *Error) WithDetails(details string) *Error {
	return &Error{
		StatusCode: e.StatusCode,
		Code:       e.Code,
		Message:    e.Message,
		Details:    details,
	}
}

// ToGRPCError converts HTTP error to gRPC error
func ToGRPCError(err error) error {
	if err == nil {
		return nil
	}

	var customErr *Error
	if errors.As(err, &customErr) {
		grpcCode := httpStatusToGRPCCode(customErr.StatusCode)
		return status.Error(grpcCode, customErr.Message)
	}

	logx.Errorf("Unknown error type: %v", err)
	return status.Error(codes.Internal, "Internal server error")
}

// httpStatusToGRPCCode converts HTTP status code to gRPC code
func httpStatusToGRPCCode(httpCode int) codes.Code {
	switch httpCode {
	case http.StatusOK:
		return codes.OK
	case http.StatusBadRequest:
		return codes.InvalidArgument
	case http.StatusUnauthorized:
		return codes.Unauthenticated
	case http.StatusForbidden:
		return codes.PermissionDenied
	case http.StatusNotFound:
		return codes.NotFound
	case http.StatusConflict:
		return codes.AlreadyExists
	case http.StatusInternalServerError:
		return codes.Internal
	case http.StatusServiceUnavailable:
		return codes.Unavailable
	default:
		return codes.Internal
	}
}

