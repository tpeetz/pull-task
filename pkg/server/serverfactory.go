package server

import "fmt"

// IssueLoader defines the interface for requesting issues.
type IssueLoader interface {
	LoadIssues() string
}

// ProjectLoader defines the interface for requesting issues.
type ProjectLoader interface {
	LoadProjects() string
}

// NewGitServer creates instance of GitServer from given configuration.
func NewGitServer(config interface{}) *GitServer {
	fmt.Printf("config: %v", config)
	return nil
}
