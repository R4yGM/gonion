package gonion

// GetSummary returns results from https://onionoo.torproject.org/summary.
func (gc *GonionClient) GetSummary(args Params) (*Summary, error) {
	summary := &Summary{}
	err := gc.getEndp("summary", args, summary)
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
	RelaysPublished  string          `json:"relays_published"`
	Relays           []SummaryRelay  `json:"relays"`
	BridgesPublished string          `json:"bridges_published"`
	Bridges          []SummaryBridge `json:"bridges"`
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
