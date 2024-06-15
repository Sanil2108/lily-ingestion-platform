package errors

import (
	"net/http"
)

// NewUnauthorizedError returns BaseError with HTTP status 401
func NewUnauthorizedError(modifiers ...func(BaseError)) BaseError {
	modifiers = append(modifiers, WithStatusCode(http.StatusUnauthorized))
	return NewBaseError(modifiers...)
}

// NewBadRequestError returns BaseError with HTTP status 400
func NewBadRequestError(modifiers ...func(BaseError)) BaseError {
	modifiers = append(modifiers, WithStatusCode(http.StatusBadRequest))
	return NewBaseError(modifiers...)
}

// NewInternalServerError returns BaseError with HTTP 500 and msg server.UKW
func NewInternalServerError(modifiers ...func(BaseError)) BaseError {
	modifiers = append(
		modifiers,
		WithStatusCode(http.StatusInternalServerError),
		WithMsg("server.UKW"),
	)

	re := NewBaseError(modifiers...)
	if re.GetMsg() == "" {
		re.SetMsg("server.UKW")
	}
	return re
}
