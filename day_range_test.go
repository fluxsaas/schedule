package schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfDayRangeIsDayRange(t *testing.T) {
	test := DayRange{1, 800, 1600}
	assert.Equal(t, test.Wday, 1)
	assert.Equal(t, test.StartsAt, 800)
	assert.Equal(t, test.EndsAt, 1600)
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
