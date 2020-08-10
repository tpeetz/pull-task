package gitlab

import "fmt"

// User represents an user account in Gitlab.
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	State    string `json:"state"`
	Avatar   string `json:"avatar_url"`
	WebURL   string `json:"web_url"`
}

func (user User) String() string {
	return fmt.Sprintf("%s (%s)", user.Name, user.Username)
}
