package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Server represents an instance of a Github server.
type Server struct {
	URL   string
	Token string
}

func (server *Server) String() string {
	return fmt.Sprintf("Github server: %v", server.URL)
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
		server.URL = "https://api.github.com"
	}
	token, ok := details["token"]
	if ok {
		serverToken, correctType := token.(string)
		if correctType {
			server.Token = serverToken
		}
	}
	return nil
}

// LoadIssues gets issues from Gitlab server.
func (server *Server) LoadIssues() error {
	fmt.Printf("Github load issues from %s\n", server.URL)
	issuesURL := fmt.Sprintf("%s/issues", server.URL)
	request, err := http.NewRequest("GET", issuesURL, nil)
	if err != nil {
		fmt.Printf("creation of request failed: %v\n", err)
		return err
	}
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", fmt.Sprintf("token %s", server.Token))
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return err
	}
	defer response.Body.Close()
	var issueList []Issue
	if err := json.NewDecoder(response.Body).Decode(&issueList); err != nil {
		fmt.Printf("Response could not parsed as JSON - %v\n", err)
		return err
	}
	fmt.Printf("Issue List:\n%v\n", issueList)
	return nil
}

// LoadProjects gets projects from Gitlab server.
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
	var projectList []interface{}
	if err := json.NewDecoder(response.Body).Decode(&projectList); err != nil {
		fmt.Printf("Response could not parsed as JSON - %v\n", err)
		return err
	}
	fmt.Printf("Project List:\n%v\n", projectList)
	return nil
}
