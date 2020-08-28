package classes

import (
	"errors"
	"github.com/google/uuid"
	"sync"
	"time"
)

type (
	// Repository struct
	Repository struct {
		lock sync.RWMutex
		Database      []*Class
	}
	// IRepository interface
	IRepository interface {
		GetClassBetweenDate(date time.Time) *Class
		GetClasses() []*Class
		SaveClass(c *Class) error
		UpdateClass(c *Class) error
	}
	Class struct {
		Id string
		StartDate               time.Time `json:"start_date"`
		EndDate                 time.Time `json:"end_date"`
		Name                    string    `json:"name"`
		Capacity                    []int    `json:"capacity"`
	}
)

func New() IRepository {

	return &Repository{
		Database:      []*Class{},
	}
}

func (r *Repository) GetClassBetweenDate(date time.Time) *Class {
	if date.IsZero() {
		return nil
	}
	r.lock.RLock()
	defer r.lock.RUnlock()
	for _, c := range r.Database {
		if c.StartDate.Equal(date) || c.EndDate.Equal(date) {
			return c
		} else if date.After(c.StartDate) && date.Before(c.EndDate) {
			return c
		}
	}
	return nil
}

func (r *Repository) GetClasses() []*Class {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.Database
}


func (r *Repository) SaveClass(c *Class) error {
	if c == nil {
		return errors.New("class is nil")
	}
	c.Id = uuid.New().String()
	r.lock.Lock()
	defer r.lock.Unlock()
	r.Database = append(r.Database, c)
	return nil
}

func (r *Repository) UpdateClass(c *Class) error {
	if c == nil {
		return errors.New("class is nil")
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	idx := -1
	for index, class := range r.Database {
		if class.Id == c.Id {
			idx = index
			break
		}
	}
	if idx == -1 {
		return errors.New("class is not found")
	}
	r.Database[idx] = c
	return nil
}