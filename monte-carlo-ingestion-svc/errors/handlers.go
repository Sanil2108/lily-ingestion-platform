package errors

// HandleError handles any input - converting it into BaseError
func HandleError(err interface{}) BaseError {
	var requestError BaseError
	existingError, ok := err.(BaseError)

	if !ok {
		requestError = NewInternalServerError(WithErrorContext(err))
	} else {
		requestError = existingError
	}

	return requestError
}
