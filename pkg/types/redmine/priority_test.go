package redmine

import (
	"reflect"
	"testing"
)

func TestCheckStructPriority(t *testing.T) {
	var redmineAttributeTestTable = []struct {
		name     string
		typeName string
	}{
		{"ID", "int"},
		{"Name", "string"},
	}
	p := Priority{}
	for index, testData := range redmineAttributeTestTable {
		givenType := reflect.TypeOf(p).Field(index).Type.Kind().String()
		fieldName := reflect.TypeOf(p).Field(index).Name
		if givenType != testData.typeName {
			t.Errorf("The expected type '%s' of field '%v', does not match type '%s'.", testData.typeName, fieldName, givenType)
		}
	}
}

func TestPriority_String(t *testing.T) {
	tests := []struct {
		name     string
		priority Priority
		want     string
	}{
		{"Priorität Niedrig", Priority{1, "Niedrig"}, "Priority: Niedrig"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.priority.String(); got != tt.want {
				t.Errorf("Priority.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
