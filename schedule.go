package schedule

// Populate calcuates and assignes workers to shifts and
// returns the produced schedule
func Populate(shifts []Shift, workers []Worker) (n []Shift, err error) {
	for _, shift := range shifts {
		for _, worker := range workers {
			if worker.isAvailableForShift(shift) {
				shift.assign(worker)
			}
		}
	}
	return shifts, nil
}
