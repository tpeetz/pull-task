package github

import "fmt"

// User represents an user in Github.
type User struct {
	ID                int    `json:"id"`
	NodeID            string `json:"node_id"`
	Login             string `json:"login"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
}

func (user User) String() string {
	return fmt.Sprintf("User %s\n", user.Login)
}
