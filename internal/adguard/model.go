package adguard

import "fmt"

// SllStats struct containing all Adguard statistics structs
type AllStats struct {
	status   *Status
	stats    *Stats
	logStats *LogStats
}

// Status struct is the Adguard statistics JSON API corresponding model.
type Status struct {
	Dhcp              bool     `json:"dhcp_available"`
	DNSAddresses      []string `json:"dns_addresses"`
	DNSPort           int      `json:"dns_port"`
	HttpPort          int      `json:"http_port"`
	Language          string   `json:"language"`
	ProtectionEnabled bool     `json:"protection_enabled"`
	Running           bool     `json:"running"`
	Version           string   `json:"version"`
}

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

type DNSHeader struct {
	Name     string `json:"Name"`
	Rrtype   int    `json:"Rrtype"`
	Class    int    `json:"Class"`
	TTL      int    `json:"Ttl"`
	Rdlength int    `json:"Rdlength"`
}

type Type65 struct {
	Hdr   DNSHeader `json:"Hdr"`
	RData string    `json:"Rdata"`
}

// DNSAnswer struct from LogData
type DNSAnswer struct {
	TTL   float64     `json:"ttl"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"` // DNSAnswer struct can change sometimes... value:string or value: { "Hdr": { "Name":string, "Rrtype":int, "Class":int, "Ttl":int, "Rdlength":int }, "RData":string }
}

// DNSQuery struct from LogData
type DNSQuery struct {
	Class string `json:"class"`
	Host  string `json:"host"`
	Type  string `json:"type"`
}

// LogData struct, sub struct of LogStats to collect the dns stats from the log
type LogData struct {
	Answer      []DNSAnswer `json:"answer"`
	DNSSec      bool        `json:"answer_dnssec"`
	Client      string      `json:"client"`
	ClientProto string      `json:"client_proto"`
	Elapsed     string      `json:"elapsedMs"`
	Question    DNSQuery    `json:"question"`
	Reason      string      `json:"reason"`
	Status      string      `json:"status"`
	Time        string      `json:"time"`
	Upstream    string      `json:"upstream"`
}

// LogStats struct for the Adguard log statistics JSON API corresponding model.
type LogStats struct {
	Data   []LogData `json:"data"`
	Oldest string    `json:"oldest"`
}

// ToString method returns a string of the current statistics struct.
func (s *Stats) ToString() string {
	return fmt.Sprintf("%d ads blocked / %d total DNS queries", s.BlockedFiltering, s.DnsQueries)
}
