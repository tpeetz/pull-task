package redmine

import "fmt"

// CustomField represents a custom field of issue in Redmine.
type CustomField struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}

func (customField *CustomField) String() string {
	return fmt.Sprintf("CustomField: %d=%s", customField.ID, customField.Name)
}
