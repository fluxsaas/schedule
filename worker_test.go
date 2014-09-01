package schedule

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIfWorkerIsWorker(t *testing.T) {
	day := DayRange{1, 800, 1600}
	worker := Worker{1, "kalle", []DayRange{day}}
	assert.Equal(t, worker.Id, 1)
	assert.Equal(t, worker.Name, "kalle")
	assert.Equal(t, worker.Available[0], day)
}

func TestWorkerShouldBeAvailableOnDayRange(t *testing.T) {
	day := DayRange{1, 800, 1600}
	worker := Worker{1, "kalle", []DayRange{day}}
	assert.True(t, worker.isAvailable(day))
}

func TestWorkerShouldNotBeAvailableOnDayRange(t *testing.T) {
	day := DayRange{1, 800, 1600}
	day2 := DayRange{1, 800, 1700}
	worker := Worker{1, "kalle", []DayRange{day}}
	assert.Equal(t, worker.isAvailable(day2), false)
}
