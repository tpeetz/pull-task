package server

// GitServer defines the general structure of a Git server.
type GitServer struct {
	url   string
	token string
}

// LoadIssues requests list of issues from GitServer instance.
func (gitServer *GitServer) LoadIssues() {

}

// LoadProjects requests list of projects from GitServer instance.
func (gitServer *GitServer) LoadProjects() {

}

// NewGitlabServer creates an instance of a Gitlab server.
func NewGitlabServer() *GitServer {
	gitServer := GitServer{}
	gitServer.url = "https://gitlab.com"
	gitServer.token = "secretToken"
	return &gitServer
}

// NewGithubServer creates an instance of a Github server.
func NewGithubServer() *GitServer {
	gitServer := GitServer{}
	gitServer.url = "https://github.com"
	gitServer.token = "secretToken"
	return &gitServer
}

// NewRedmineServer creates an instance of a Redmine server.
func NewRedmineServer() *GitServer {
	gitServer := GitServer{}
	gitServer.url = "https://redmine.com"
	gitServer.token = "secretToken"
	return &gitServer
}
