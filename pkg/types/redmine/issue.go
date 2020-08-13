package redmine

import (
	"encoding/json"
	"fmt"
)

// Issue represents issue in Redmine.
type Issue struct {
	ID           int           `json:"id"`
	Project      Project       `json:"project"`
	Tracker      Tracker       `json:"tracker"`
	Status       Status        `json:"status"`
	Priority     Priority      `json:"priority"`
	Author       User          `json:"author"`
	Assigned     User          `json:"assigned_to"`
	Parent       ParentIssue   `json:"parent"`
	Subject      string        `json:"subject"`
	Description  string        `json:"description,omitempty"`
	StartDate    string        `json:"start_date,omitempty"`
	DueDate      string        `json:"due_date,omitempty"`
	DoneRatio    int           `json:"done_ratio"`
	Estimated    json.Number   `json:"estimated_hours"`
	Created      string        `json:"created_on,omitempty"`
	Updated      string        `json:"updated_on,omitempty"`
	CustomFields []CustomField `json:"custom_fields,omitempty"`
}

func (issue Issue) String() string {
	return fmt.Sprintf("Issue %d: %s\n  %s\n  Start  : %s\n  Due    : %s\n  %s\n", issue.ID, issue.Subject, issue.Project, issue.StartDate, issue.DueDate, issue.Status)
}

// Short returns ID and subject of issue.
func (issue Issue) Short() string {
	return fmt.Sprintf("Issue %d: %s", issue.ID, issue.Subject)
}

// SingleIssue represents a JSON answer of a single issue in Redmine.
type SingleIssue struct {
	Issue Issue `json:"issue"`
}
