package adguard

import "fmt"

const (
	enabledStatus = "enabled"
)

type Stats struct {
	AvgProcessingTime    float64            `json:"avg_processing_time"`
}
