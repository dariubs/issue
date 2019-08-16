package base

import "time"

// Label type
type Label struct {
	Name        string
	Description string
	Color       string

	User User
	Time time.Time
}
