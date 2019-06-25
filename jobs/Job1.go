package jobs

import "fmt"

// Job Specific Functions
type Job1 struct {
	Message string
}

// ReminderEmails.Run() will get triggered automatically.
func (e Job1) Run() {
	// Queries the DB
	// Sends some email

	fmt.Printf("[Job 1] %s \n", e.Message)
}
