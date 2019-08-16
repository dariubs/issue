package base

import "time"

// Milestone type
type Milestone struct {
	Title       string
	DueDate     time.Time
	Description string

	User User
	Time time.Time
}
