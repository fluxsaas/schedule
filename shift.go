package schedule

import "time"

// Shift struct represents a time ranges in a schedule
// including one or many `assigned` workers.
type Shift struct {
	Id       int       `json:"id"`        // uniq Id
	StartsAt time.Time `json:"starts_at"` // start date
	EndsAt   time.Time `json:"ends_at"`   // end date
	Workers  []Worker  `json:"workers"`   // assigned worker for this shift
}

// isAssigned checks if the worker is already assigned to this shift
func (shift Shift) isAssigned(worker Worker) bool {
	return false
}

// isAssigned checks if the worker is already assigned to this shift
func (shift Shift) assign(worker Worker) Shift {
	return shift
}
