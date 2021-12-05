package httphandler

import (
	"net/http"
	"path"
	"strings"
)

// CustomHTTPHandler implement http.Handler
// We need this custom handler to workaround a issue from ARM, where ARM send linked notification to us with // in url path.
// http library in golang redirect this kind of request to normalized path.
type CustomHTTPHandler struct {
	handler http.Handler
}

// compiler to check if CustomHTTPHandler implement http.Handler
var _ http.Handler = &CustomHTTPHandler{}

// NewCustomHTTPHandler create new CustomHTTPHandler
// a concrete http.Handler need here to process request, after the workaround.
func NewCustomHTTPHandler(handler http.Handler) *CustomHTTPHandler {
	return &CustomHTTPHandler{
		handler: handler,
	}
}

// ServeHTTP implement http.Handler
// workaround the ARM issue, then process request.
func (c *CustomHTTPHandler) ServeHTTP(httpwriter http.ResponseWriter, httpRequest *http.Request) {
	pathCleaned := cleanPath(httpRequest.URL.Path)

	// check if workaround required
	if pathCleaned != httpRequest.URL.Path {
		// the workaround is remove double forward slash in url path
		// and proceed the request.
		httpRequest.URL.Path = pathCleaned
	}

	c.handler.ServeHTTP(httpwriter, httpRequest)
}

// cleanPath returns the canonical path for p, eliminating . and .. elements.
// !! this is copied from /usr/local/go/src/net/http/server.go.
func cleanPath(p string) string {
	if p == "" {
		return "/"
	}
	if p[0] != '/' {
		p = "/" + p
	}
	np := path.Clean(p)
	// path.Clean removes trailing slash except for root;
	// put the trailing slash back if necessary.
	if p[len(p)-1] == '/' && np != "/" {
		// Fast path for common case of p being the string we want:
		if len(p) == len(np)+1 && strings.HasPrefix(p, np) {
			np = p
		} else {
			np += "/"
		}
	}
	return np
}
