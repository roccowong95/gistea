package pkg

import (
	"net"
	"net/http"
	"time"
)

var (
	_defaultKeepAliveSec   = 10
	_defaultMaxIdleConns   = 10
	_defaultIdleTimeoutMin = 5
)

// TimeoutClient returns an *http.Client with given timeouts.
func TimeoutClient(ctimeout, totaltimeout time.Duration) *http.Client {
	tr := &http.Transport{
		MaxIdleConns: _defaultMaxIdleConns,
		DialContext: (&net.Dialer{
			// tcp keepalive interval
			KeepAlive: time.Duration(_defaultKeepAliveSec) * time.Second,
			// connect timeout
			Timeout: ctimeout,
		}).DialContext,
		// max idle time
		IdleConnTimeout: time.Duration(_defaultIdleTimeoutMin) * time.Minute,
	}
	ret := &http.Client{
		// max time before quitting a request
		Timeout:   totaltimeout,
		Transport: tr,
	}
	return ret
}
