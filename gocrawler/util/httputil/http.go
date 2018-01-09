package httputil

import (
	"net"
	"net/http"
	"time"
)

func DoGet(url string) (resp *http.Response, err error) {
	return DoGetWithTimeout(url, time.Second*5)
}

func DoGetWithTimeout(url string, timeout time.Duration) (resp *http.Response, err error) {
	client := &http.Client{
		Transport: &http.Transport{
			Dial: func(netw, addr string) (net.Conn, error) {
				conn, err := net.DialTimeout(netw, addr, timeout)
				if err != nil {
					return nil, err
				}
				conn.SetDeadline(time.Now().Add(timeout))
				return conn, nil
			},
			ResponseHeaderTimeout: timeout,
		},
	}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}
	return res, nil
}
