package server

import "fmt"

// Configuration defines the placeholder for storing server configuration objects.
type Configuration struct {
	list []GitServer
}

// Add GitServer instance to Server Configuration
func (configuration *Configuration) Add(server GitServer) {
	//fmt.Printf("Add server to list: %v\n", server)
	configuration.list = append(configuration.list, server)
	//fmt.Printf("new length of list: %v", len(configuration.list))
}

// List returns the list of configured servers
func (configuration *Configuration) List() []GitServer {
	return configuration.list
}

// String returns stringified list of configured GitServer.
func (configuration *Configuration) String() string {
	result := ""
	for server := range configuration.list {
		result += fmt.Sprintf("%v\n", server)
	}
	return result
}
