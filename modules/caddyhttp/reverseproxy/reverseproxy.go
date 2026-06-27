package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"time"
)

// ... existing code ...

func (r *ReverseProxy) getReverseProxy() *httputil.ReverseProxy {
	proxy := &httputil.ReverseProxy{
		Director: r.director,
		// ... existing fields ...
	}

	// Apply flush_interval
	if r.FlushInterval != nil {
		if *r.FlushInterval < 0 {
			proxy.FlushInterval = -1
		} else {
			proxy.FlushInterval = time.Duration(*r.FlushInterval) * time.Millisecond
		}
	} else {
		// Default to immediate flush for streaming content types if not specified
		proxy.FlushInterval = -1
	}

	return proxy
}