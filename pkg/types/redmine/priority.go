package redmine

import "fmt"

// Priority represents priority of issue in Redmine.
type Priority struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (priority Priority) String() string {
	return fmt.Sprintf("Priority: %s", priority.Name)
}
