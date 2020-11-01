package adguard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/eko/adguard-exporter/internal/metrics"
)

var (
	loginURLPattern = "%s://%s:%d/login.html"
	statsURLPattern = "%s://%s:%d/control/stats"
)

// Client struct is a AdGuard  client to request an instance of a AdGuard  ad blocker.
type Client struct {
	httpClient http.Client
	interval   time.Duration
	protocol   string
	hostname   string
	port       uint16
	password   string
	sessionID  string
}

// NewClient method initializes a new AdGuard  client.
func NewClient(protocol, hostname string, port uint16, password, interval time.Duration) *Client {
	if protocol != "http" && protocol != "https" {
		log.Printf("protocol %s is invalid. Must be http or https.", protocol)
		os.Exit(1)
	}

	return &Client{
		protocol: protocol,
		hostname: hostname,
		port:     port,
		password: password,
		interval: interval,
		httpClient: http.Client{
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

// Scrape method authenticates and retrieves statistics from AdGuard  JSON API
// and then pass them as Prometheus metrics.
func (c *Client) Scrape() {
	for range time.Tick(c.interval) {
		stats := c.getStatistics()

		c.setMetrics(stats)

		log.Printf("New tick of statistics: %s", stats.ToString())
	}
}

func (c *Client) setMetrics(stats *Stats) {
	metrics.AvgProcessingTime.WithLabelValues(c.hostname).Set(float64(stats.AvgProcessingTime))
}

func (c *Client) getPHPSessionID() (sessionID string) {
	loginURL := fmt.Sprintf(loginURLPattern, c.protocol, c.hostname, c.port)
	values := url.Values{"pw": []string{c.password}}

	req, err := http.NewRequest("POST", loginURL, strings.NewReader(values.Encode()))
	if err != nil {
		log.Fatal("An error has occured when creating HTTP statistics request", err)
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(values.Encode())))

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Printf("An error has occured during login to Adguard: %v", err)
	}

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "PHPSESSID" {
			sessionID = cookie.Value
			break
		}
	}

	return
}

func (c *Client) getStatistics() *Stats {
	var stats Stats

	statsURL := fmt.Sprintf(statsURLPattern, c.protocol, c.hostname, c.port)

	req, err := http.NewRequest("GET", statsURL, nil)
	if err != nil {
		log.Fatal("An error has occured when creating HTTP statistics request", err)
	}

	if c.isUsingPassword() {
		c.authenticateRequest(req)
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		log.Println("An error has occured during retrieving Adguard statistics", err)
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

func (c *Client) isUsingPassword() bool {
	return len(c.password) > 0
}

func (c *Client) authenticateRequest(req *http.Request) {
	cookie := http.Cookie{Name: "PHPSESSID", Value: c.getPHPSessionID()}
	req.AddCookie(&cookie)
}
