package gonion_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

func TestGetWeights(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		Client          gonion.HTTPClient
		Params          gonion.Params
		ExpectedWeights *gonion.Weights
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			Params:          gonion.Params{},
			ExpectedWeights: nil,
			ExpectedErr:     gonion.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			Params:          gonion.Params{},
			ExpectedWeights: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			Params:          gonion.Params{},
			ExpectedWeights: nil,
			ExpectedErr: &gonion.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:          gonion.Params{},
			ExpectedWeights: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"version":"8.0","build_revision":"a0fbbe2","relays_published":"2021-10-01 11:00:00","relays":[{"fingerprint":"000853C5B75A25D6960A0910D40A5F2B210B7D3A","consensus_weight_fraction":{"1_month":{"first":"2021-09-01 14:00:00","last":"2021-10-01 10:00:00","interval":14400,"factor":9.860446146146146e-8,"count":180,"values":[729,726,722]}},"guard_probability":{"1_month":{"first":"2021-09-01 14:00:00","last":"2021-10-01 10:00:00","interval":14400,"factor":1.5754877277277277e-7,"count":180,"values":[781,779]}},"middle_probability":{"1_month":{"first":"2021-09-01 14:00:00","last":"2021-10-01 10:00:00","interval":14400,"factor":1.181378158158158e-7,"count":180,"values":[769,766]}},"exit_probability":{"1_month":{"first":"2021-09-01 14:00:00","last":"2021-10-01 10:00:00","interval":14400,"factor":0,"count":180,"values":[0,0]}},"consensus_weight":{"1_month":{"first":"2021-09-01 14:00:00","last":"2021-10-01 10:00:00","interval":14400,"factor":9.75975975975976,"count":180,"values":[963,963,961]}}}],"relays_truncated":7616,"bridges_published":"2021-10-01 10:50:31","bridges":[]}`, http.StatusOK, nil),
			Params: gonion.Params{
				Limit:  i(1),
				Fields: &gonion.CommaSepList{"test1", "test2"},
			},
			ExpectedWeights: &gonion.Weights{
				Version:         "8.0",
				BuildRevision:   "a0fbbe2",
				RelaysPublished: "2021-10-01 11:00:00",
				Relays: []gonion.WeightNode{
					{
						Fingerprint: "000853C5B75A25D6960A0910D40A5F2B210B7D3A",
						ConsensusWeightFraction: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-09-01 14:00:00",
								Last:     "2021-10-01 10:00:00",
								Interval: 14400,
								Factor:   9.860446146146146e-8,
								Count:    180,
								Values: []*int{
									i(729),
									i(726),
									i(722),
								},
							},
						},
						GuardProbability: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-09-01 14:00:00",
								Last:     "2021-10-01 10:00:00",
								Interval: 14400,
								Factor:   1.5754877277277277e-7,
								Count:    180,
								Values: []*int{
									i(781),
									i(779),
								},
							},
						},
						MiddleProbability: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-09-01 14:00:00",
								Last:     "2021-10-01 10:00:00",
								Interval: 14400,
								Factor:   1.181378158158158e-7,
								Count:    180,
								Values: []*int{
									i(769),
									i(766),
								},
							},
						},
						ExitProbability: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-09-01 14:00:00",
								Last:     "2021-10-01 10:00:00",
								Interval: 14400,
								Factor:   0,
								Count:    180,
								Values: []*int{
									i(0),
									i(0),
								},
							},
						},
						ConsensusWeight: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-09-01 14:00:00",
								Last:     "2021-10-01 10:00:00",
								Interval: 14400,
								Factor:   9.75975975975976,
								Count:    180,
								Values: []*int{
									i(963),
									i(963),
									i(961),
								},
							},
						},
					},
				},
				RelaysTruncated:  i(7616),
				BridgesPublished: "2021-10-01 10:50:31",
				Bridges:          []gonion.WeightNode{},
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			weights, err := gonion.GetWeights(tt.Client, tt.Params)

			assert.Equal(tt.ExpectedWeights, weights)
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
