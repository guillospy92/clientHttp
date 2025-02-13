package gohttp

import (
	"net/http"

	"github.com/guillospy92/clientHttp/core"
)

// ClientInterface interface contains the necessary methods to interact with the http client (GET,POST,PUT,DELETE,PATCH)
type ClientInterface interface {
	Get(url string, headers http.Header) (core.ResponseInterface, error)
	Post(url string, headers http.Header, body any) (core.ResponseInterface, error)
	Put(url string, headers http.Header, body any) (core.ResponseInterface, error)
	Delete(url string, headers http.Header, body any) (core.ResponseInterface, error)
	Patch(url string, headers http.Header, body any) (core.ResponseInterface, error)
}

// merClient contains a client of the go core where it is in charge of handling the request
// contains a builder that is represented by a clientBuilder which contains previous configurations for the client
type merClient struct {
	builder *clientBuilder
	client  *http.Client
}

// Get interacts with a request and brings a response type of structure that implements
// this interface ResponseInterface where the http method is GET
func (c *merClient) Get(url string, headers http.Header) (response core.ResponseInterface, err error) {
	return c.do(http.MethodGet, url, headers, nil)
}

// Post interacts with a request and brings a response type of structure that implements
// this interface ResponseInterface where the http method is POST
func (c *merClient) Post(url string, headers http.Header, body any) (response core.ResponseInterface, err error) {
	return c.do(http.MethodPost, url, headers, body)
}

// Put interacts with a request and brings a response type of structure that implements
// this interface ResponseInterface where the http method is PUT
func (c *merClient) Put(url string, headers http.Header, body any) (response core.ResponseInterface, err error) {
	return c.do(http.MethodPut, url, headers, body)
}

// Delete interacts with a request and brings a response type of structure that implements
// this interface ResponseInterface where the http method is DELETE
func (c *merClient) Delete(url string, headers http.Header, body any) (response core.ResponseInterface, err error) {
	return c.do(http.MethodDelete, url, headers, body)
}

// Patch interacts with a request and brings a response type of structure that implements
// this interface ResponseInterface where the http method is PATCH
func (c *merClient) Patch(url string, headers http.Header, body any) (response core.ResponseInterface, err error) {
	return c.do(http.MethodPatch, url, headers, body)
}
