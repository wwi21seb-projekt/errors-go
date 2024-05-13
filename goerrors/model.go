package goerrors

type CustomError struct {
	Title      string `json:"title"`
	Code       string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"http_status"`
}
