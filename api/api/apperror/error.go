package apperror

type AppError struct {
	Code        int
	Description string
	Err         error
}

func (e *AppError) Error() string {
	return e.Description
}

func NewAppError(code int, description string, err error) *AppError {
	return &AppError{Code: code, Description: description, Err: err}
}
