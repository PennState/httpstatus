package httpstatus

import (
	"fmt"
)

type class string
type Class class // TODO: this is awkward

// Status classes are grouped by their hundreds digit and are convenient
// for broad comparisons of returned status codes.  The status classes
// provided below are defined in section 6 of RFC7231
// (https://tools.ietf.org/html/rfc7231#section-6).
const (
	ClassInvalid       class = ""
	ClassInformational class = "1xx"
	ClassSuccessful    class = "2xx"
	ClassRedirection   class = "3xx"
	ClassClientError   class = "4xx"
	ClassServerError   class = "5xx"
)

// Family returns the family "name" based on the provided HTTP status
// code.
// TODO: This will be calculated per
func Family(code int) Class {
	if code < 100 || code > 599 {
		return Class(ClassInvalid)
	}
	return Class(class(fmt.Sprintf("%dxx", code/100)))
}
