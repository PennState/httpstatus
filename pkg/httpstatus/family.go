package httpstatus

import (
	"fmt"
)

type family string

// Status families are grouped by their hundreds digit and are
// convenient for broad comparisons of returned status codes.
const (
	FamilyInvalid       family = ""
	FamilyInformational family = "1xx"
	FamilySuccessful    family = "2xx"
	FamilyRedirection   family = "3xx"
	FamilyClientError   family = "4xx"
	FamilyServerError   family = "5xx"
)

// Family returns the family "name" based on the provided HTTP
// status code.
func Family(code int) family {
	if code < 100 || code > 599 {
		return FamilyInvalid
	}
	return family(fmt.Sprintf("%dxx", code/100)) //, nil
}
