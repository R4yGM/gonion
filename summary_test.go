package gonion_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

func TestGetSummary(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		Client          gonion.HTTPClient
		Params          gonion.Params
		ExpectedSummary *gonion.Summary
		ExpectedErr     error
	}{
		"nil-client": {
			Client:          nil,
			Params:          gonion.Params{},
			ExpectedSummary: nil,
			ExpectedErr:     gonion.ErrNilClient,
		},
		"failing-client": {
			Client:          newFakeHTTPClient(``, 0, errFake),
			Params:          gonion.Params{},
			ExpectedSummary: nil,
			ExpectedErr:     errFake,
		},
		"unexpected-statuscode": {
			Client:          newFakeHTTPClient(``, 0, nil),
			Params:          gonion.Params{},
			ExpectedSummary: nil,
			ExpectedErr: &gonion.ErrUnexpectedStatus{
				StatusCode: 0,
				Body:       []byte(``),
			},
		},
		"failing-unmarshal": {
			Client:          newFakeHTTPClient(`{[}]`, http.StatusOK, nil),
			Params:          gonion.Params{},
			ExpectedSummary: nil,
			ExpectedErr: &json.SyntaxError{
				Offset: 2,
			},
		},
		"valid-call": {
			Client: newFakeHTTPClient(`{"version":"8.0","build_revision":"a0fbbe2","relays_published":"2021-09-30 14:00:00","relays":[{"n":"lyesn","f":"000853C5B75A25D6960A0910D40A5F2B210B7D3A","a":["165.22.200.169"],"r":true},{"n":"seele","f":"000A10D43011EA4928A35F610405F92B4433B4DC","a":["98.45.181.220"],"r":true},{"n":"CalyxInstitute14","f":"0011BD2485AD45D984EC4159C88FC066E5E3300E","a":["162.247.74.201"],"r":true},{"n":"skylarkRelay","f":"00240ECB2B535AA4C1E1874D744DFA6AF2E5E941","a":["95.111.230.178"],"r":true}],"relays_truncated":7290,"bridges_published":"2021-09-30 13:41:52","bridges":[],"bridges_truncated":1453}`, http.StatusOK, nil),
			Params: gonion.Params{
				Limit:  ptr(4),
				Fields: &gonion.CommaSepList{"test1", "test2"},
			},
			ExpectedSummary: &gonion.Summary{
				Version:         "8.0",
				BuildRevision:   "a0fbbe2",
				RelaysPublished: "2021-09-30 14:00:00",
				Relays: []gonion.SummaryRelay{
					{
						N: "lyesn",
						F: "000853C5B75A25D6960A0910D40A5F2B210B7D3A",
						A: []string{
							"165.22.200.169",
						},
						R: true,
					}, {
						N: "seele",
						F: "000A10D43011EA4928A35F610405F92B4433B4DC",
						A: []string{
							"98.45.181.220",
						},
						R: true,
					}, {
						N: "CalyxInstitute14",
						F: "0011BD2485AD45D984EC4159C88FC066E5E3300E",
						A: []string{
							"162.247.74.201",
						},
						R: true,
					}, {
						N: "skylarkRelay",
						F: "00240ECB2B535AA4C1E1874D744DFA6AF2E5E941",
						A: []string{
							"95.111.230.178",
						},
						R: true,
					},
				},
				RelaysTruncated:  ptr(7290),
				BridgesPublished: "2021-09-30 13:41:52",
				Bridges:          []gonion.SummaryBridge{},
				BridgesTruncated: ptr(1453),
			},
			ExpectedErr: nil,
		},
	}

	for testname, tt := range tests {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			summary, err := gonion.GetSummary(tt.Client, tt.Params)

			assert.Equal(tt.ExpectedSummary, summary)
			checkErr(err, tt.ExpectedErr, t)
		})
	}
}
