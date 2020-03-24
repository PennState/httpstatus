package httptest

// ErroringReadCloser provides an implementation that errors to be
// purposely added to both the reading and closing methods.  If neither
// of the errors are included, then this reader simply returns a zero
// byte read.
type ErroringReadCloser struct {
	ReadError  error
	CloseError error
}

// Read implements the io.Reader interface
// https://golang.org/pkg/io/#Reader
func (erc ErroringReadCloser) Read(p []byte) (n int, err error) {
	return 0, erc.ReadError
}

// Close implements the io.Closer interface
// https://golang.org/pkg/io/#Closer
func (erc ErroringReadCloser) Close() error {
	return erc.CloseError
}
