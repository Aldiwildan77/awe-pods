package handler

// ValidationError for store error message
type ValidationError struct {
	Message string
}

func (validation ValidationError) Error() string {
	return validation.Message
}
