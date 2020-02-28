package httputil

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	emptyString      = ""
	testErrorMessage = "This is a test error"
	testString       = "This is a test"
)

func TestBody(t *testing.T) {
	tests := []struct {
		Name   string
		Body   io.ReadCloser
		Output string
	}{
		{"Erroring reader", erroringReadCloser{}, emptyString},
		{"Empty body", newStringReadCloser(emptyString), emptyString},
		{"Non-empty body", newStringReadCloser(testString), testString},
	}
	for idx := range tests {
		test := tests[idx]
		t.Run(test.Name, func(t *testing.T) {
			resp := http.Response{
				Body: test.Body,
			}
			assert.Equal(t, test.Output, BodyAsStringNoError(&resp))
		})
	}
}

func newStringReadCloser(s string) io.ReadCloser {
	return ioutil.NopCloser(strings.NewReader(s))
}

type erroringReadCloser struct{}

func (erc erroringReadCloser) Read(p []byte) (n int, err error) {
	return 0, errors.New(testErrorMessage)
}

func (erc erroringReadCloser) Close() error {
	return nil
}
