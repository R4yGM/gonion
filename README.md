# Gonion

[![GoDoc](https://godoc.org/github.com/R4yGM/gonion/gonion?status.svg)](http://godoc.org/github.com/R4yGM/gonion)
[![Go Report Card](https://goreportcard.com/badge/github.com/R4yGM/gonion)](https://goreportcard.com/report/github.com/R4yGM/gonion)

Lightweight Golang wrapper for querying Tor network data using the Onionoo service.

## How to use

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/R4yGM/gonion"
)

func main() {
	// Build the HTTP client to use.
	client := &http.Client{}

	// Build parameters.
	params := gonion.Params{
		Search:             str("R4y"),
		Running:            b(true),
		RecommendedVersion: b(true),
	}

	// Issue the request.
	res, err := gonion.GetDetails(client, params)
	if err != nil {
		log.Fatal(err)
	}

	// Use the results.
	for _, relay := range res.Relays {
		fmt.Println(relay.Nickname)
	}
}

func str(str string) *string {
	return &str
}

func b(b bool) *bool {
	return &b
}
```

## Support

The following endpoints are supported (tested through unit and integration tests):
 - bandwidth
 - clients
 - details
 - summary
 - uptime
 - weights
