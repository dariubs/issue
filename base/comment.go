package base

import "time"

// Comment type
type Comment struct {
	Content string
	Status  string

	User User
	Time time.Time
}
