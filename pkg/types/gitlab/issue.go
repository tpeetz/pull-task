package gitlab

import (
	"fmt"
)

// Issue represents issues in Gitlab.
type Issue struct {
	ID          int       `json:"id"`
	InternalID  int       `json:"iid"`
	ProjectID   int       `json:"project_id"`
	Title       string    `json:"title"`
	Description string    `json:"description,omitempty"`
	State       string    `json:"state"`
	Labels      []string  `json:"labels,omitempty"`
	Created     string    `json:"created_at,omitempty"`
	Updated     string    `json:"updated_at,omitempty"`
	Closed      string    `json:"closed_at,omitempty"`
	ClosedBy    User      `json:"closed_by,omitempty"`
	Author      User      `json:"author"`
	Assignee    User      `json:"assignee"`
	DueDate     string    `json:"due_date,omitempty"`
	WebURL      string    `json:"web_url"`
	TimeStats   TimeStats `json:"time_stats"`
}

func (r Issue) String() string {
	return fmt.Sprintf("Issue %d: %s\n  Project:  %d\n  Start  : %s\n  Due    : %s\n  %s\n  %v\n  %s\n", r.ID, r.Title, r.ProjectID, r.Created, r.DueDate, r.State, r.Labels, r.WebURL)
}

// Short returns issue details in short.
func (r Issue) Short() string {
	return fmt.Sprintf("Issue %d: %s", r.ID, r.Title)
}
