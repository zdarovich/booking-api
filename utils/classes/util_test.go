package classes

import (
	"fmt"
	"gotest.tools/assert"
	"testing"
	"time"
)

func TestCreateClass_WithOneHourOffset_ShouldCreateOneLengthCapacity(t *testing.T) {
	name := "test"
	capacity := 10
	start := time.Now()
	end := start.Add(1 * time.Hour)
	actual := CreateClass(name, capacity, start, end)

	expected := []int{10}
	assert.Equal(t, actual.Name, name)
	assert.DeepEqual(t, actual.Capacity, expected)
	assert.Equal(t, actual.StartDate, start)
	assert.Equal(t, actual.EndDate, end)
}

func TestCreateClass_WithSevenDaysOffset_ShouldCreateSevenLengthCapacity(t *testing.T) {
	name := "test"
	capacity := 10
	start := time.Now()
	end := start.AddDate(0, 0, 7)
	actual := CreateClass(name, capacity, start, end)

	expected := []int{10, 10, 10, 10, 10, 10, 10}
	assert.Equal(t, actual.Name, name)
	assert.DeepEqual(t, actual.Capacity, expected)
	assert.Equal(t, actual.StartDate, start)
	assert.Equal(t, actual.EndDate, end)
}

func TestDecreaseCapacity_WithFifthDayBooking_ShouldDecreaseCapacityByOne(t *testing.T) {
	name := "test"
	capacity := 10
	start := time.Now()
	end := start.AddDate(0, 0, 7)
	actual := CreateClass(name, capacity, start, end)

	fifthDay := start.AddDate(0, 0, 4)

	expected := []int{10, 10, 10, 10, 9, 10, 10}

	DecreaseCapacity(fifthDay, actual)
	fmt.Print(actual.Capacity)
	assert.Equal(t, actual.Name, name)
	assert.DeepEqual(t, actual.Capacity, expected)
	assert.Equal(t, actual.StartDate, start)
	assert.Equal(t, actual.EndDate, end)
}

func TestDecreaseCapacity_WithTodayBooking_ShouldDecreaseCapacityByOne(t *testing.T) {
	name := "test"
	capacity := 10
	start := time.Now()
	end := start.AddDate(0, 0, 7)
	actual := CreateClass(name, capacity, start, end)

	expected := []int{9, 10, 10, 10, 10, 10, 10}

	DecreaseCapacity(start, actual)
	fmt.Print(actual.Capacity)
	assert.Equal(t, actual.Name, name)
	assert.DeepEqual(t, actual.Capacity, expected)
	assert.Equal(t, actual.StartDate, start)
	assert.Equal(t, actual.EndDate, end)
}
