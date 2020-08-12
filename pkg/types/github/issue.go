package github

import "fmt"

// Issue represents issue in Github.
type Issue struct {
	ID                    int        `json:"id"`
	NodeID                string     `json:"node_id"`
	Title                 string     `json:"title"`
	Number                int        `json:"number"`
	User                  User       `json:"user"`
	Labels                []Label    `json:"labels"`
	URL                   string     `json:"url"`
	RepositoryURL         string     `json:"repository_url"`
	LabelsURL             string     `json:"labels_url"`
	CommentsURL           string     `json:"comments_url"`
	EventsURL             string     `json:"events_url"`
	HTMLURL               string     `json:"html_url"`
	State                 string     `json:"state"`
	Locked                bool       `json:"locked"`
	Assignee              User       `json:"assignee"`
	Assignees             []User     `json:"assignees"`
	Milestone             Milestone  `json:"milestone,omitempty"`
	Comments              int        `json:"comments"`
	Created               string     `json:"created_at,omitempty"`
	Updated               string     `json:"updated_at,omitempty"`
	Closed                string     `json:"closed_at,omitempty"`
	AuthorAssociation     string     `json:"author_association"`
	ActiveLockReason      string     `json:"active_lock_reason,omitempty"`
	Repository            Repository `json:"repository"`
	Body                  string     `json:"body"`
	PerformedViaGithubApp string     `json:"performed_via_github_app,omitempty"`
}

func (issue Issue) String() string {
	//return fmt.Sprintf("Issue %d: %s\n  Project:  %d\n  Start  : %s\n  Due    : %s\n  %s\n  %v\n  %s\n", r.ID, r.Title, r.ProjectID, r.Created, r.DueDate, r.State, r.Labels, r.WebURL)
	return fmt.Sprintf("Issue %d: %s\n", issue.ID, issue.Title)
}

// Short returns issue details in short.
func (issue Issue) Short() string {
	return fmt.Sprintf("Issue %d: %s", issue.ID, issue.Title)
}
