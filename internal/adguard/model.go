package adguard

import "fmt"

// Stats struct is the Adguard statistics JSON API corresponding model.
type Stats struct {
	AvgProcessingTime     float64          `json:"avg_processing_time"`
	DnsQueries            int              `json:"num_dns_queries"`
	BlockedFiltering      int              `json:"num_blocked_filtering"`
	ParentalFiltering     int              `json:"num_replaced_parental"`
	SafeBrowsingFiltering int              `json:"num_replaced_safebrowsing"`
	SafeSearchFiltering   int              `json:"num_replaced_safesearch"`
	TopQueries            []map[string]int `json:"top_queried_domains"`
	TopBlocked            []map[string]int `json:"top_blocked_domains"`
	TopClients            []map[string]int `json:"top_clients"`
}

// DNS answer struct in the log
type DNSAnswer struct {
	Ttl   float64 `json:"ttl"`
	Type  string  `json:"type"`
	Value string  `json:"value"`
}

// DNS query struct in the log
type DNSQuery struct {
	Class string `json:"class"`
	Host  string `json:"host"`
	Type  string `json:"type"`
}

// LogStats struct for the Adguard log statistics JSON API corresponding model.
type LogStats struct {
	DNS          []DNSAnswer `json:"answer"`
	DNSSec       bool        `json:"answer_dnssec"`
	Client       string      `json:"client"`
	Client_proto string      `json:"client_proto"`
	Elapsed      string      `json:"elapsedMs"`
	Question     DNSQuery    `json:"question"`
	Reason       string      `json:"reason"`
	Status       string      `json:"status"`
	Time         string      `json:"time"`
	Upstream     string      `json:"upstream"`
}

type LogData struct {
	Data   []LogStats `json:"data"`
	Oldest string     `json:"oldest"`
}

// ToString method returns a string of the current statistics struct.
func (s *Stats) ToString() string {
	return fmt.Sprintf("%d ads blocked / %d total DNS queries", s.BlockedFiltering, s.DnsQueries)
}
