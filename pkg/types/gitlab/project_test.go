package gitlab

import (
	"fmt"
	"reflect"
	"testing"
)

var gitlabProjectTestTable = []struct {
	name     string
	typeName string
}{
	{"ID", "int"},
	{"Name", "string"},
	{"Path", "string"},
	{"PathNamespace", "string"},
	{"Description", "string"},
	{"Archived", "bool"},
	{"Visibility", "string"},
	{"Owner", "struct"},
}

func TestCheckGitlabStructProject(t *testing.T) {
	d := Project{}
	for index, testData := range gitlabProjectTestTable {
		givenType := reflect.TypeOf(d).Field(index).Type.Kind().String()
		if givenType != testData.typeName {
			t.Errorf("The expected type '%s', does not match type '%s'.", testData.typeName, givenType)
		}
	}
}

func TestProjectString(t *testing.T) {
	project := Project{}
	project.ID = 2
	project.Name = "Testfall"
	expected := fmt.Sprintf("Project (%d): %s\n  Path with namespace: %s\n  Description: %s\n  Owner: %s", project.ID, project.Name, project.PathNamespace, project.Description, project.Owner)
	result := project.String()
	if expected != result {
		t.Errorf("The expected string '%s', does not match result '%s'.", expected, result)
	}
}

func TestProject_Short(t *testing.T) {
	type fields struct {
		ID            int
		Name          string
		Path          string
		PathNamespace string
		Description   string
		Archived      bool
		Visibility    string
		Owner         User
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Testfall", fields{1, "Testfall", "testfall", "testfall", "", false, "public", User{1, "user", "username", "active", "avatar", "weburl"}}, "Project: Testfall"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Project{
				ID:            tt.fields.ID,
				Name:          tt.fields.Name,
				Path:          tt.fields.Path,
				PathNamespace: tt.fields.PathNamespace,
				Description:   tt.fields.Description,
				Archived:      tt.fields.Archived,
				Visibility:    tt.fields.Visibility,
				Owner:         tt.fields.Owner,
			}
			if got := r.Short(); got != tt.want {
				t.Errorf("Project.Short() = %v, want %v", got, tt.want)
			}
		})
	}
}
