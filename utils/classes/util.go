package classes

import (
	"github.com/zdarovich/booking-api/repositories/classes"
	"time"
)


func CreateClass(name string, capacity int, start, end time.Time) *classes.Class {
	class := new(classes.Class)
	class.StartDate = start
	class.EndDate = end
	class.Name = name

	days := DaysBetween(start, end)
	if days == 0 {
		days = 1
	}
	capArr := make([]int, days)
	for i := 0; i < len(capArr); i ++ {
		capArr[i] = capacity
	}
	class.Capacity = capArr
	return class
}

func DecreaseCapacity(date time.Time, class *classes.Class) {
	var idx, capacity int
	if date.Equal(class.StartDate) {
		capacity = class.Capacity[0]
		idx = 0
	} else if date.Equal(class.EndDate) {
		capacity = class.Capacity[len(class.Capacity) - 1]
		idx = len(class.Capacity) - 1
	} else {
		days := DaysBetween(class.StartDate, class.EndDate)

		offsetDays := DaysBetween(date, class.EndDate)

		idx = days - offsetDays
		capacity = class.Capacity[idx]
	}
	if capacity == 0 {
		return
	}
	capacity--
	class.Capacity[idx] = capacity
}


func DaysBetween(a, b time.Time) int {
	if a.After(b) {
		a, b = b, a
	}

	days := -a.YearDay()
	for year := a.Year(); year < b.Year(); year++ {
		days += time.Date(year, time.December, 31, 0, 0, 0, 0, time.UTC).YearDay()
	}
	days += b.YearDay()

	return days
}
