package base

import "time"

// Issue type
type Issue struct {
	Title       string
	Description string

	Labels    []Label
	Milestone Milestone

	User User
	Time time.Time

	Comments []Comment
}
