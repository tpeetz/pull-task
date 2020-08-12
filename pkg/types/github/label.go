package github

import "fmt"

// Label represents issue label.
type Label struct {
	ID          int    `json:"id"`
	NodeID      string `json:"node_id"`
	URL         string `json:"url"`
	Name        string `json:"name"`
	Color       string `json:"color"`
	DEfault     bool   `json:"default"`
	Description string `json:"description"`
}

func (label Label) String() string {
	return fmt.Sprintf("License %s\n", label.Name)
}
