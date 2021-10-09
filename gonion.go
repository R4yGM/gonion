package gonion

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"sync"

	"github.com/gorilla/schema"
)

var (
	// ErrNilClient is an error meaning the GonionClient tried to
	// get instanciated using a nil HTTPClient.
	ErrNilClient = errors.New("given client is nil")
)

// HTTPClient defines what a client have to implement.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

var _ HTTPClient = (*http.Client)(nil)

// ErrUnexpectedStatus is an error meaning the Onionoo server
// answered with an unexpected status code (something unexpected
// happened).
type ErrUnexpectedStatus struct {
	Body       []byte
	StatusCode int
}

func (e ErrUnexpectedStatus) Error() string {
	return fmt.Sprintf("unexpected status %d with body %s", e.StatusCode, e.Body)
}

var _ error = (*ErrUnexpectedStatus)(nil)

// CommaSepList is a slice that will get encoded by schema as
// a comma separated list.
type CommaSepList []string

func commaSepListEncode(val reflect.Value) string {
	builder := strings.Builder{}

	len := val.Len()
	strs := val.Slice(0, val.Len())
	for i := 0; i < len; i++ {
		index := strs.Index(i)
		s := index.String()
		builder.WriteString(s)

		if i != len-1 {
			builder.WriteString(",")
		}
	}

	return builder.String()
}

// Params represents a Onionioo parameters.
// See https://metrics.torproject.org/onionoo.html#parameters.
type Params struct {
	Type               *string       `schema:"type,omitempty"`
	Running            *bool         `schema:"running,omitempty"`
	Search             *string       `schema:"search,omitempty"`
	Lookup             *string       `schema:"lookup,omitempty"`
	Country            *string       `schema:"country,omitempty"`
	As                 *string       `schema:"as,omitempty"`
	AsName             *string       `schema:"as_name,omitempty"`
	Flag               *string       `schema:"flag,omitempty"`
	FirstSeenDays      *string       `schema:"first_seen_days,omitempty"`
	LastSeenDays       *string       `schema:"last_seen_days,omitempty"`
	Contact            *string       `schema:"contact,omitempty"`
	Family             *string       `schema:"family,omitempty"`
	Version            *string       `schema:"version,omitempty"`
	OS                 *string       `schema:"os,omitempty"`
	HostName           *string       `schema:"host_name,omitempty"`
	RecommendedVersion *bool         `schema:"recommended_version,omitempty"`
	Fields             *CommaSepList `schema:"fields,omitempty"`
	Order              *CommaSepList `schema:"order,omitempty"`
	Offset             *int          `schema:"offset,omitempty"`
	Limit              *int          `schema:"limit,omitempty"`
}

var encoderOnce = sync.Once{}
var encoder *schema.Encoder

func getEncoder() *schema.Encoder {
	encoderOnce.Do(func() {
		encoder = schema.NewEncoder()
		encoder.RegisterEncoder(CommaSepList{}, commaSepListEncode)
	})
	return encoder
}

func getEndp(client HTTPClient, endp string, params Params, dst interface{}) error {
	if client == nil {
		return ErrNilClient
	}

	// Create request
	req, _ := http.NewRequest(http.MethodGet, "https://onionoo.torproject.org/"+endp, nil)
	req.Header.Set("Accept", "application/json")
	q := url.Values{}
	_ = getEncoder().Encode(params, q)
	req.URL.RawQuery = q.Encode()

	// Issue request
	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	// Check status code
	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusNotModified {
		return &ErrUnexpectedStatus{
			Body:       body,
			StatusCode: res.StatusCode,
		}
	}

	// Unmarshal response
	err = json.Unmarshal(body, dst)
	if err != nil {
		return err
	}

	return nil
}
