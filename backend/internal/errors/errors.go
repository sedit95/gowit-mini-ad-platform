package errors

import "fmt"

// AppError represents an application-level error.
type AppError struct {
	Code    string
	Message string
}

func (e *AppError) Error() string {
	return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

// Predefined error codes
const (
	CodeValidationError  = "validation_error"
	CodeNotFound         = "not_found"
	CodeInvalidStatus    = "invalid_status"
	CodeCampaignNotActive = "campaign_not_active"
	CodeBudgetExhausted  = "budget_exhausted"
	CodeInternalError    = "internal_error"
)

// New creates a new AppError.
func New(code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

// Validation creates a validation_error.
func Validation(message string) *AppError {
	return &AppError{Code: CodeValidationError, Message: message}
}

// NotFound creates a not_found error.
func NotFound(message string) *AppError {
	return &AppError{Code: CodeNotFound, Message: message}
}

// Internal creates an internal_error.
func Internal(message string) *AppError {
	return &AppError{Code: CodeInternalError, Message: message}
}
