package redmine

import "fmt"

// ParentIssue holds the ID of the parent issue of the current one.
type ParentIssue struct {
	ID int `json:"id"`
}

func (parent *ParentIssue) String() string {
	return fmt.Sprintf("Parent Issue: %d", parent.ID)
}
