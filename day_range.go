package schedule

// DayRange struct represents a time ranges on a wday
type DayRange struct {
	Wday     int // 0 - 6
	StartsAt int // 0 - 2359
	EndsAt   int // 0 - 2359
}

// within checks if day range is within other day range
func (day DayRange) within(other DayRange) bool {
	if day.Wday == other.Wday && day.StartsAt >= other.StartsAt && day.EndsAt <= other.EndsAt {
		return true
	}
	return false
}

// withinMany checks if one day range is within many day ranges
func (day DayRange) withinMany(days []DayRange) bool {
	for _, other := range days {
		if day.within(other) {
			return true
		}
	}
	return false
}
