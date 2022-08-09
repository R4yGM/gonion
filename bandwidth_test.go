package gonion_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

func TestGetBandwidth(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		Client            gonion.HTTPClient
		Params            gonion.Params
		ExpectedBandwidth *gonion.Bandwidth
		ExpectedErr       error
	}{
		"nil-client": {
			Client:            nil,
			Params:            gonion.Params{},
			ExpectedBandwidth: nil,
			ExpectedErr:       gonion.ErrNilClient,
		},
		"failing-client": {
			Client:            newFakeHTTPClient(``, 0, errFake),
			Params:            gonion.Params{},
			ExpectedBandwidth: nil,
			ExpectedErr:       errFake,
		},
		"unexpected-statuscode": {
			Client:            newFakeHTTPClient(``, 0, nil),
			Params:            gonion.Params{},
			ExpectedBandwidth: nil,
			ExpectedErr: &gonion.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:            newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:            gonion.Params{},
			ExpectedBandwidth: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"version":"8.0","build_revision":"a0fbbe2","relays_published":"2021-09-30 13:00:00","relays":[{"fingerprint":"000853C5B75A25D6960A0910D40A5F2B210B7D3A","write_history":{"1_month":{"first":"2021-08-30 12:00:00","last":"2021-09-29 12:00:00","interval":86400,"factor":2723.6236232700335,"count":31,"values":[999]},"6_months":{"first":"2021-08-22 12:00:00","last":"2021-09-29 12:00:00","interval":86400,"factor":2882.4591997800267,"count":39,"values":[1,1,28,320]}},"read_history":{"1_month":{"first":"2021-08-30 12:00:00","last":"2021-09-29 12:00:00","interval":86400,"factor":2697.803559480341,"count":31,"values":[999,564]},"6_months":{"first":"2021-08-22 12:00:00","last":"2021-09-29 12:00:00","interval":86400,"factor":2858.2600114872075,"count":39,"values":[2]}}}],"relays_truncated":7291,"bridges_published":"2021-09-30 12:41:52","bridges":[],"bridges_truncated":1454}`, http.StatusOK, nil),
			Params: gonion.Params{
				Limit:  ptr(4),
				Fields: &gonion.CommaSepList{"test1", "test2"},
			},
			ExpectedBandwidth: &gonion.Bandwidth{
				Version:         "8.0",
				BuildRevision:   "a0fbbe2",
				RelaysPublished: "2021-09-30 13:00:00",
				Relays: []gonion.BandwidthNode{
					{
						Fingerprint: "000853C5B75A25D6960A0910D40A5F2B210B7D3A",
						WriteHistory: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-08-30 12:00:00",
								Last:     "2021-09-29 12:00:00",
								Interval: 86400,
								Factor:   2723.6236232700335,
								Count:    31,
								Values: []*int{
									ptr(999),
								},
							},
							SixMonths: &gonion.HistoryEntry{
								First:    "2021-08-22 12:00:00",
								Last:     "2021-09-29 12:00:00",
								Interval: 86400,
								Factor:   2882.4591997800267,
								Count:    39,
								Values: []*int{
									ptr(1),
									ptr(1),
									ptr(28),
									ptr(320),
								},
							},
						},
						ReadHistory: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2021-08-30 12:00:00",
								Last:     "2021-09-29 12:00:00",
								Interval: 86400,
								Factor:   2697.803559480341,
								Count:    31,
								Values: []*int{
									ptr(999),
									ptr(564),
								},
							},
							SixMonths: &gonion.HistoryEntry{
								First:    "2021-08-22 12:00:00",
								Last:     "2021-09-29 12:00:00",
								Interval: 86400,
								Factor:   2858.2600114872075,
								Count:    39,
								Values: []*int{
									ptr(2),
								},
							},
						},
					},
				},
				RelaysTruncated:  ptr(7291),
				BridgesPublished: "2021-09-30 12:41:52",
				Bridges:          []gonion.BandwidthNode{},
				BridgesTruncated: ptr(1454),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			bandwidth, err := gonion.GetBandwidth(tt.Client, tt.Params)

			assert.Equal(tt.ExpectedBandwidth, bandwidth)
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
