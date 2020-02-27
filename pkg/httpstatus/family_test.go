package httpstatus

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFamily(t *testing.T) {
	tests := []struct {
		Name   string
		Code   int
		Family family
		Error  error
	}{
		{"Invalid - too low", 99, "", errors.New("blah")},
		{"Invalid - too high", 600, "", errors.New("Blah")},
		{"Switching Protocols", 101, "1xx", nil},
		{"Created", 201, "2xx", nil},
		{"Moved Permanently", 301, "3xx", nil},
		{"Unauthorized", 401, "4xx", nil},
		{"Not Implemented", 501, "5xx", nil},
	}
	for idx := range tests {
		test := tests[idx]
		t.Run(test.Name, func(t *testing.T) {
			fam, err := Family(test.Code)
			if test.Error != nil {
				assert.Error(t, err)
				//assert.EqualError(t, err, test.Error.Error())
				return
			}
			assert.NoError(t, err)
			assert.Equal(t, test.Family, fam)
		})
	}
}
