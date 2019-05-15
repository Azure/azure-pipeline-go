//+build darwin,unix,linux

package pipeline

import "net/http"

func proxyMiddleman() func(req *http.Request) (i *url.URL, e error) {
	return http.ProxyFromEnvironment
}
