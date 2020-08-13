package server

import (
	"github.com/tpeetz/pull-task/pkg/types/github"
	"github.com/tpeetz/pull-task/pkg/types/gitlab"
	"github.com/tpeetz/pull-task/pkg/types/redmine"
)

// GitServer definees the methods for server (Gitlab, Github, Redmine)
type GitServer interface {
	Configure(map[string]interface{}) error
	LoadIssues() error
	LoadProjects() error
	String() string
}

// UnkownServiceTypeError defines the error condition for unkown service type.
type UnkownServiceTypeError struct {
	Service string
}

func (uste *UnkownServiceTypeError) Error() string {
	return uste.Service + " not known"
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
		server = &github.Server{}
	case "redmine":
		server = &redmine.Server{}
	default:
		return nil, &UnkownServiceTypeError{service.(string)}
	}
	server.Configure(details)
	return server, nil
}
