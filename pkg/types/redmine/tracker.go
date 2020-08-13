package redmine

import "fmt"

// Tracker represents a issue type in Redmine.
type Tracker struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (tracker *Tracker) String() string {
	return fmt.Sprintf("Tracker %s", tracker.Name)
}
