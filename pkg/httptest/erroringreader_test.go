package httptest

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErroringReadCloser(t *testing.T) {
	tests := []struct {
		name string
		inp  ErroringReadCloser
	}{
		{
			name: "Error in reader",
			inp:  ErroringReadCloser{},
		},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			defer test.inp.Close()
			bytes := []byte{}
			cnt, err := test.inp.Read(bytes)
			if err != nil {
				assert.Equal(t, 0, cnt)
				return
			}
		})
	}
}
