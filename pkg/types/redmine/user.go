package redmine

import "fmt"

// User represents an user in Redmine.
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (user *User) String() string {
	return fmt.Sprintf("User: %d=%s", user.ID, user.Name)
}
