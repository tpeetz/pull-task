package server

import (
	"reflect"
	"testing"
)

var (
	gitlab   = make(map[string]string)
	glServer = &GitServer{url: "https://gitlab.com", token: "secretToken"}
)

func TestNewGitServer(t *testing.T) {
	gitlab["type"] = "Gitlab"
	type args struct {
		config interface{}
	}
	tests := []struct {
		name string
		args args
		want *GitServer
	}{
		{"", args{config: gitlab}, nil},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGitServer(tt.args.config); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGitServer() = %v, want %v", got, tt.want)
			}
		})
	}
}
