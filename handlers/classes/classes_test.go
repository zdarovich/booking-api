package classes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zdarovich/booking-api/handlers/constants"
	"github.com/zdarovich/booking-api/handlers/errors"
	"github.com/zdarovich/booking-api/handlers/response"
	"github.com/zdarovich/booking-api/repositories/classes"
	"gotest.tools/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestHandler_PostClasses_WithValidClass_ShouldReturn201(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	classesRepo := classes.New()

	classesHandler := Handler{
		ClassesRepository:  classesRepo,
	}
	r.POST("/classes", classesHandler.PostClasses)

	start := time.Now().Add(1 * time.Hour)
	end := time.Now().Add(2 * time.Hour)
	name := "Pillates"
	capacity := 10
	c := Class{
		StartDate: start.Format(constants.LayoutISO),
		EndDate:   end.Format(constants.LayoutISO),
		Name:      name,
		Capacity: capacity,
	}
	b, err := json.Marshal(&c)
	assert.NilError(t, err)

	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(b))
	assert.NilError(t, err)

	r.ServeHTTP(w, req)
	res, err := ioutil.ReadAll(w.Body)
	assert.NilError(t, err)
	fmt.Println(string(res))
	assert.Equal(t, w.Code, http.StatusCreated)

	type Response struct {
		Data Class
	}
	var resp Response
	err = json.Unmarshal(res, &resp)
	assert.NilError(t, err)

	assert.Equal(t, resp.Data.StartDate, start.Format(constants.LayoutISO))
	assert.Equal(t, resp.Data.EndDate, end.Format(constants.LayoutISO))
	assert.Equal(t, resp.Data.Name, name)
	assert.Equal(t, resp.Data.Capacity, capacity)
}

func TestHandler_PostClasses_WithWrongDate_ShouldReturn400AndError(t *testing.T) {
	w := httptest.NewRecorder()
	r := gin.Default()
	classesRepo := classes.New()

	classesHandler := Handler{
		ClassesRepository:  classesRepo,
	}
	r.POST("/classes", classesHandler.PostClasses)

	end := time.Now().Add(2 * time.Hour)
	name := "Pillates"
	capacity := 10
	c := Class{
		EndDate:   end.Format(constants.LayoutISO),
		Name:      name,
		Capacity: capacity,
	}
	b, err := json.Marshal(&c)
	assert.NilError(t, err)

	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(b))
	assert.NilError(t, err)

	r.ServeHTTP(w, req)
	res, err := ioutil.ReadAll(w.Body)
	assert.NilError(t, err)
	fmt.Println(string(res))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	expected := response.ErrorResponse{
		Code:    constants.CodeRequiredParameterMissing,
		Message: "Required parameter missing",
		Errors:[]errors.FieldError{
			{Field:   "start_date", Message: "Field is empty"},
			{Field:   "end_date", Message: "Field is empty"},
		},
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(res, &resp)
	assert.NilError(t, err)

	assert.DeepEqual(t, resp, expected)
}

func TestHandler_PostClasses_WithDuplicateClass_ShouldReturn400AndError(t *testing.T) {
	r := gin.Default()
	classesRepo := classes.New()

	classesHandler := Handler{
		ClassesRepository:  classesRepo,
	}
	r.POST("/classes", classesHandler.PostClasses)

	start := time.Now().Add(1 * time.Hour)
	end := time.Now().Add(2 * time.Hour)
	name := "Pillates"
	capacity := 10
	c := Class{
		StartDate: start.Format(constants.LayoutISO),
		EndDate:   end.Format(constants.LayoutISO),
		Name:      name,
		Capacity: capacity,
	}
	b, err := json.Marshal(&c)
	assert.NilError(t, err)

	w := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/classes", bytes.NewBuffer(b))
	assert.NilError(t, err)
	r.ServeHTTP(w, req)
	res, err := ioutil.ReadAll(w.Body)
	assert.NilError(t, err)
	fmt.Println(string(res))
	assert.Equal(t, w.Code, http.StatusCreated)

	w = httptest.NewRecorder()
	req, err = http.NewRequest("POST", "/classes", bytes.NewBuffer(b))
	assert.NilError(t, err)
	r.ServeHTTP(w, req)
	res, err = ioutil.ReadAll(w.Body)
	assert.NilError(t, err)
	fmt.Println(string(res))
	assert.Equal(t, w.Code, http.StatusBadRequest)

	expected := response.ErrorResponse{
		Code:    constants.CodeEntityExistsAlready,
		Message: "Class exists already",
	}

	var resp response.ErrorResponse
	err = json.Unmarshal(res, &resp)
	assert.NilError(t, err)

	assert.DeepEqual(t, resp, expected)
}
