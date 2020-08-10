package gitlab

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCheckStructIssue(t *testing.T) {
	var taskwarriorAttributeTestTable = []struct {
		name     string
		typeName string
	}{
		{"ID", "int"},
		{"InternalID", "int"},
		{"ProjectID", "int"},
		{"Title", "string"},
		{"Description", "string"},
		{"State", "string"},
		{"Labels", "slice"},
		{"Created", "string"},
		{"Updated", "string"},
		{"Closed", "string"},
		{"ClosedBy", "struct"},
		{"Author", "struct"},
		{"Assignee", "struct"},
		{"DueDate", "string"},
		{"WebURL", "string"},
		{"TimeStats", "struct"},
	}
	d := Issue{}
	for index, testData := range taskwarriorAttributeTestTable {
		givenType := reflect.TypeOf(d).Field(index).Type.Kind().String()
		if givenType != testData.typeName {
			t.Errorf("The expected type '%s', does not match type '%s'.", testData.typeName, givenType)
		}
	}
}

func TestIssue_String(t *testing.T) {
	type fields struct {
		ID          int
		InternalID  int
		ProjectID   int
		Title       string
		Description string
		State       string
		Labels      []string
		Created     string
		Updated     string
		Closed      string
		ClosedBy    User
		Author      User
		Assignee    User
		DueDate     string
		WebURL      string
		TimeStats   TimeStats
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Empty issue", fields{}, fmt.Sprintf("Issue %d: %s\n  Project:  %d\n  Start  : %s\n  Due    : %s\n  %s\n  []\n  %s\n", 0, "", 0, "", "", "", "")},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Issue{
				ID:          tt.fields.ID,
				InternalID:  tt.fields.InternalID,
				ProjectID:   tt.fields.ProjectID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				State:       tt.fields.State,
				Labels:      tt.fields.Labels,
				Created:     tt.fields.Created,
				Updated:     tt.fields.Updated,
				Closed:      tt.fields.Closed,
				ClosedBy:    tt.fields.ClosedBy,
				Author:      tt.fields.Author,
				Assignee:    tt.fields.Assignee,
				DueDate:     tt.fields.DueDate,
				WebURL:      tt.fields.WebURL,
				TimeStats:   tt.fields.TimeStats,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Issue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIssue_Short(t *testing.T) {
	type fields struct {
		ID          int
		InternalID  int
		ProjectID   int
		Title       string
		Description string
		State       string
		Labels      []string
		Created     string
		Updated     string
		Closed      string
		ClosedBy    User
		Author      User
		Assignee    User
		DueDate     string
		WebURL      string
		TimeStats   TimeStats
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Empty issue", fields{}, "Issue 0: "},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Issue{
				ID:          tt.fields.ID,
				InternalID:  tt.fields.InternalID,
				ProjectID:   tt.fields.ProjectID,
				Title:       tt.fields.Title,
				Description: tt.fields.Description,
				State:       tt.fields.State,
				Labels:      tt.fields.Labels,
				Created:     tt.fields.Created,
				Updated:     tt.fields.Updated,
				Closed:      tt.fields.Closed,
				ClosedBy:    tt.fields.ClosedBy,
				Author:      tt.fields.Author,
				Assignee:    tt.fields.Assignee,
				DueDate:     tt.fields.DueDate,
				WebURL:      tt.fields.WebURL,
				TimeStats:   tt.fields.TimeStats,
			}
			if got := r.Short(); got != tt.want {
				t.Errorf("Issue.Short() = %v, want %v", got, tt.want)
			}
		})
	}
}
