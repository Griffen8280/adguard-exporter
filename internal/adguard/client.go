package adguard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ebrianne/adguard-exporter/internal/metrics"
)

var (
	statsURLPattern    = "%s://%s:%d/control/stats"
	logstatsURLPattern = "%s://%s:%d/control/querylog"
	m                  map[string]int
)

// Client struct is a AdGuard  client to request an instance of a AdGuard  ad blocker.
type Client struct {
	httpClient  http.Client
	interval    time.Duration
	protocol    string
	hostname    string
	port        uint16
	b64password string
}

// NewClient method initializes a new AdGuard  client.
func NewClient(protocol, hostname string, port uint16, b64password string, interval time.Duration) *Client {
	if protocol != "http" {
		log.Printf("protocol %s is invalid. Must be http.", protocol)
		os.Exit(1)
	}

	return &Client{
		protocol:    protocol,
		hostname:    hostname,
		port:        port,
		b64password: b64password,
		interval:    interval,
		httpClient:  http.Client{},
	}
}

// Scrape method authenticates and retrieves statistics from AdGuard  JSON API
// and then pass them as Prometheus metrics.
func (c *Client) Scrape() {
	for range time.Tick(c.interval) {

		//Get the general stats
		stats := c.getStatistics()
		c.setMetrics(stats)

		//Get the log stats
		logdata := c.getLogStatistics()
		c.setLogMetrics(logdata)

		log.Printf("New tick of statistics: %s", stats.ToString())
	}
}

// Function to set the general stats
func (c *Client) setMetrics(stats *Stats) {
	metrics.AvgProcessingTime.WithLabelValues(c.hostname).Set(float64(stats.AvgProcessingTime))
	metrics.DnsQueries.WithLabelValues(c.hostname).Set(float64(stats.DnsQueries))
	metrics.BlockedFiltering.WithLabelValues(c.hostname).Set(float64(stats.BlockedFiltering))
	metrics.ParentalFiltering.WithLabelValues(c.hostname).Set(float64(stats.ParentalFiltering))
	metrics.SafeBrowsingFiltering.WithLabelValues(c.hostname).Set(float64(stats.SafeBrowsingFiltering))
	metrics.SafeSearchFiltering.WithLabelValues(c.hostname).Set(float64(stats.SafeSearchFiltering))

	for l := range stats.TopQueries {
		for domain, value := range stats.TopQueries[l] {
			metrics.TopQueries.WithLabelValues(c.hostname, domain).Set(float64(value))
		}
	}

	for l := range stats.TopBlocked {
		for domain, value := range stats.TopBlocked[l] {
			metrics.TopBlocked.WithLabelValues(c.hostname, domain).Set(float64(value))
		}
	}

	for l := range stats.TopClients {
		for source, value := range stats.TopClients[l] {
			metrics.TopClients.WithLabelValues(c.hostname, source).Set(float64(value))
		}
	}
}

// Function to get the general stats
func (c *Client) getStatistics() *Stats {
	log.Printf("Getting general statistics")

	var stats Stats
	statsURL := fmt.Sprintf(statsURLPattern, c.protocol, c.hostname, c.port)

	req, err := http.NewRequest("GET", statsURL, nil)
	if err != nil {
		log.Fatal("An error has occurred when creating HTTP statistics request ", err)
	}

	if c.isUsingPassword() {
		c.authenticateRequest(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("An error has occurred during login to Adguard: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Unable to read Adguard statistics HTTP response", err)
	}

	err = json.Unmarshal(body, &stats)
	if err != nil {
		log.Println("Unable to unmarshal Adguard statistics to statistics struct model", err)
	}

	return &stats
}

// Function to get the log metrics
func (c *Client) setLogMetrics(logdata *LogData) {
	m = make(map[string]int)
	for i := range logdata.Data {
		logstats := logdata.Data[i]
		if logstats.DNS != nil {
			for j := range logstats.DNS {
				dnsType := logstats.DNS[j].Type
				m[dnsType] += 1
			}
		}
	}

	for key, value := range m {
		metrics.QueryTypes.WithLabelValues(c.hostname, key).Set(float64(value))
	}

	for k := range m {
		delete(m, k)
	}
}

// Function to get the log stats
func (c *Client) getLogStatistics() *LogData {
	log.Printf("Getting log statistics")

	var logdata LogData
	logstatsURL := fmt.Sprintf(logstatsURLPattern, c.protocol, c.hostname, c.port)

	req, err := http.NewRequest("GET", logstatsURL, nil)
	if err != nil {
		log.Fatal("An error has occurred when creating HTTP statistics request ", err)
	}

	if c.isUsingPassword() {
		c.authenticateRequest(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("An error has occurred during login to Adguard: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Unable to read Adguard statistics HTTP response", err)
	}

	err = json.Unmarshal(body, &logdata)
	if err != nil {
		log.Println("Unable to unmarshal Adguard log statistics to log statistics struct model", err)
	}

	return &logdata
}

func (c *Client) isUsingPassword() bool {
	return len(c.b64password) > 0
}

func (c *Client) authenticateRequest(req *http.Request) {
	req.Header.Add("Authorization", "Basic "+c.b64password)
}
