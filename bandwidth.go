package gonion

// GetBandwidth returns results from https://onionoo.torproject.org/bandwidth.
func GetBandwidth(client HTTPClient, args Params) (*Bandwidth, error) {
	bandwidth := &Bandwidth{}
	err := getEndp(client, "bandwidth", args, bandwidth)
	if err != nil {
		return nil, err
	}
	return bandwidth, nil
}

// Bandwidth represents the datastructure defined by Onionoo.
// See https://metrics.torproject.org/onionoo.html#bandwidth.
type Bandwidth struct {
	Version          string          `json:"version"`
	BuildRevision    string          `json:"build_revision"`
	RelaysSkipped    *int            `json:"relays_skipped,omitempty"`
	RelaysPublished  string          `json:"relays_published"`
	Relays           []BandwidthNode `json:"relays"`
	RelaysTruncated  *int            `json:"relays_truncated,omitempty"`
	BridgesPublished string          `json:"bridges_published"`
	Bridges          []BandwidthNode `json:"bridges"`
	BridgesTruncated *int            `json:"bridges_truncated,omitempty"`
}

// BandwidthNode is a sub-Bandwidth datastructure.
type BandwidthNode struct {
	Fingerprint        string              `json:"fingerprint"`
	WriteHistory       *History            `json:"write_history,omitempty"`
	ReadHistory        *History            `json:"read_history,omitempty"`
	OverloadRateLimits *OverloadRateLimits `json:"overload_ratelimits,omitempty"`
}

// History is a sub-BandwidthNode datastructure.
type History struct {
	OneMonth  *HistoryEntry `json:"1_month,omitempty"`
	SixMonths *HistoryEntry `json:"6_months,omitempty"`
	OneYear   *HistoryEntry `json:"1_year,omitempty"`
	FiveYears *HistoryEntry `json:"5_years,omitempty"`
}

// HistoryEntry is a sub-History datastructure.
type HistoryEntry struct {
	First    string  `json:"first"`
	Last     string  `json:"last"`
	Interval int     `json:"interval"`
	Factor   float64 `json:"factor"`
	Count    int     `json:"count"`
	Values   []*int  `json:"values,omitempty"`
}

// OverloadRateLimits is a sub-BandwidthNode datastructure.
type OverloadRateLimits struct {
	BurstLimit float64 `json:"burst-limit"`
	RateLimit  float64 `json:"rate-limit"`
	ReadCount  float64 `json:"read-count"`
	Timestamp  float64 `json:"timestamp"`
	WriteCount float64 `json:"write-count"`
}
