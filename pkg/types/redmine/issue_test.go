package redmine

import (
	"reflect"
	"testing"
)

func TestCheckStructIssue(t *testing.T) {
	var redmineAttributeTestTable = []struct {
		name     string
		typeName string
	}{
		{"ID", "int"},
		{"Project", "struct"},
		{"Tracker", "struct"},
		{"Status", "struct"},
		{"Priority", "struct"},
		{"Author", "struct"},
		{"Assigned", "struct"},
		{"Parent", "struct"},
		{"Subject", "string"},
		{"Description", "string"},
		{"StartDate", "string"},
		{"DueDate", "string"},
		{"DoneRatio", "int"},
		{"Estimated", "string"},
		{"Created", "string"},
		{"Updated", "string"},
		{"CustomFields", "slice"},
	}
	d := Issue{}
	for index, testData := range redmineAttributeTestTable {
		givenType := reflect.TypeOf(d).Field(index).Type.Kind().String()
		fieldName := reflect.TypeOf(d).Field(index).Name
		if givenType != testData.typeName {
			t.Errorf("The expected type '%s' of field '%v', does not match type '%s'.", testData.typeName, fieldName, givenType)
		}
	}
}

func TestIssue_String(t *testing.T) {
	tests := []struct {
		name  string
		issue Issue
		want  string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.issue.String(); got != tt.want {
				t.Errorf("Issue.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
