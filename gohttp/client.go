package gohttp

import (
	"net/http"
)

type MerClientInterface interface {
	Get (url string, headers http.Header) (*http.Response, error)
	Post (url string, headers http.Header, body interface{}) (*http.Response, error)
	Put (url string, headers http.Header, body interface{}) (*http.Response, error)
	Delete (url string, headers http.Header, body interface{}) (*http.Response, error)
	Patch (url string, headers http.Header, body interface{}) (*http.Response, error)
}

type merClient struct {
	builder *clientBuilder
	client *http.Client
}

func (c *merClient) Get(url string, headers http.Header) (*http.Response, error) {
	return c.do(http.MethodGet, url, headers, nil)
}

func (c *merClient) Post(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPost, url, headers, body)
}

func (c *merClient) Put(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPut, url, headers, body)
}

func (c *merClient) Delete(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodDelete, url, headers, body)
}

func (c *merClient) Patch(url string, headers http.Header, body interface{}) (*http.Response, error) {
	return c.do(http.MethodPatch, url, headers, body)
}


