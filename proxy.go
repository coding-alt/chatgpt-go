package main

import (
	"errors"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func SetupProxy(proxyOption string) (*http.Transport, error) {
	if proxyOption == "DIRECT" {
		return &http.Transport{}, nil
	}

	parts := strings.Split(proxyOption, " ")
	if len(parts) != 2 {
		return nil, errors.New("invalid proxy option format")
	}

	// 拼接代理 URL
	proxyURL, err := url.Parse(parts[0] + "://" + parts[1])
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		ForceAttemptHTTP2:     true,
		MaxIdleConns:          100,
		MaxIdleConnsPerHost:   100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		ExpectContinueTimeout: 1 * time.Second,
	}

	return transport, nil
}
