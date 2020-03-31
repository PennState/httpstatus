package httperror

import (
	"fmt"
)

type HTTPError struct {
	Code        int
	Description string
	// TODO: Replace Code and Description with the line below
	// Status httpstatus.Status
	Body string
}

// func New(code int, body string) HTTPError {
// 	return HTTPError{
// 		Status: httpstatus.Of(code),
// 		Body:   body,
// 	}
// }

func (he HTTPError) Error() string {
	return fmt.Sprintf("%d (%s) - %s", he.Code, he.Description, he.Body)
}
