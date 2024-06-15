package errors

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

// stackTracer is the interface for stack trace
type stackTracer interface {
	StackTrace() errors.StackTrace
}

// BaseError is the abstraction over errors for a request
type BaseError interface {
	error
	GetStatusCode() int
	SetStatusCode(int)
	GetUserContext() interface{}
	SetUserContext(interface{})
	GetErrorContext() interface{}
	SetErrorContext(interface{})
	SetStackTrace(string)
	GetStackTrace() string
	GetMsg() string
	SetMsg(string)
	Error() string
}

// BaseErrorImpl implements BaseError and error
type BaseErrorImpl struct {
	// This is short error code to quickly identify the error class, user.TOO_MANY_REQUESTS. This is
	// sent back to the client
	Msg string
	// HTTP status code, eg 500, 404, 401, 429
	StatusCode int
	// Stacktrace for error handlers
	StackTrace string
	// Associated metadata with the error for detailed debugging. For example, it can include things
	// like resourceId very specific to the context of error. This is logged into tools like BigQuery
	// or Sentry. It is never sent to the client, so we can put any metadata required.
	ErrorContext interface{}
	// UserContext: Helpful information related to the error which can be sent to the client. Please
	// take care to not send any sensitive information in this field.
	UserContext interface{}

	error
}

// NewBaseError is the constructor for NewRequestImpl */
func NewBaseError(modifiers ...func(BaseError)) BaseError {
	re := &BaseErrorImpl{}
	for _, modifier := range modifiers {
		modifier(re)
	}
	if err, ok := errors.New(re.Msg).(stackTracer); ok && re.GetStackTrace() == "" {
		trace := strings.Builder{}
		for _, f := range err.StackTrace() {
			trace.WriteString(fmt.Sprintf("%+s:%d\n", f, f))
		}
		re.SetStackTrace(trace.String())
	}
	return re
}

// SetStackTrace sets the stack trace
func (re *BaseErrorImpl) SetStackTrace(trace string) {
	re.StackTrace = trace
}

// GetStackTrace returns back the stack trace
func (re *BaseErrorImpl) GetStackTrace() string {
	return re.StackTrace
}

// SetStatusCode sets the status code
func (re *BaseErrorImpl) SetStatusCode(statusCode int) {
	re.StatusCode = statusCode
}

// GetStatusCode returns back the status code
func (re *BaseErrorImpl) GetStatusCode() int {
	if re.StatusCode == 0 {
		return http.StatusInternalServerError
	}
	return re.StatusCode
}

// GetUserContext returns back the user context associated with the request
func (re *BaseErrorImpl) GetUserContext() interface{} {
	return re.UserContext
}

// SetUserContext sets the user context
func (re *BaseErrorImpl) SetUserContext(userContext interface{}) {
	re.UserContext = userContext
}

// GetErrorContext returns back the error context for internal logging
func (re *BaseErrorImpl) GetErrorContext() interface{} {
	return re.ErrorContext
}

// SetErrorContext sets the error context
func (re *BaseErrorImpl) SetErrorContext(errorContext interface{}) {
	re.ErrorContext = errorContext
}

// GetMsg returns back the error message
func (re *BaseErrorImpl) GetMsg() string {
	return re.Msg
}

// SetMsg sets the error message
func (re *BaseErrorImpl) SetMsg(msg string) {
	re.Msg = msg
}

// Error method to make BaseError compatible with error
func (re *BaseErrorImpl) Error() string {
	statusCode := re.GetStatusCode()
	userContext := re.GetUserContext()
	errorContext := re.GetErrorContext()
	stackTrace := re.GetStackTrace()
	msg := re.GetMsg()
	return fmt.Sprintf("BaseError:\nMsg: %s\nStatusCode: %d\nUserContext: %v\nErrorContext: %v\nStackTrace: %s", msg, statusCode, userContext, errorContext, stackTrace)
}

/* Modifiers for the BaseError */

// WithError constructs BaseError using incoming error
func WithError(err interface{}) func(BaseError) {
	return func(re BaseError) {
		if in, ok := err.(BaseError); ok {
			if in.GetStatusCode() != 0 {
				re.SetStatusCode(in.GetStatusCode())
			}
			if in.GetUserContext() != nil {
				re.SetUserContext(in.GetUserContext())
			}
			if in.GetErrorContext() != nil {
				re.SetErrorContext(in.GetErrorContext())
			}
			if in.GetStackTrace() != "" {
				re.SetStackTrace(in.GetStackTrace())
			}
			if in.GetMsg() != "" {
				re.SetMsg(in.GetMsg())
			}
		} else if in, ok := err.(error); ok {
			re.SetMsg(in.Error())
		} else if in, ok := err.(string); ok {
			re.SetMsg(in)
		} else {
			re.SetErrorContext(err)
		}
	}
}

// WithMsg sets message for the request error
func WithMsg(msg string) func(BaseError) {
	return func(re BaseError) {
		if re.GetMsg() == "" {
			re.SetMsg(msg)
		}
	}
}

// WithUserContext helps add user context to builder for BaseError
func WithUserContext(userContext interface{}) func(BaseError) {
	return func(re BaseError) {
		re.SetUserContext(userContext)
	}
}

// WithErrorContext helps add error context to builder for BaseError
func WithErrorContext(errorContext interface{}) func(BaseError) {
	return func(re BaseError) {
		re.SetErrorContext(errorContext)
	}
}

// WithStatusCode helps add HTTP status code to builder for BaseError
func WithStatusCode(statusCode int) func(BaseError) {
	return func(re BaseError) {
		re.SetStatusCode(statusCode)
	}
}

// WithStackTrace helps add stack trace to builder for BaseError
func WithStackTrace(stackTrace string) func(BaseError) {
	return func(re BaseError) {
		re.SetStackTrace(stackTrace)
	}
}
