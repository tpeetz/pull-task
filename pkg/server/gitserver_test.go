package server

import (
	"reflect"
	"testing"

	"github.com/tpeetz/pull-task/pkg/types/gitlab"
)

var (
	gitlabConfig = map[string]interface{}{
		"service": "gitlab",
	}
	glServer     = &gitlab.Server{URL: "https://gitlab.com"}
	githubConfig = map[string]interface{}{
		"service": "github",
	}
	//ghServer      = &GitServer{url: "https://github.com", token: "secretToken"}
	redmineConfig = map[string]interface{}{
		"service": "redmine",
	}
	//rmServer = &GitServer{url: "https://redmine.com", token: "secretToken"}
)

func TestNewGitServer(t *testing.T) {
	type args struct {
		details map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    GitServer
		wantErr bool
	}{
		{"gitlab", args{gitlabConfig}, glServer, false},
		{"github", args{githubConfig}, nil, true},
		{"redmine", args{githubConfig}, nil, true},
		{"unknown", args{githubConfig}, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewGitServer(tt.args.details)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewGitServer() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
