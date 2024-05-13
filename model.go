package errors

type CustomError struct {
	Message    string
	Code       string
	HttpStatus int
}
