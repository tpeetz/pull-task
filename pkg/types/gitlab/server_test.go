package gitlab

import (
	"testing"
)

func TestServer_String(t *testing.T) {
	type fields struct {
		URL   string
		Token string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"Default Gitlab server", fields{URL: "https://gitlab.com", Token: "secretToken"}, "Gitlab server: https://gitlab.com"},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &Server{
				URL:   tt.fields.URL,
				Token: tt.fields.Token,
			}
			if got := server.String(); got != tt.want {
				t.Errorf("Server.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestServer_Configure(t *testing.T) {
	type fields struct {
		URL   string
		Token string
	}
	type args struct {
		details map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &Server{
				URL:   tt.fields.URL,
				Token: tt.fields.Token,
			}
			if err := server.Configure(tt.args.details); (err != nil) != tt.wantErr {
				t.Errorf("Server.Configure() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_LoadIssues(t *testing.T) {
	type fields struct {
		URL   string
		Token string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &Server{
				URL:   tt.fields.URL,
				Token: tt.fields.Token,
			}
			if err := server.LoadIssues(); (err != nil) != tt.wantErr {
				t.Errorf("Server.LoadIssues() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestServer_LoadProjects(t *testing.T) {
	type fields struct {
		URL   string
		Token string
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &Server{
				URL:   tt.fields.URL,
				Token: tt.fields.Token,
			}
			if err := server.LoadProjects(); (err != nil) != tt.wantErr {
				t.Errorf("Server.LoadProjects() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
