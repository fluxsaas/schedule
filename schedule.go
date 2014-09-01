// The main schedule gives you methods to combine and build
// a schedule from workers and shifts.
// This first version implements following logic:
// 1. collect all day ranges in schedule (aka `shifts`)
// 2. iterate over each shift and find an first available workers.
// Enhanced version should include more options on worker:
// Departments, WorkingHours, Rotating, Preferences, Blocking days,
// Maybe the is a more generic approach on restriction...maybe `rule sets` ?
package main

func main() {
}
