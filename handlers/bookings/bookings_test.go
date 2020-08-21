package bookings

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zdarovich/booking-api/handlers/constants"
	"github.com/zdarovich/booking-api/handlers/response"
	"github.com/zdarovich/booking-api/repositories/bookings"
	"github.com/zdarovich/booking-api/repositories/classes"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_PostBooking_WithValidBookingAndClass_ShouldReturn201(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	classesRepo := classes.New()

	newClass := &classes.Class{
		StartDate: time.Now().Add(-24 * time.Hour),
		EndDate:   time.Now().Add(24 * time.Hour),
		Name:      "Football",
		Capacity:  []int{10, 10, 10},
	}
	err := classesRepo.SaveClass(newClass)
	assert.NilError(t, err)

	bookingsRepo := bookings.New()

	bookingsHandler := Handler{
		ClassesRepository:  classesRepo,
		BookingsRepository:  bookingsRepo,
	}
	r.POST("/bookings", bookingsHandler.PostBooking)

	date := time.Now()
	name := "Rob Pike"
	c := Booking{
		Date: date.Format(constants.LayoutISO),
		Name:      name,
	}
	fmt.Printf("%+v \n", c)

	b, err := json.Marshal(&c)
	assert.NilError(t, err)

	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(b))
	assert.NilError(t, err)

	r.ServeHTTP(w, req)
	res, err := ioutil.ReadAll(w.Body)
	assert.NilError(t, err)
	fmt.Println(string(res))
	assert.Equal(t, w.Code, http.StatusCreated)

	type Response struct {
		Data Booking
	}
	var resp Response
	err = json.Unmarshal(res, &resp)
	assert.NilError(t, err)

	assert.Equal(t, resp.Data.Date, date.Format(constants.LayoutISO))
	assert.Equal(t, resp.Data.Name, name)

	created := classesRepo.GetClassBetweenDate(date)
	assert.DeepEqual(t, created.Capacity, []int{10, 9, 10})

}

func TestHandler_PostBooking_WithUnexistentClass_ShouldReturn400AndError(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	classesRepo := classes.New()
	bookingsRepo := bookings.New()

	bookingsHandler := Handler{
		ClassesRepository:  classesRepo,
		BookingsRepository:  bookingsRepo,
	}
	r.POST("/bookings", bookingsHandler.PostBooking)

	date := time.Now().Add(1 * time.Hour)
	name := "Rob Pike"
	c := Booking{
		Date: date.Format(constants.LayoutISO),
		Name:      name,
	}
	b, err := json.Marshal(&c)
	assert.NilError(t, err)

	req, err := http.NewRequest("POST", "/bookings", bytes.NewBuffer(b))
	assert.NilError(t, err)

	r.ServeHTTP(w, req)
	res, err := ioutil.ReadAll(w.Body)
	assert.NilError(t, err)
	fmt.Println(string(res))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	expected := response.ErrorResponse{
		Code:    constants.CodeEntityNotFound,
		Message: "Class was not found",
	}
	var resp response.ErrorResponse
	err = json.Unmarshal(res, &resp)
	assert.NilError(t, err)

	assert.DeepEqual(t, expected, resp)
}
