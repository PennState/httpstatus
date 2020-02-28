//go:generate httpstatus

// Package httpstatus expands the simple Status and StatusCode values
// returned by the net/http package to included classes, references
// as well as an HTTPError that provides more useful information
// from the http.Response object.
package httpstatus
