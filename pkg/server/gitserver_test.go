package server

import (
	"reflect"
	"testing"

	"github.com/tpeetz/pull-task/pkg/types/github"
	"github.com/tpeetz/pull-task/pkg/types/gitlab"
	"github.com/tpeetz/pull-task/pkg/types/redmine"
)

var (
	gitlabConfig = map[string]interface{}{
		"service": "gitlab",
	}
	glServer     = &gitlab.Server{URL: "https://gitlab.com"}
	githubConfig = map[string]interface{}{
		"service": "github",
	}
	ghServer      = &github.Server{URL: "https://api.github.com"}
	redmineConfig = map[string]interface{}{
		"service": "redmine",
		"limit":   120,
	}
	rmServer      = &redmine.Server{URL: "https://redmine.example.com"}
	unknownConfig = map[string]interface{}{
		"service": "unknown",
	}
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
		{"github", args{githubConfig}, ghServer, false},
		{"redmine", args{redmineConfig}, rmServer, false},
		{"unknown", args{unknownConfig}, nil, true},
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
