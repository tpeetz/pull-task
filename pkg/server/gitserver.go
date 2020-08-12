package server

import (
	"errors"

	"github.com/tpeetz/pull-task/pkg/types/gitlab"
)

// GitServer definees the methods for server (Gitlab, Github, Redmine)
type GitServer interface {
	Configure(map[string]interface{}) error
	LoadIssues() error
	LoadProjects() error
	String() string
}

// NewGitServer creates instance of GitServer from given configuration.
func NewGitServer(details map[string]interface{}) (GitServer, error) {
	//fmt.Printf("details: %v\n", details)
	var server GitServer
	service := details["service"]
	switch service {
	case "gitlab":
		server = &gitlab.Server{}
	case "github":
		return nil, errors.New("service type github unknown")
	case "redmine":
		return nil, errors.New("service type redmine unknown")
	}
	server.Configure(details)
	return server, nil
}
