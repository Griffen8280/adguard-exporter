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

  DnsQueries = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "num_dns_queries",
      Namespace: "adguard",
      Help:      "Number of DNS queries",
    },
    []string{"hostname"},
  )

  BlockedFiltering = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "num_blocked_filtering",
      Namespace: "adguard",
      Help:      "This represent the number of domains blocked",
    },
    []string{"hostname"},
  )

  ParentalFiltering = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "num_replaced_parental",
      Namespace: "adguard",
      Help:      "This represent the number of domains blocked (parental)",
    },
    []string{"hostname"},
  )

  SafeBrowsingFiltering = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "num_replaced_safebrowsing",
      Namespace: "adguard",
      Help:      "This represent the number of domains blocked (safe browsing)",
    },
    []string{"hostname"},
  )

  SafeSearchFiltering = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "num_replaced_safesearch",
      Namespace: "adguard",
      Help:      "This represent the number of domains blocked (safe search)",
    },
    []string{"hostname"},
  )

  TopQueries = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "top_queried_domains",
      Namespace: "adguard",
      Help:      "This represent the top queried domains",
    },
    []string{"hostname", "domain"},
  )

  TopBlocked = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "top_blocked_domains",
      Namespace: "adguard",
      Help:      "This represent the top bloacked domains",
    },
    []string{"hostname", "domain"},
  )

  TopClients = prometheus.NewGaugeVec(
    prometheus.GaugeOpts{
      Name:      "top_clients",
      Namespace: "adguard",
      Help:      "This represent the top clients",
    },
    []string{"hostname", "client"},
  )
)

// Init initializes all Prometheus metrics made available by AdGuard  exporter.
func Init() {
  initMetric("avg_processing_time", AvgProcessingTime)
  initMetric("num_dns_queries", DnsQueries)
  initMetric("num_blocked_filtering", BlockedFiltering)
  initMetric("num_replaced_parental", ParentalFiltering)
  initMetric("num_replaced_safebrowsing", SafeBrowsingFiltering)
  initMetric("num_replaced_safesearch", SafeSearchFiltering)
  initMetric("top_queried_domains", TopQueries)
  initMetric("top_blocked_domains", TopBlocked)
  initMetric("top_clients", TopClients)
}

func initMetric(name string, metric *prometheus.GaugeVec) {
  prometheus.MustRegister(metric)
  log.Printf("New Prometheus metric registered: %s", name)
}
