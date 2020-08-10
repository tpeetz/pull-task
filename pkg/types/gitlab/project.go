package gitlab

import (
	"fmt"
)

var (
	projectMap = make(map[int]Project, 0)
)

// Project represents a project in Gitlab.
type Project struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Path          string `json:"path"`
	PathNamespace string `json:"path_with_namespace"`
	Description   string `json:"description"`
	Archived      bool   `json:"archived"`
	Visibility    string `json:"visibility"`
	Owner         User   `json:"owner"`
}

func (r Project) String() string {
	return fmt.Sprintf("Project (%d): %s\n  Path with namespace: %s\n  Description: %s\n  Owner: %s", r.ID, r.Name, r.PathNamespace, r.Description, r.Owner)
}

// Short returns project details in short.
func (r Project) Short() string {
	return fmt.Sprintf("Project: %s", r.Name)
}
