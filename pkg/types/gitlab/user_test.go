package gitlab

import (
	"reflect"
	"testing"
)

var gitlabSimpleDataTestTable = []struct {
	name     string
	typeName string
}{
	{"ID", "int"},
	{"Name", "string"},
}

func TestCheckGitlabStructUser(t *testing.T) {
	d := User{}
	for index, testData := range gitlabSimpleDataTestTable {
		givenType := reflect.TypeOf(d).Field(index).Type.Kind().String()
		if givenType != testData.typeName {
			t.Errorf("The expected type '%s', does not match type '%s'.", testData.typeName, givenType)
		}
	}
}

func TestUser_String(t *testing.T) {
	type fields struct {
		ID       int
		Name     string
		Username string
		State    string
		Avatar   string
		WebURL   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Administrator account", fields{1, "Administrator", "admin", "", "", ""}, "Administrator (admin)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user := User{
				ID:       tt.fields.ID,
				Name:     tt.fields.Name,
				Username: tt.fields.Username,
				State:    tt.fields.State,
				Avatar:   tt.fields.Avatar,
				WebURL:   tt.fields.WebURL,
			}
			if got := user.String(); got != tt.want {
				t.Errorf("User.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
