package gonion

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
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

// GonionClient combines all that is needed for a Gonion client
// to work.
type GonionClient struct {
	client    HTTPClient
	UserAgent string
}

// NewGonionClient takes a HTTPClient and a User Agent string, and
// returns a new GonionClient.
func NewGonionClient(client HTTPClient, userAgent string) (*GonionClient, error) {
	if client == nil {
		return nil, ErrNilClient
	}
	return &GonionClient{
		client:    client,
		UserAgent: userAgent,
	}, nil
}

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

func (gc *GonionClient) getEndp(endp string, args Params, dst interface{}) error {
	// Create request
	req, _ := http.NewRequest(http.MethodGet, "https://onionoo.torproject.org/"+endp, nil)

	// Set headers
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", gc.UserAgent)

	// Encode query parameters
	q, err := args.QueryParams()
	if err != nil {
		return err
	}
	req.URL.RawQuery = q.Encode()

	// Issue request
	res, err := gc.client.Do(req)
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
