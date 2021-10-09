package gonion

// GetClients returns results from https://onionoo.torproject.org/clients.
func GetClients(client HTTPClient, args Params) (*Clients, error) {
	clients := &Clients{}
	err := getEndp(client, "clients", args, clients)
	if err != nil {
		return nil, err
	}
	return clients, nil
}

// Clients represents the datastructure defined by Onionoo.
// See https://metrics.torproject.org/onionoo.html#clients.
type Clients struct {
	Version          string          `json:"version"`
	BuildRevision    string          `json:"build_revision"`
	RelaysSkipped    *int            `json:"relays_skipped,omitempty"`
	RelaysPublished  string          `json:"relays_published"`
	Relays           []interface{}   `json:"relays"`
	BridgesSkipped   *int            `json:"bridges_skipped,omitempty"`
	BridgesPublished string          `json:"bridges_published"`
	Bridges          []ClientsBridge `json:"bridges"`
	BridgesTruncated *int            `json:"bridges_truncated,omitempty"`
}

// ClientsBridge is a sub-Clients datastructure.
type ClientsBridge struct {
	Fingerprint    string   `json:"fingerprint"`
	AverageClients *History `json:"average_clients,omitempty"`
}
