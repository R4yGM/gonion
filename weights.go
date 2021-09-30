package gonion

// GetWeights returns results from https://onionoo.torproject.org/weights.
func (gc *GonionClient) GetWeights(args Params) (*Weights, error) {
	weights := &Weights{}
	err := gc.getEndp("weights", args, weights)
	if err != nil {
		return nil, err
	}
	return weights, nil
}

// Weights represents the datastructure defined by Onionoo.
// See https://metrics.torproject.org/onionoo.html#weights.
type Weights struct {
	Version          string       `json:"version"`
	BuildRevision    string       `json:"build_revision"`
	RelaysPublished  string       `json:"relays_published"`
	Relays           []WeightNode `json:"relays"`
	RelaysTruncated  *int         `json:"relays_truncated,omitempty"`
	BridgesPublished string       `json:"bridges_published"`
	Bridges          []WeightNode `json:"bridges"`
}

// WeightNode is a sub-Weights datastructure.
type WeightNode struct {
	Fingerprint             string   `json:"fingerprint"`
	ConsensusWeightFraction *History `json:"consensus_weight_fraction,omitempty"`
	GuardProbability        *History `json:"guard_probability,omitempty"`
	MiddleProbability       *History `json:"middle_probability,omitempty"`
	ExitProbability         *History `json:"exit_probability,omitempty"`
	ConsensusWeight         *History `json:"consensus_weight,omitempty"`
}
