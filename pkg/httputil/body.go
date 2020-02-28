package httputil

import (
	"io/ioutil"
	"net/http"
)

// BodyAsBytes returns the body contents of an HTTP response as a []byte
// or an error if there's a problem reading the data.
func BodyAsBytes(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}
	return bytes, nil
}

// BodyAsBytesNoError returns the body contents of an HTTP response as a
// []byte or an empty []byte if there's a problem reading the data.
func BodyAsBytesNoError(resp *http.Response) []byte {
	bytes, _ := BodyAsBytes(resp)
	return bytes
}

// BodyAsString returns the body contents of an HTTP response as a string
// or an error if theres a problem reading the data.
func BodyAsString(resp *http.Response) (string, error) {
	bytes, err := BodyAsBytes(resp)
	return string(bytes), err
}

// BodyAsStringNoError returns the body contents of an HTTP response as
// a string or an empty string if there's a problem reading the data.
func BodyAsStringNoError(resp *http.Response) string {
	str, _ := BodyAsString(resp)
	return str
}
