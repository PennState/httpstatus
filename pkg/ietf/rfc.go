package ietf

import (
	"fmt"
	"net/url"
	"strings"
)

// RFC represents a reference to an (adopted) IETF RFC
type RFC struct {
	Document string
	Section  string
}

// URL returns the URL associated with the RFC Number and Section.
func (rfc RFC) URL() *url.URL {
	num, sec := normalize(rfc.Document), normalize(rfc.Section)
	r := fmt.Sprintf("https://tools.ietf.org/html/%s#%s", num, sec)
	u, _ := url.Parse(r)
	return u
}

func normalize(s string) string {
	return url.PathEscape(strings.ToLower(strings.ReplaceAll(s, "Section ", "Section-")))
}
