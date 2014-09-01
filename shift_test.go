package schedule

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestIfShiftIsShift(t *testing.T) {
	startsAt, e := time.Parse(time.RFC3339, "2016-04-10T12:00:00+02:00")
	EndsAt, e := time.Parse(time.RFC3339, "2016-04-10T16:00:00+02:00")
	if e != nil {
		fmt.Println(e)
	}
	worker := Worker{1, "kalle", []DayRange{DayRange{1, 800, 1600}}}
	workers := []Worker{worker}
	shift := Shift{1, startsAt, EndsAt, workers}
	assert.Equal(t, shift.Id, 1)
	assert.Equal(t, shift.StartsAt, startsAt)
	assert.Equal(t, shift.EndsAt, EndsAt)
	assert.Equal(t, shift.Workers[0], worker)
}

func TestShiftJsonDecoding(t *testing.T) {
	b := []byte(`{"id":1,"starts_at":"2016-04-10T12:00:00+02:00","ends_at":"2016-04-10T16:00:00+02:00"}`)
	startsAt, e := time.Parse(time.RFC3339, "2016-04-10T12:00:00+02:00")
	EndsAt, e := time.Parse(time.RFC3339, "2016-04-10T16:00:00+02:00")
	if e != nil {
		fmt.Println(e)
	}
	var shift Shift
	err := json.Unmarshal(b, &shift)
	assert.Nil(t, err)
	assert.Equal(t, shift.Id, 1)
	assert.Equal(t, shift.StartsAt, startsAt)
	assert.Equal(t, shift.EndsAt, EndsAt)
}

func TestIsAssigned(t *testing.T) {
	startsAt, e := time.Parse(time.RFC3339, "2016-04-10T12:00:00+02:00")
	endsAt, e := time.Parse(time.RFC3339, "2016-04-10T16:00:00+02:00")
	if e != nil {
		fmt.Println(e)
	}
	worker := Worker{1, "kalle", []DayRange{DayRange{1, 800, 1600}}}
	workers := []Worker{worker}
	shift := Shift{1, startsAt, endsAt, workers}
	assert.True(t, shift.isAssigned(worker))
}

func TestAssign(t *testing.T) {
	startsAt, e := time.Parse(time.RFC3339, "2016-04-10T12:00:00+02:00")
	endsAt, e := time.Parse(time.RFC3339, "2016-04-10T16:00:00+02:00")
	if e != nil {
		fmt.Println(e)
	}
	workers := make([]Worker, 1, 1) // []Worker{worker}
	shift := Shift{1, startsAt, endsAt, workers}

	worker := Worker{1, "kalle", []DayRange{DayRange{1, 800, 1600}}}
	shift.assign(worker)

	assert.Equal(t, shift.Workers[0], worker)
}
