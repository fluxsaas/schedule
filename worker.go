// The Worker represents a specific user with preferenzes
// and skillset
package main

type Worker struct {
	Id        int
	Name      string
	Available []DayRange
}

// check if a worker is available on a specific day
// this is a strict check. partial overlapping is ignored for now.
func (worker *Worker) IsAvailable(day DayRange) bool {
	return day.withinMany(worker.Available)
}
