package bookings

import (
	"github.com/gin-gonic/gin"
	"github.com/zdarovich/booking-api/handlers/constants"
	error2 "github.com/zdarovich/booking-api/handlers/errors"
	"github.com/zdarovich/booking-api/handlers/response"
	"github.com/zdarovich/booking-api/repositories/bookings"
	"github.com/zdarovich/booking-api/repositories/classes"
	classessUtil "github.com/zdarovich/booking-api/utils/classes"
	"net/http"
	"time"
)



type (
	Handler struct {
		BookingsRepository  bookings.IRepository
		ClassesRepository classes.IRepository
	}

	Booking struct {
		Date               string `json:"date"`
		Name                    string    `json:"name"`
	}

)

func (h *Handler) PostBooking(ctx *gin.Context) {
	var req Booking
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
	} else if len(req.Date) == 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeRequiredParameterMissing,
			Message: "Required parameter missing",
			Errors:[]error2.FieldError{
				{Field:   "date", Message: "Field is empty"},
			},
		})
		return
	}

	bookDate, err := time.Parse(constants.LayoutISO, req.Date)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeParameterInvalidFormat,
			Message: "Parameter is not 2006-01-02 format",
			Errors:[]error2.FieldError{
				{Field:   "date", Message: "Field is invalid"},
			},
		})
		return
	}
	now := time.Now()
	if now.Format(constants.LayoutISO) != req.Date && bookDate.Before(now) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeParameterInvalidFormat,
			Message: "Date is before valid time",
			Errors:[]error2.FieldError{
				{Field:   "date", Message: "Field is invalid"},
			},
		})
		return
	}
	c := h.ClassesRepository.GetClassBetweenDate(bookDate)
	if c == nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, &response.ErrorResponse{
			Code:    constants.CodeEntityNotFound,
			Message: "Class was not found",
		})
		return
	}

	classessUtil.DecreaseCapacity(bookDate, c)
	err = h.ClassesRepository.UpdateClass(c)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, &response.ErrorResponse{
			Code:        constants.CodeInternalError,
			Message:     "Server internal errors",
			Description: err.Error(),
		})
		return
	}

	b := new(bookings.Booking)
	b.Date = bookDate
	b.Name = req.Name
	err = h.BookingsRepository.SaveBooking(b)
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

func (h *Handler) GetBookings(ctx *gin.Context) {
	bs := h.BookingsRepository.GetBookings()
	ctx.JSON(http.StatusOK, &response.SuccessResponse{Data: bs})
}

