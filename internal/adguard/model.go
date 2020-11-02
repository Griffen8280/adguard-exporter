package adguard

import "fmt"

// Stats struct is the Adguard statistics JSON API corresponding model.
type Stats struct {
  AvgProcessingTime     float64            `json:"avg_processing_time"`
  DnsQueries            int                `json:"num_dns_queries"`
  BlockedFiltering      int                `json:"num_blocked_filtering"`
  ParentalFiltering     int                `json:"num_replaced_parental"`
  SafeBrowsingFiltering int                `json:"num_replaced_safebrowsing"`
  SafeSearchFiltering   int                `json:"num_replaced_safesearch"`
  TopQueries            []map[string]int   `json:"top_queried_domains"`
  TopBlocked            []map[string]int   `json:"top_blocked_domains"`
  TopClients            []map[string]int   `json:"top_clients"`
}

// ToString method returns a string of the current statistics struct.
func (s *Stats) ToString() string {
  return fmt.Sprintf("%d ads blocked / %d total DNS queries", s.BlockedFiltering, s.DnsQueries)
}
