package redmine

import "fmt"

// Project represents project in Redmine.
type Project struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (project Project) String() string {
	return fmt.Sprintf("Project: %s", project.Name)
}
