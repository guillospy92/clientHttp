package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"io"
	"log"

	"github.com/guillospy92/clientHttp/core"
	"github.com/guillospy92/clientHttp/gohttp/gohttpmock"

	"net"
	"net/http"
	"strings"
	"time"
)

// do trigger to make the request POST, PUT, DELETE, PATCH
func (c *merClient) do(method string, url string, headers http.Header, body any) (core.ResponseInterface, error) {
	fullHeaders := c.getRequestHeaders(headers)

	requestBody, err := c.getRequestBody(fullHeaders.Get("Content-type"), body)
	if err != nil {
		return nil, errors.New("error Unable for request body *****")
	}

	request, err := http.NewRequest(method, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.New("error Unable for request headers *****")
	}

	request.Header = fullHeaders

	response, err := c.getHTTPClient().Do(request)
	if err != nil {
		return nil, err
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			log.Printf("error closing request body %v", err)
		}
	}(response.Body)

	responseBody, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}
	return &core.Response{
		StatusCode: response.StatusCode,
		Status:     response.Status,
		Headers:    response.Header,
		Body:       responseBody,
	}, nil
}

// getMaxIdleConnection get value of maximum open requests for a client
func (c *merClient) getMaxIdleConnection() int {
	if c.builder.maxIdleConnection > 0 {
		return c.builder.maxIdleConnection
	}
	return defaultMaxIdleConnection
}

// getResponseTimeOut get value of maximum waiting time for a request to respond
func (c *merClient) getResponseTimeOut() time.Duration {
	if c.builder.responseTimeOut > 0 {
		return c.builder.responseTimeOut
	}
	return defaultResponseTimeOut
}

// getConnectionTimeOut get value of time limit to establish a connection
func (c *merClient) getConnectionTimeOut() time.Duration {
	if c.builder.connectionTimeOut > 0 {
		return c.builder.connectionTimeOut
	}
	return defaultConnectionTimeOut
}

// getTimeOut get value of maximum waiting time for a request to respond
func (c *merClient) getTimeOut() time.Duration {
	if c.builder.timeOut > 0 {
		return c.builder.timeOut
	}
	return defaultTimeOut
}

// getRequestBody transform the body according to the application / json key of the headers
func (*merClient) getRequestBody(contentType string, body any) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch contentType {
	case strings.ToLower("application/json"):
		return json.Marshal(body)
	case strings.ToLower("application/xml"):
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

// getRequestHeaders get the headers set by the builder to create the request
func (c *merClient) getRequestHeaders(requestHeaders http.Header) http.Header {
	result := make(http.Header)

	// Add common headers to the request
	for header, value := range c.builder.Header {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// Add custom headers to the request
	for header, value := range requestHeaders {
		if len(value) > 0 {
			result.Set(header, value[0])
		}
	}

	// if attribute user-Agent in builder is different ""
	if c.builder.userAgent != "" {
		if result.Get("User-Agent") != "" {
			return result
		}
		result.Set("User-Agent", c.builder.userAgent)
	}

	return result
}

// getHTTPClient returns a http client singleton with the previous builder configurations
func (c *merClient) getHTTPClient() core.HTTPClient {
	// if mock is active in testing
	if gohttpmock.MockUpServer.Enabled {
		return &gohttpmock.HTTPClientMock{}
	}
	// client singleton if is initialized
	if c.client != nil {
		return c.client
	}
	// if client not is initialized return new client
	c.client = &http.Client{
		Timeout: c.getTimeOut(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   c.getMaxIdleConnection(),
			ResponseHeaderTimeout: c.getResponseTimeOut(),
			DialContext: (&net.Dialer{
				Timeout: c.getConnectionTimeOut(),
			}).DialContext,
		},
	}
	return c.client
}
