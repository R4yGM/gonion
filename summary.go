package gonion

// GetSummary returns results from https://onionoo.torproject.org/summary.
func GetSummary(client HTTPClient, args Params) (*Summary, error) {
	summary := &Summary{}
	err := getEndp(client, "summary", args, summary)
	if err != nil {
		return nil, err
	}
	return summary, nil
}

// Summary represents the datastructure defined by Onionoo.
// See https://metrics.torproject.org/onionoo.html#summary.
type Summary struct {
	Version          string          `json:"version"`
	BuildRevision    string          `json:"build_revision"`
	RelaysSkipped    *int            `json:"relays_skipped,omitempty"`
	RelaysPublished  string          `json:"relays_published"`
	Relays           []SummaryRelay  `json:"relays"`
	RelaysTruncated  *int            `json:"relays_truncated,omitempty"`
	BridgesPublished string          `json:"bridges_published"`
	Bridges          []SummaryBridge `json:"bridges"`
	BridgesTruncated *int            `json:"bridges_truncated,omitempty"`
}

// SummaryRelay is a sub-Summary datastructure.
type SummaryRelay struct {
	N string   `json:"n"`
	F string   `json:"f"`
	A []string `json:"a"`
	R bool     `json:"r"`
}

// SummaryBridge is a sub-Summary datastructure.
type SummaryBridge struct {
	N string `json:"n"`
	H string `json:"h"`
	R bool   `json:"r"`
}
