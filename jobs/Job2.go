package jobs

import "fmt"

// Job Specific Functions
type Job2 struct {
	Message string
}

// ReminderEmails.Run() will get triggered automatically.
func (e Job2) Run() {
	// Queries the DB
	// Sends some email

	fmt.Printf("[Job 2] %s \n", e.Message)
}
