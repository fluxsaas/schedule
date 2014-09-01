// The Worker represents a specific user with preferenzes
// and skillset
package schedule

type Worker struct {
	Id        int
	Name      string
	Available []DayRange
}

// Checks if a worker is available on a specific day. (strict)
// We do not use pointer (*Worker) because this method does not
// mutate the receiving struct
func (worker Worker) IsAvailable(day DayRange) bool {
	return day.withinMany(worker.Available)
}
