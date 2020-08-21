package api

import (
	"github.com/gin-gonic/gin"
	"github.com/zdarovich/booking-api/handlers/bookings"
	"github.com/zdarovich/booking-api/handlers/classes"
	bRepo "github.com/zdarovich/booking-api/repositories/bookings"
	cRepo "github.com/zdarovich/booking-api/repositories/classes"
)

func NewRouter() *gin.Engine {
	bookingsRepo := bRepo.New()
	classesRepo := cRepo.New()
	bookingsHandler := bookings.Handler{
		BookingsRepository: bookingsRepo,
		ClassesRepository:  classesRepo,
	}
	classesHandler := classes.Handler{
		ClassesRepository:  classesRepo,
	}

	router := gin.Default()
	router.POST("/classes", classesHandler.PostClasses)
	router.GET("/classes", classesHandler.GetClasses)
	router.POST("/bookings", bookingsHandler.PostBooking)
	router.GET("/bookings", bookingsHandler.GetBookings)
	return router
}