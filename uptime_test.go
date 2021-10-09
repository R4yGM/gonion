package gonion_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

func TestGetUptime(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		Client         gonion.HTTPClient
		Params         gonion.Params
		ExpectedUptime *gonion.Uptime
		ExpectedErr    error
	}{
		"nil-client": {
			Client:         nil,
			Params:         gonion.Params{},
			ExpectedUptime: nil,
			ExpectedErr:    gonion.ErrNilClient,
		},
		"failing-client": {
			Client:         newFakeHTTPClient(``, 0, errFake),
			Params:         gonion.Params{},
			ExpectedUptime: nil,
			ExpectedErr:    errFake,
		},
		"unexpected-statuscode": {
			Client:         newFakeHTTPClient(``, 0, nil),
			Params:         gonion.Params{},
			ExpectedUptime: nil,
			ExpectedErr: &gonion.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:         newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:         gonion.Params{},
			ExpectedUptime: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"version":"8.0","build_revision":"a0fbbe2","relays_published":"2021-10-01 11:00:00","relays":[{"fingerprint":"000853C5B75A25D6960A0910D40A5F2B210B7D3A","uptime":{"1_month":{"first":"2021-09-01 14:00:00","last":"2021-10-01 10:00:00","interval":14400,"factor":0.001001001001001001,"count":180,"values":[999,999]},"6_months":{"first":"2021-08-22 18:00:00","last":"2021-10-01 06:00:00","interval":43200,"factor":0.001001001001001001,"count":80,"values":[999]}},"flags":{"Guard":{"1_month":{"first":"2021-09-01 14:00:00","last":"2021-10-01 10:00:00","interval":14400,"factor":0.001001001001001001,"count":180,"values":[999]},"6_months":{"first":"2021-08-30 18:00:00","last":"2021-10-01 06:00:00","interval":43200,"factor":0.001001001001001001,"count":64,"values":[999,999]}}}}],"relays_truncated":7616,"bridges_published":"2021-10-01 10:50:31","bridges":[],"bridges_truncated":1450}`, http.StatusOK, nil),
			Params: gonion.Params{
				Limit: i(1),
			},
			ExpectedUptime: &gonion.Uptime{
				Version:         "8.0",
				BuildRevision:   "a0fbbe2",
				RelaysPublished: "2021-10-01 11:00:00",
				Relays: []gonion.UptimeRelay{
					{
						Fingerprint: "000853C5B75A25D6960A0910D40A5F2B210B7D3A",
						Uptime: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-09-01 14:00:00",
								Last:     "2021-10-01 10:00:00",
								Interval: 14400,
								Factor:   0.001001001001001001,
								Count:    180,
								Values: []*int{
									i(999),
									i(999),
								},
							},
							SixMonths: &gonion.HistoryEntry{
								First:    "2021-08-22 18:00:00",
								Last:     "2021-10-01 06:00:00",
								Interval: 43200,
								Factor:   0.001001001001001001,
								Count:    80,
								Values: []*int{
									i(999),
								},
							},
						},
						Flags: &gonion.Flags{
							Guard: &gonion.History{
								OneMonth: &gonion.HistoryEntry{
									First:    "2021-09-01 14:00:00",
									Last:     "2021-10-01 10:00:00",
									Interval: 14400,
									Factor:   0.001001001001001001,
									Count:    180,
									Values: []*int{
										i(999),
									},
								},
								SixMonths: &gonion.HistoryEntry{
									First:    "2021-08-30 18:00:00",
									Last:     "2021-10-01 06:00:00",
									Interval: 43200,
									Factor:   0.001001001001001001,
									Count:    64,
									Values: []*int{
										i(999),
										i(999),
									},
								},
							},
						},
					},
				},
				RelaysTruncated:  i(7616),
				BridgesPublished: "2021-10-01 10:50:31",
				Bridges:          []interface{}{},
				BridgesTruncated: i(1450),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			uptime, err := gonion.GetUptime(tt.Client, tt.Params)

			assert.Equal(tt.ExpectedUptime, uptime)
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
