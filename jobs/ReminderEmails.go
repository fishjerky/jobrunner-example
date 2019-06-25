package jobs

import "fmt"

// Job Specific Functions
type ReminderEmails struct {
	// filtered
}

// ReminderEmails.Run() will get triggered automatically.
func (e ReminderEmails) Run() {
	// Queries the DB
	// Sends some email

	fmt.Printf("[ReminderEmails]Every 5 sec send reminder emails \n")
}
