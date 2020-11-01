package adguard

import "fmt"

const (
	enabledStatus = "enabled"
)

type Stats struct {
	AvgProcessingTime    float64            `json:"avg_processing_time"`
}

// ToString method returns a string of the current statistics struct.
func (s *Stats) ToString() string {
	return fmt.Sprintf("Average Processing Time %d s", s.AvgProcessingTime)
}
