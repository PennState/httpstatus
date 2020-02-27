package httpstatus

import (
	"fmt"
)

type family string

const (
	FamilyInformational family = "1xx"
	FamilySuccessful    family = "2xx"
	FamilyRedirection   family = "3xx"
	FamilyClientError   family = "4xx"
	FamilyServerError   family = "5xx"
)

func Family(code int) (family, error) {
	if code < 100 || code > 599 {
		return "", fmt.Errorf("not in the valid HTTP status code range: %d", code)
	}
	return family(fmt.Sprintf("%dxx", code/100)), nil
}
