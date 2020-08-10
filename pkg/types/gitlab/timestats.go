package gitlab

import "fmt"

// TimeStats represents time statitics in Gitlab.
type TimeStats struct {
	Estimate          int    `json:"time_estimate"`
	TimeSpent         int    `json:"total_time_spent"`
	HumanTimeEstimate string `json:"human_time_estimate,omitempty"`
	HumanTimeSpent    string `json:"human_time_spent,omitempty"`
}

func (timeStats TimeStats) String() string {
	return fmt.Sprintf("Estimate: %v", timeStats.Estimate)
}
