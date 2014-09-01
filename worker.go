package schedule

// Worker struct represents a specific worker with preferences and skillset
type Worker struct {
	Id        int
	Name      string
	Available []DayRange
}

// isAvailable checks whether or not a worker is available on a specific day.
// We do not use pointer (*Worker) because this method does not mutate the
// receiving struct.
func (worker Worker) isAvailable(day DayRange) bool {
	return day.withinMany(worker.Available)
}
