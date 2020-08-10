package gitlab

import (
	"reflect"
	"testing"
)

var gitlabTimestatsTestTable = []struct {
	name     string
	typeName string
}{
	{"Estimate", "int"},
	{"TimeSpent", "int"},
}

func TestCheckGitlabStructTimeStats(t *testing.T) {
	d := TimeStats{}
	for index, testData := range gitlabTimestatsTestTable {
		givenType := reflect.TypeOf(d).Field(index).Type.Kind().String()
		if givenType != testData.typeName {
			t.Errorf("The expected type '%s', does not match type '%s'.", testData.typeName, givenType)
		}
	}
}

func TestTimeStats_String(t *testing.T) {
	type fields struct {
		Estimate          int
		TimeSpent         int
		HumanTimeEstimate string
		HumanTimeSpent    string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Estimates 3 hours", fields{3, 0, "", ""}, "Estimate: 3"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			timeStats := TimeStats{
				Estimate:          tt.fields.Estimate,
				TimeSpent:         tt.fields.TimeSpent,
				HumanTimeEstimate: tt.fields.HumanTimeEstimate,
				HumanTimeSpent:    tt.fields.HumanTimeSpent,
			}
			if got := timeStats.String(); got != tt.want {
				t.Errorf("TimeStats.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
