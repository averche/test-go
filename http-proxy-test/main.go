package httpproxytest

import (
	"net/http"
	"net/url"
)

func main() {
	proxyAddr := "http://:8888"
	proxyURL, _ := url.Parse(proxyAddr)

	transport := &http.Transport{}
	transport.Proxy = http.ProxyURL(proxyURL)

	client := &http.Client{}
	client.Transport = transport

	// Since the proxy is explictly defined in the Transport, all requests,
	// including requests to `localhost`, `127.0.0.1`, and `::`, will use the
	// proxy.
	_, err := client.Get("http://127.0.0.1:8000")
	if err != nil {
		panic(err)
	}
}
