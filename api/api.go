package api

import (
	"fmt"
)

type (
	// api struct
	api struct {
	}
	// IAPI interface
	IAPI interface {
		Run()
	}
)

func New() IAPI {

	return &api{
	}
}


// Run starts the api
func (api *api) Run() {

	router := NewRouter()

	router.Run(fmt.Sprintf(":%d", 8081))
}
