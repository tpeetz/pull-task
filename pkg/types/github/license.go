package github

import "fmt"

// License holds information about repository license.
type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxID string `json:"spdx_id"`
	URL    string `json:"url"`
	NodeID string `json:"node_id"`
}

func (license License) String() string {
	return fmt.Sprintf("License %s\n", license.Name)
}
