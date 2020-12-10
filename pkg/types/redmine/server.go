package redmine

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Server represents a instance on an Redmine server.
type Server struct {
	URL   string
	Token string
	Limit int
}

func (server *Server) String() string {
	return fmt.Sprintf("Redmine server: %v", server.URL)
}

// Configure reads the configuration details from map and sets the server instance.
func (server *Server) Configure(details map[string]interface{}) error {
	url, ok := details["domain"]
	if ok {
		serverURL, correctType := url.(string)
		if correctType {
			server.URL = serverURL
		}
	} else {
		server.URL = "https://redmine.example.com"
	}
	token, ok := details["api"]
	if ok {
		serverToken, correctType := token.(string)
		if correctType {
			server.Token = serverToken
		}
	}
	limit, ok := details["limit"]
	if ok {
		serverLimit, correctType := limit.(int)
		if correctType {
			server.Limit = serverLimit
		}
	} else {
		server.Limit = 120
	}
	return nil
}

// LoadIssues gets issues from Redmine server.
func (server *Server) LoadIssues() error {
	fmt.Printf("Redmine load issues from %s\n", server.URL)
	issuesURL := fmt.Sprintf("%s/issues.json?limit=%d", server.URL, server.Limit)
	request, err := http.NewRequest("GET", issuesURL, nil)
	if err != nil {
		fmt.Printf("creation of request failed: %v\n", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("X-Redmine-API-Key", server.Token)
	//fmt.Printf("used request:\n%v\n", request)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	}
	defer response.Body.Close()
	var issueList IssueList
	if err := json.NewDecoder(response.Body).Decode(&issueList); err != nil {
		fmt.Printf("Response could not parsed as JSON - %v\n", err)
		return err
	}
	fmt.Printf("Issue List:\n%v\n", issueList)
	return nil
}

// LoadProjects gets projects from Redmine server.
func (server *Server) LoadProjects() error {
	fmt.Printf("Gitlab load projects from %s\n", server.URL)
	projectsURL := fmt.Sprintf("%s/%s/projects?per_page=%d", server.URL, "api/v4", 120)
	request, err := http.NewRequest("GET", projectsURL, nil)
	if err != nil {
		fmt.Printf("creation of request failed: %v\n", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("PRIVATE-TOKEN", server.Token)
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	}
	defer response.Body.Close()
	var projectList []Project
	if err := json.NewDecoder(response.Body).Decode(&projectList); err != nil {
		fmt.Printf("Response could not parsed as JSON - %v\n", err)
		return err
	}
	fmt.Printf("Project List:\n%v\n", projectList)
	return nil
}
