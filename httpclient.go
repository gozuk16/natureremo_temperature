package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
)

func newRequest(m map[string]string) (*http.Request, error) {
	url := m["url"]

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set(m["headerKey"], m["headerValue"])

	dump, err := httputil.DumpRequestOut(req, true)
	fmt.Printf("%s", dump)
	if err != nil {
		log.Fatal("Error requesting dump")
	}

	return req, err
}

func getResponse(m map[string]string) (*http.Response, error) {
	req, err := newRequest(m)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	} else if res.StatusCode != 200 {
		return nil, fmt.Errorf("http status %d", res.StatusCode)
	}

	return res, err
}
