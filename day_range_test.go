package schedule

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfDayRangeIsDayRange(t *testing.T) {
	test := DayRange{1, 800, 1600}
	assert.Equal(t, 1, test.Wday)
	assert.Equal(t, 800, test.StartsAt)
	assert.Equal(t, 1600, test.EndsAt)
}

func TestDayRangeJsonDecoding(t *testing.T) {
	b := []byte(`{"wday":1,"starts_at": 1200,"ends_at": 2000}`)
	var dayRange DayRange
	err := json.Unmarshal(b, &dayRange)
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, 1, dayRange.Wday)
	assert.Equal(t, 1200, dayRange.StartsAt)
	assert.Equal(t, 2000, dayRange.EndsAt)
}

func TestThatDayRangeIsInDayRanges(t *testing.T) {
	day := DayRange{1, 800, 1600}
	day1 := DayRange{2, 800, 1600}
	days := []DayRange{day, day1}
	assert.True(t, day.withinMany(days), "day should be in days")

	day2 := DayRange{2, 2000, 2300}
	assert.False(t, day2.withinMany(days), "day should not be in day if time does not match")

	day3 := DayRange{2, 800, 1200}
	assert.True(t, day3.withinMany(days))
}
