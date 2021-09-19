# gonion

Lightweight Golang wrapper for querying Tor network data using the Onionoo service.

[![GoDoc](https://godoc.org/github.com/R4yGM/gonion/gonion?status.svg)](http://godoc.org/github.com/R4yGM/gonion)
[![Go Report Card](https://goreportcard.com/badge/github.com/R4yGM/gonion)](https://goreportcard.com/report/github.com/R4yGM/gonion)

```go
package main

import (
        "github.com/R4yGM/gonion"
        "net/http"
        "time"
        "fmt"
)

func main(){

        var netClient = &http.Client{
                Timeout: time.Second * 10,
        }

        g := gonion.Client{HttpClient: netClient}
        res := g.Details(gonion.Params{Lookup : "37CB803A9B74A7B7693040ED94E4AA9E66838021", Running: true, RecommendedVersion: true})
        fmt.Println(res)
}
```

## Installation

The Golang wrapper has been tested with Golang 1.6+. It may worker with older versions although it has not been tested.

To use it, just include it to your ``import`` and run ``go get``:
```bash
go get github.com/R4yGM/gonion
```

```go
import (
	...
	"github.com/R4yGM/gonion"
)
```
