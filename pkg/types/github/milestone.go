package github

import "fmt"

// Milestone represents milestone of a Github project.
type Milestone struct {
	ID           int    `json:"id"`
	NodeID       string `json:"node_id"`
	Number       int    `json:"number"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Creator      User   `json:"creator"`
	URL          string `json:"url"`
	HTMLURL      string `json:"html_url"`
	LabelsURL    string `json:"labels_url"`
	OpenIssues   int    `json:"open_issues"`
	ClosedIssues int    `json:"closed_issues"`
	State        string `json:"state"`
	Created      string `json:"created_at,omitempty"`
	Updated      string `json:"updated_at,omitempty"`
	Due          string `json:"due_on,omitempty"`
	Closed       string `json:"closed_at,omitempty"`
}

func (milestone Milestone) String() string {
	return fmt.Sprintf("Issue %d: %s\n", milestone.ID, milestone.Title)
}
