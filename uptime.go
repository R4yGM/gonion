package gonion

// GetUptime returns results from https://onionoo.torproject.org/uptime.
func GetUptime(client HTTPClient, args Params, opts ...Option) (*Uptime, error) {
	uptime := &Uptime{}
	if err := getEndp(client, "uptime", args, uptime, opts...); err != nil {
		return nil, err
	}
	return uptime, nil
}

// Uptime represents the datastructure defined by Onionoo.
// See https://metrics.torproject.org/onionoo.html#uptime.
type Uptime struct {
	Version          string        `json:"version"`
	BuildRevision    string        `json:"build_revision"`
	RelaysPublished  string        `json:"relays_published"`
	RelaysSkipped    *int          `json:"relays_skipped,omitempty"`
	Relays           []UptimeRelay `json:"relays"`
	RelaysTruncated  *int          `json:"relays_truncated,omitempty"`
	BridgesPublished string        `json:"bridges_published"`
	Bridges          []any         `json:"bridges"`
	BridgesTruncated *int          `json:"bridges_truncated,omitempty"`
}

// UptimeRelay is a sub-Uptime datastructure.
type UptimeRelay struct {
	Fingerprint string   `json:"fingerprint"`
	Uptime      *History `json:"uptime,omitempty"`
	Flags       *Flags   `json:"flags,omitempty"`
}

// Flags is a sub-UptimeRelay datastructure.
type Flags struct {
	Exit      *History `json:"Exit,omitempty"`
	Fast      *History `json:"Fast,omitempty"`
	Guard     *History `json:"Guard,omitempty"`
	HSDir     *History `json:"HSDir,omitempty"`
	Running   *History `json:"Running,omitempty"`
	Stable    *History `json:"Stable,omitempty"`
	StaleDesc *History `json:"StaleDesc,omitempty"`
	V2Dir     *History `json:"V2Dir,omitempty"`
	Valid     *History `json:"Valid,omitempty"`
}

// UptimeBridge is a sub-Uptime datastructure.
type UptimeBridge struct {
	Fingerprint string  `json:"fingerprint"`
	Uptime      History `json:"uptime"`
}
