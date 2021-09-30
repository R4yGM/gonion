package integration_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

func TestGetBandwidth(t *testing.T) {
	for testname, params := range paramsSet {
		t.Run(testname, func(t *testing.T) {
			assert := assert.New(t)

			mdwClient := &MdwClient{}
			client, _ := gonion.NewGonionClient(mdwClient, "")

			bandwidth, err := client.GetBandwidth(params)

			// Ensure no error
			if !assert.Nil(err) {
				t.Errorf("Last body [%s]\n", mdwClient.LastBody)
			}

			// Reencode to JSON
			buf := &bytes.Buffer{}
			_ = json.NewEncoder(buf).Encode(bandwidth)

			// Decode both to interfaces
			var expected interface{}
			var actual interface{}
			_ = json.Unmarshal(mdwClient.LastBody, &expected)
			_ = json.Unmarshal(buf.Bytes(), &actual)

			// Compares both to check valid API (and not nil)
			assert.NotNil(expected)
			assert.Equal(expected, actual)
		})
	}
}
