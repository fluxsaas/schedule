package schedule

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestPopulateSchedule(t *testing.T) {
	// Tuesday 02.09.2014
	b := []byte(`{"id":1,"starts_at":"2014-09-02T12:30:00+02:00","ends_at":"2014-09-02T16:30:00+02:00"}`)
	var shift Shift
	json.Unmarshal(b, &shift)
	shifts := []Shift{shift}

	a := []byte(`{"id":1,"name":"kalle","available":[{"wday": 2, "starts_at": 1200, "ends_at": 1640}]}`)
	var worker Worker
	json.Unmarshal(a, &worker)
	workers := []Worker{worker}

	// Call Populate function
	populated_shifts, err := Populate(shifts, workers)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---populated_shifts")
	fmt.Println(populated_shifts)
}
