//go:build integration
// +build integration

package integration_test

import "github.com/R4yGM/gonion"

var paramsSet = map[string]gonion.Params{
	"limit-4-offset-10": {
		Limit:  i(4),
		Offset: i(10),
	},
	"search-maria": {
		Search: str("maria"),
	},
}

func str(str string) *string {
	return &str
}

func i(i int) *int {
	return &i
}
