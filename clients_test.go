package gonion_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

func TestGetClients(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		Client          gonion.HTTPClient
		Params          gonion.Params
		ExpectedClients *gonion.Clients
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			Params:          gonion.Params{},
			ExpectedClients: nil,
			ExpectedErr:     gonion.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			Params:          gonion.Params{},
			ExpectedClients: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			Params:          gonion.Params{},
			ExpectedClients: nil,
			ExpectedErr: &gonion.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:          gonion.Params{},
			ExpectedClients: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"version":"8.0","build_revision":"a0fbbe2","relays_published":"2021-09-30 14:00:00","relays":[],"bridges_published":"2021-09-30 13:41:52","bridges":[{"fingerprint":"000A69A59E5624D2ED2AEF827429108EF5FC0F4A"},{"fingerprint":"00782946F4C54CE1D028F21E541EF8440ECAA0EE","average_clients":{"1_month":{"first":"2020-03-06 12:00:00","last":"2020-04-05 12:00:00","interval":86400,"factor":0.0057462462462462465,"count":31,"values":[487,501]},"6_months":{"first":"2019-10-06 12:00:00","last":"2020-04-05 12:00:00","interval":86400,"factor":3.693037037037037,"count":183,"values":[74,366,664]},"1_year":{"first":"2019-04-07 00:00:00","last":"2020-04-05 00:00:00","interval":172800,"factor":3.3355055555555557,"count":183,"values":[0,null]},"5_years":{"first":"2015-05-10 00:00:00","last":"2020-03-24 00:00:00","interval":864000,"factor":3.028332762762763,"count":179,"values":[]}}}],"bridges_truncated":1449}`, http.StatusOK, nil),
			Params: gonion.Params{
				Limit:  i(4),
				Fields: &gonion.CommaSepList{"test1", "test2"},
			},
			ExpectedClients: &gonion.Clients{
				Version:          "8.0",
				BuildRevision:    "a0fbbe2",
				RelaysPublished:  "2021-09-30 14:00:00",
				Relays:           []interface{}{},
				BridgesPublished: "2021-09-30 13:41:52",
				Bridges: []gonion.ClientsBridge{
					{
						Fingerprint: "000A69A59E5624D2ED2AEF827429108EF5FC0F4A",
					}, {
						Fingerprint: "00782946F4C54CE1D028F21E541EF8440ECAA0EE",
						AverageClients: &gonion.History{
							OneMonth: &gonion.HistoryEntry{
								First:    "2020-03-06 12:00:00",
								Last:     "2020-04-05 12:00:00",
								Interval: 86400,
								Factor:   0.0057462462462462465,
								Count:    31,
								Values: []*int{
									i(487),
									i(501),
								},
							},
							SixMonths: &gonion.HistoryEntry{
								First:    "2019-10-06 12:00:00",
								Last:     "2020-04-05 12:00:00",
								Interval: 86400,
								Factor:   3.693037037037037,
								Count:    183,
								Values: []*int{
									i(74),
									i(366),
									i(664),
								},
							},
							OneYear: &gonion.HistoryEntry{
								First:    "2019-04-07 00:00:00",
								Last:     "2020-04-05 00:00:00",
								Interval: 172800,
								Factor:   3.3355055555555557,
								Count:    183,
								Values: []*int{
									i(0),
									nil,
								},
							},
							FiveYears: &gonion.HistoryEntry{
								First:    "2015-05-10 00:00:00",
								Last:     "2020-03-24 00:00:00",
								Interval: 864000,
								Factor:   3.028332762762763,
								Count:    179,
								Values:   []*int{},
							},
						},
					},
				},
				BridgesTruncated: i(1449),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			clients, err := gonion.GetClients(tt.Client, tt.Params)

			assert.Equal(tt.ExpectedClients, clients)
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
