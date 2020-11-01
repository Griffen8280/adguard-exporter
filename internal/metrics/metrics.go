package metrics

import (
	"log"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	AvgProcessingTime = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name:      "avg_processing_time",
			Namespace: "adguard",
			Help:      "This represent the average processing time for a DNS query in s",
		},
		[]string{"hostname"},
	)
)

// Init initializes all Prometheus metrics made available by AdGuard  exporter.
func Init() {
	initMetric("avg_processing_time", AvgProcessingTime)
}

func initMetric(name string, metric *prometheus.GaugeVec) {
	prometheus.MustRegister(metric)
	log.Printf("New Prometheus metric registered: %s", name)
}
