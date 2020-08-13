package redmine

import (
	"fmt"
	"strings"
)

// Status represents the issue status in Redmine.
type Status struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (status Status) String() string {
	return fmt.Sprintf("Status : %s", status.Name)
}

// Convert removes whitespace from status name.
func (status Status) Convert() string {
	return strings.Replace(status.Name, " ", "", 1)
}
