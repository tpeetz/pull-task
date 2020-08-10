package server

import (
	"testing"
)

func TestConfiguration_Add(t *testing.T) {
	type fields struct {
		list []GitServer
	}
	type args struct {
		server GitServer
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{"Add nil pointer", fields{}, args{nil}},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configuration := &Configuration{
				list: tt.fields.list,
			}
			configuration.Add(tt.args.server)
		})
	}
}

func TestConfiguration_String(t *testing.T) {
	type fields struct {
		list []GitServer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Empty list", fields{}, ""},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			configuration := &Configuration{
				list: tt.fields.list,
			}
			if got := configuration.String(); got != tt.want {
				t.Errorf("Configuration.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
