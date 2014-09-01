package schedule

// Worker struct represents a specific worker with preferences and skillset
type Worker struct {
	Id        int        `json:"id"`
	Name      string     `json:"name"`
	Available []DayRange `json:"available"`
}

// isAvailable checks whether or not a worker is Available on a specific day.
// We do not use pointer (*Worker) because this method does not mutate the
// receiving struct.
func (worker Worker) isAvailable(dayRange DayRange) bool {
	return dayRange.withinMany(worker.Available)
}

// isAvailableForShift checks whether or not a worker is Available for a shift.
func (worker Worker) isAvailableForShift(shift Shift) bool {
	wday := int(shift.StartsAt.Weekday())
	startsAt := shift.StartsAt.Hour()*100 + shift.StartsAt.Minute()
	endsAt := shift.EndsAt.Hour()*100 + shift.EndsAt.Minute()
	dayRange := DayRange{wday, startsAt, endsAt}
	return shift.isAssigned(worker) && dayRange.withinMany(worker.Available)
}
