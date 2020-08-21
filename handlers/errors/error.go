package errors

import "fmt"

type (
	CodeError struct {
		Code          int     `json:"code"`
		Message         string  `json:"message,omitempty"`
		Errors         []FieldError `json:"errors,omitempty"`
	}
	FieldError struct {
		Field         string  `json:"field,omitempty"`
		Message         string  `json:"message,omitempty"`
	}
)

func (e *CodeError) Error() string {
	return fmt.Sprintf("ErrorCode: %d ErrorMsg %s: ",e.Code, e.Message)
}

