package ietf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURL(t *testing.T) {
	tests := []struct {
		name string
		inp  RFC
		exp  string
	}{
		{"Valid", RFC{Document: "RFC1234", Section: "Section 6.4.1"}, "https://tools.ietf.org/html/rfc1234#section-6.4.1"},
		{"Empty Number", RFC{Section: "Section 6.4.1"}, "https://tools.ietf.org/html/#section-6.4.1"},
		{"Empty Section", RFC{Document: "RFC1234"}, "https://tools.ietf.org/html/rfc1234"},
		{"Requires encoding", RFC{Document: "RFC1234", Section: "\"This is a test\""},
			"https://tools.ietf.org/html/rfc1234#%22this%20is%20a%20test%22"},
	}

	for idx := range tests {
		test := tests[idx]
		t.Run(test.name, func(t *testing.T) {
			act := test.inp.URL()
			assert.Equal(t, test.exp, act.String())
		})
	}
}
