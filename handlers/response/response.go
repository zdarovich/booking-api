package response

import error2 "github.com/zdarovich/booking-api/handlers/errors"

type (

	SuccessResponse struct {
		Data interface{} `json:"data"`
	}
	ErrorResponse struct {
		Code          int                  `json:"code"`
		Message         string             `json:"message,omitempty"`
		Description         string         `json:"desc,omitempty"`
		Errors         []error2.FieldError `json:"errors,omitempty"`
	}

)


