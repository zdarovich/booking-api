package bookings

import (
	"errors"
	"sync"
	"time"
)

type (
	// Repository struct
	Repository struct {
		lock sync.RWMutex
		Database      []*Booking
	}
	// IRepository interface
	IRepository interface {
		SaveBooking(b *Booking) error
		GetBookings() []*Booking
	}
	Booking struct {
		Date               time.Time `json:"start_date"`
		Name                    string    `json:"name"`
	}
)


func New() IRepository {

	return &Repository{
		Database:      []*Booking{},
	}
}

func (r *Repository) SaveBooking(b *Booking) error {
	if b == nil {
		return errors.New("booking is nil")
	}
	r.lock.Lock()
	defer r.lock.Unlock()
	r.Database = append(r.Database, b)
	return nil
}

func (r *Repository) GetBookings() []*Booking {
	r.lock.RLock()
	defer r.lock.RUnlock()
	return r.Database
}