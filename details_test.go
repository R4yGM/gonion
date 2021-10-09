package gonion_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

func TestGetDetails(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		Client          gonion.HTTPClient
		Params          gonion.Params
		ExpectedDetails *gonion.Details
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			Params:          gonion.Params{},
			ExpectedDetails: nil,
			ExpectedErr:     gonion.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			Params:          gonion.Params{},
			ExpectedDetails: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			Params:          gonion.Params{},
			ExpectedDetails: nil,
			ExpectedErr: &gonion.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:          gonion.Params{},
			ExpectedDetails: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"version":"8.0","build_revision":"a0fbbe2","relays_published":"2021-09-30 14:00:00","relays":[{"nickname":"lyesn","fingerprint":"000853C5B75A25D6960A0910D40A5F2B210B7D3A","or_addresses":["165.22.200.169:443"],"last_seen":"2021-09-30 14:00:00","last_changed_address_or_port":"2021-08-22 19:00:00","first_seen":"2021-08-22 19:00:00","running":true,"flags":["Fast","Guard","HSDir","Running","Stable","V2Dir","Valid"],"country":"us","country_name":"United States of America","as":"AS14061","consensus_weight":9700,"last_restarted":"2021-08-22 18:39:41","bandwidth_rate":1073741824,"bandwidth_burst":1073741824,"observed_bandwidth":6251355,"advertised_bandwidth":6251355,"exit_policy":["reject *:*"],"exit_policy_summary":{"reject":["1-65535"]},"contact":"lyesn@everyones.tech","platform":"Tor 0.4.6.7 on FreeBSD","version":"0.4.6.7","version_status":"recommended","effective_family":["000853C5B75A25D6960A0910D40A5F2B210B7D3A"],"consensus_weight_fraction":0.00007400913,"guard_probability":0.0001231322,"middle_probability":0.00009454931,"exit_probability":0,"recommended_version":true,"measured":true},{"nickname":"seele","fingerprint":"000A10D43011EA4928A35F610405F92B4433B4DC","or_addresses":["98.45.181.220:9001"],"last_seen":"2021-09-30 14:00:00","last_changed_address_or_port":"2021-07-22 05:00:00","first_seen":"2014-08-23 06:00:00","running":true,"flags":["Running","Stable","V2Dir","Valid"],"country":"us","country_name":"United States of America","as":"AS7922","consensus_weight":11,"last_restarted":"2021-08-28 17:33:51","bandwidth_rate":102400,"bandwidth_burst":204800,"observed_bandwidth":112793,"advertised_bandwidth":102400,"exit_policy":["reject *:*"],"exit_policy_summary":{"reject":["1-65535"]},"contact":"Nicolas Noble <pixel@nobis-crew.org>","platform":"Tor 0.4.5.10 on Linux","version":"0.4.5.10","version_status":"recommended","effective_family":["000A10D43011EA4928A35F610405F92B4433B4DC"],"consensus_weight_fraction":8.392788e-8,"guard_probability":0,"middle_probability":2.4688202e-7,"exit_probability":0,"recommended_version":true,"measured":true}],"relays_truncated":7290,"bridges_published":"2021-09-30 13:41:52","bridges":[],"bridges_truncated":1453}`, http.StatusOK, nil),
			Params: gonion.Params{
				Limit:  i(4),
				Fields: &gonion.CommaSepList{"test1", "test2"},
			},
			ExpectedDetails: &gonion.Details{
				Version:         "8.0",
				BuildRevision:   "a0fbbe2",
				RelaysPublished: "2021-09-30 14:00:00",
				Relays: []gonion.DetailedRelay{
					{
						Nickname:    "lyesn",
						Fingerprint: "000853C5B75A25D6960A0910D40A5F2B210B7D3A",
						OrAddresses: []string{
							"165.22.200.169:443",
						},
						LastSeen:                 "2021-09-30 14:00:00",
						LastChangedAddressOrPort: "2021-08-22 19:00:00",
						FirstSeen:                "2021-08-22 19:00:00",
						Running:                  true,
						Flags: []string{
							"Fast",
							"Guard",
							"HSDir",
							"Running",
							"Stable",
							"V2Dir",
							"Valid",
						},
						Country:             "us",
						CountryName:         "United States of America",
						As:                  "AS14061",
						ConsensusWeight:     9700,
						LastRestarted:       "2021-08-22 18:39:41",
						BandwidthRate:       1073741824,
						BandwidthBurst:      1073741824,
						ObservedBandwidth:   6251355,
						AdvertisedBandwidth: 6251355,
						ExitPolicy: []string{
							"reject *:*",
						},
						ExitPolicySummary: &gonion.ExitPolicySummary{
							Reject: &[]string{
								"1-65535",
							},
						},
						Contact:       str("lyesn@everyones.tech"),
						Platform:      "Tor 0.4.6.7 on FreeBSD",
						Version:       "0.4.6.7",
						VersionStatus: "recommended",
						EffectiveFamily: []string{
							"000853C5B75A25D6960A0910D40A5F2B210B7D3A",
						},
						ConsensusWeightFraction: f(0.00007400913),
						GuardProbability:        f(0.0001231322),
						MiddleProbability:       f(0.00009454931),
						ExitProbability:         f(0),
						RecommendedVersion:      true,
						Measured:                true,
					}, {
						Nickname:    "seele",
						Fingerprint: "000A10D43011EA4928A35F610405F92B4433B4DC",
						OrAddresses: []string{
							"98.45.181.220:9001",
						},
						LastSeen:                 "2021-09-30 14:00:00",
						LastChangedAddressOrPort: "2021-07-22 05:00:00",
						FirstSeen:                "2014-08-23 06:00:00",
						Running:                  true,
						Flags: []string{
							"Running",
							"Stable",
							"V2Dir",
							"Valid",
						},
						Country:             "us",
						CountryName:         "United States of America",
						As:                  "AS7922",
						ConsensusWeight:     11,
						LastRestarted:       "2021-08-28 17:33:51",
						BandwidthRate:       102400,
						BandwidthBurst:      204800,
						ObservedBandwidth:   112793,
						AdvertisedBandwidth: 102400,
						ExitPolicy: []string{
							"reject *:*",
						},
						ExitPolicySummary: &gonion.ExitPolicySummary{
							Reject: &[]string{
								"1-65535",
							},
						},
						Contact:       str("Nicolas Noble <pixel@nobis-crew.org>"),
						Platform:      "Tor 0.4.5.10 on Linux",
						Version:       "0.4.5.10",
						VersionStatus: "recommended",
						EffectiveFamily: []string{
							"000A10D43011EA4928A35F610405F92B4433B4DC",
						},
						ConsensusWeightFraction: f(8.392788e-8),
						GuardProbability:        f(0),
						MiddleProbability:       f(2.4688202e-7),
						ExitProbability:         f(0),
						RecommendedVersion:      true,
						Measured:                true,
					},
				},
				RelaysTruncated:  i(7290),
				BridgesPublished: "2021-09-30 13:41:52",
				Bridges:          []gonion.DetailedBridges{},
				BridgesTruncated: i(1453),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			details, err := gonion.GetDetails(tt.Client, tt.Params)

			assert.Equal(tt.ExpectedDetails, details)
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
