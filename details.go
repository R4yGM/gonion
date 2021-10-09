package gonion

// GetDetails returns results from https://onionoo.torproject.org/details.
func GetDetails(client HTTPClient, args Params) (*Details, error) {
	details := &Details{}
	err := getEndp(client, "details", args, details)
	if err != nil {
		return nil, err
	}
	return details, nil
}

// Details represents the datastructure defined by Onionoo.
// See https://metrics.torproject.org/onionoo.html#details.
type Details struct {
	Version          string            `json:"version"`
	BuildRevision    string            `json:"build_revision"`
	RelaysSkipped    *int              `json:"relays_skipped,omitempty"`
	RelaysPublished  string            `json:"relays_published"`
	RelaysTruncated  *int              `json:"relays_truncated,omitempty"`
	Relays           []DetailedRelay   `json:"relays"`
	BridgesPublished string            `json:"bridges_published"`
	Bridges          []DetailedBridges `json:"bridges"`
	BridgesTruncated *int              `json:"bridges_truncated,omitempty"`
}

// DetailedRelay is a sub-Details datastructure.
type DetailedRelay struct {
	Nickname                 string             `json:"nickname"`
	Fingerprint              string             `json:"fingerprint"`
	OrAddresses              []string           `json:"or_addresses"`
	LastSeen                 string             `json:"last_seen"`
	Hibernating              *bool              `json:"hibernating,omitempty"`
	LastChangedAddressOrPort string             `json:"last_changed_address_or_port"`
	FirstSeen                string             `json:"first_seen"`
	Running                  bool               `json:"running"`
	UnverifiedHostNames      *[]string          `json:"unverified_host_names,omitempty"`
	UnreachableOrAddresses   *[]string          `json:"unreachable_or_addresses,omitempty"`
	Flags                    []string           `json:"flags"`
	Country                  string             `json:"country"`
	CountryName              string             `json:"country_name"`
	As                       string             `json:"as"`
	AsName                   *string            `json:"as_name,omitempty"`
	ConsensusWeight          int                `json:"consensus_weight"`
	LastRestarted            string             `json:"last_restarted"`
	BandwidthRate            int                `json:"bandwidth_rate"`
	BandwidthBurst           int                `json:"bandwidth_burst"`
	ObservedBandwidth        int                `json:"observed_bandwidth"`
	AdvertisedBandwidth      int                `json:"advertised_bandwidth"`
	AllegedFamily            *[]string          `json:"alleged_family,omitempty"`
	ExitPolicy               []string           `json:"exit_policy"`
	ExitPolicySummary        *ExitPolicySummary `json:"exit_policy_summary,omitempty"`
	ExitPolicyV6Summary      *ExitPolicySummary `json:"exit_policy_v6_summary,omitempty"`
	Contact                  *string            `json:"contact,omitempty"`
	OverloadGeneralTimestamp *float64           `json:"overload_general_timestamp,omitempty"`
	Platform                 string             `json:"platform"`
	Version                  string             `json:"version"`
	VersionStatus            string             `json:"version_status"`
	EffectiveFamily          []string           `json:"effective_family"`
	ConsensusWeightFraction  *float64           `json:"consensus_weight_fraction,omitempty"`
	GuardProbability         *float64           `json:"guard_probability,omitempty"`
	IndirectFamily           *[]string          `json:"indirect_family,omitempty"`
	MiddleProbability        *float64           `json:"middle_probability,omitempty"`
	ExitProbability          *float64           `json:"exit_probability,omitempty"`
	RecommendedVersion       bool               `json:"recommended_version"`
	Measured                 bool               `json:"measured"`
	ExitAddresses            *[]string          `json:"exit_addresses,omitempty"`
	DirAddress               *string            `json:"dir_address,omitempty"`
	VerifiedHostNames        *[]string          `json:"verified_host_names,omitempty"`
}

// ExitPolicySummary is a sub-DetailedRelay datastructure.
type ExitPolicySummary struct {
	Reject *[]string `json:"reject,omitempty"`
	Accept *[]string `json:"accept,omitempty"`
}

// DetailedBridges is a sub-Details datastructure.
type DetailedBridges struct {
	Nickname                 string    `json:"nickname"`
	HashedFingerprint        string    `json:"hashed_fingerprint"`
	OrAddresses              []string  `json:"or_addresses"`
	LastSeen                 string    `json:"last_seen"`
	FirstSeen                string    `json:"first_seen"`
	Running                  bool      `json:"running"`
	Flags                    []string  `json:"flags"`
	LastRestarted            string    `json:"last_restarted"`
	AdvertisedBandwidth      int       `json:"advertised_bandwidth"`
	Platform                 string    `json:"platform"`
	Version                  string    `json:"version"`
	VersionStatus            string    `json:"version_status"`
	RecommendedVersion       bool      `json:"recommended_version"`
	BridgedbDistributor      *string   `json:"bridgedb_distributor,omitempty"`
	Transports               *[]string `json:"transports,omitempty"`
	OverloadGeneralTimestamp *float64  `json:"overload_general_timestamp,omitempty"`
}
