package httptest

import "net/http"

// MockRoundTripper implements the http.RoundTripper interface as noted
// below but doesn't actually perform any network calls.  The Resp should
// be set up before each HTTP call.  The Req is available for inspection
// and/or assertions after the HTTP call.
type MockRoundTripper struct {
	Req  *http.Request
	Resp *http.Response
	Err  error
}

// RoundTrip implements the http.RoundTripper interface.
// https://pkg.go.dev/net/http?tab=doc#RoundTripper
func (mrt MockRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	mrt.Req = req
	return mrt.Resp, mrt.Err
}
