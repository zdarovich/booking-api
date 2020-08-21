package classes

import (
	"github.com/gin-gonic/gin"
	"github.com/zdarovich/booking-api/handlers/constants"
	error2 "github.com/zdarovich/booking-api/handlers/errors"
	"github.com/zdarovich/booking-api/handlers/response"
	"github.com/zdarovich/booking-api/repositories/classes"
	classessUtil "github.com/zdarovich/booking-api/utils/classes"
	"net/http"
	"time"
)

type (
	Handler struct {
		ClassesRepository classes.IRepository
	}

	Class struct {
		StartDate               string `json:"start_date"`
		EndDate                 string `json:"end_date"`
		Name                    string    `json:"name"`
		Capacity                    int    `json:"capacity"`
	}
)

func (h *Handler) PostClasses(ctx *gin.Context) {

	var req Class
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:        constants.CodeRequiredParameterMissing,
			Message:     "Invalid message format",
			Description: err.Error(),
		})
		return
	}

	if len(req.Name) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeRequiredParameterMissing,
			Message: "Required parameter missing",
			Errors:[]error2.FieldError{
				{Field:   "name", Message: "Field is empty"},
			},
		})
		return
	} else if len(req.StartDate) == 0 || len(req.EndDate) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeRequiredParameterMissing,
			Message: "Required parameter missing",
			Errors:[]error2.FieldError{
				{Field:   "start_date", Message: "Field is empty"},
				{Field:   "end_date", Message: "Field is empty"},
			},
		})
		return
	} else if req.Capacity <= 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeRequiredParameterMissing,
			Message: "Required parameter missing",
			Errors:[]error2.FieldError{
				{Field:   "capacity", Message: "Field is empty"},
			},
		})
		return
	}

	var start, end time.Time
	var err error
	start, err = time.Parse(constants.LayoutISO, req.StartDate)
	end, err = time.Parse(constants.LayoutISO, req.EndDate)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeParameterInvalidFormat,
			Message: "Parameter is not 2006-01-02 format",
			Errors:[]error2.FieldError{
				{Field:   "start_date", Message: "Field is invalid"},
				{Field:   "end_date", Message: "Field is invalid"},
			},
		})
		return
	}
	if start.Before(time.Now().AddDate(0, 0, -1)) || start.After(end.AddDate(0, 0, 1)) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeParameterInvalidFormat,
			Message: "Date is before valid time",
			Errors:[]error2.FieldError{
				{Field:   "start_date", Message: "Field is invalid"},
				{Field:   "end_date", Message: "Field is invalid"},
			},
		})
		return
	}

	c := h.ClassesRepository.GetClassBetweenDate(start)
	if c != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeEntityExistsAlready,
			Message: "Class exists already",
		})
		return
	}
	c = h.ClassesRepository.GetClassBetweenDate(end)
	if c != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeEntityExistsAlready,
			Message: "Class exists already",
		})
		return
	}

	class := classessUtil.CreateClass(req.Name, req.Capacity, start, end)

	err = h.ClassesRepository.SaveClass(class)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &response.ErrorResponse{
			Code:        constants.CodeInternalError,
			Message:     "Server internal errors",
			Description: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, &response.SuccessResponse{Data: req})
}

func (h *Handler) GetClasses(ctx *gin.Context) {
	cs := h.ClassesRepository.GetClasses()
	ctx.JSON(http.StatusOK, &response.SuccessResponse{Data: cs})
}
