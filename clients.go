package gonion

// GetClients returns results from https://onionoo.torproject.org/clients.
func (gc *GonionClient) GetClients(args Params) (*Clients, error) {
	clients := &Clients{}
	err := gc.getEndp("clients", args, clients)
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
	RelaysPublished  string          `json:"relays_published"`
	Relays           []interface{}   `json:"relays"`
	BridgesPublished string          `json:"bridges_published"`
	Bridges          []ClientsBridge `json:"bridges"`
	BridgesTruncated *int            `json:"bridges_truncated,omitempty"`
}

// ClientsBridge is a sub-Clients datastructure.
type ClientsBridge struct {
	Fingerprint    string   `json:"fingerprint"`
	AverageClients *History `json:"average_clients,omitempty"`
}
