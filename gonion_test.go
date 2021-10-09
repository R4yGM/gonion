package gonion_test

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"reflect"
	"testing"

	"github.com/R4yGM/gonion"
	"github.com/stretchr/testify/assert"
)

var (
	errStrTypeOf = reflect.TypeOf(errors.New(""))
	errFake      = errors.New("this is a fake error")
)

func checkErr(err, expErr error, t *testing.T) {
	// Check err type
	typeErr := reflect.TypeOf(err)
	typeExpErr := reflect.TypeOf(expErr)
	if typeErr != typeExpErr {
		t.Fatalf("Failed to get expected error type: got \"%s\" instead of \"%s\".", typeErr, typeExpErr)
	}

	// Check Error content is not empty
	if err != nil && err.Error() == "" {
		t.Error("Error should not have an empty content.")
	}

	// Check if the error is generated using errors.New
	if typeErr == errStrTypeOf {
		if err.Error() != expErr.Error() {
			t.Errorf("Error message differs: got \"%s\" instead of \"%s\".", err, expErr)
		}
		return
	}

	assert := assert.New(t)
	switch err.(type) {
	case *json.SyntaxError:
		castedErr := err.(*json.SyntaxError)
		castedExpErr := expErr.(*json.SyntaxError)

		assert.Equal(castedExpErr.Offset, castedErr.Offset)

	case *gonion.ErrUnexpectedStatus:
		castedErr := err.(*gonion.ErrUnexpectedStatus)
		castedExpErr := expErr.(*gonion.ErrUnexpectedStatus)

		assert.Equal(castedExpErr.StatusCode, castedErr.StatusCode)
		assert.Equal(castedExpErr.Body, castedErr.Body)

	case nil:
		return

	default:
		t.Logf("\033[31mcheckErr Unsupported type: %s\033[0m\n", typeErr)
	}
}

// fakeHTTPClient is an implementation of HTTPClient that
// does nothing expect returning what you said it to.
type fakeHTTPClient struct {
	Response *http.Response
	Err      error
}

func (f fakeHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return f.Response, f.Err
}

var _ gonion.HTTPClient = (*fakeHTTPClient)(nil)

func newFakeHTTPClient(body string, statusCode int, err error) gonion.HTTPClient {
	return &fakeHTTPClient{
		Response: &http.Response{
			StatusCode: statusCode,
			Body:       newFakeReadCloser(body),
		},
		Err: err,
	}
}

// FakeReadCloser mocks an io.ReadCloser.
type fakeReadCloser struct {
	data      []byte
	readIndex int64
}

func (f *fakeReadCloser) Read(p []byte) (n int, err error) {
	if f.readIndex >= int64(len(f.data)) {
		err = io.EOF
		return
	}

	n = copy(p, f.data[f.readIndex:])
	f.readIndex += int64(n)
	return
}

func (f *fakeReadCloser) Close() error {
	return nil
}

var _ io.ReadCloser = (*fakeReadCloser)(nil)

func newFakeReadCloser(str string) *fakeReadCloser {
	return &fakeReadCloser{
		data: []byte(str),
	}
}

func str(str string) *string {
	return &str
}

func i(i int) *int {
	return &i
}

func f(f float64) *float64 {
	return &f
}
