package gohttp

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
)

func (c *merClient) do(method string, url string, headers http.Header, body interface{}) (*http.Response, error) {
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

	return c.getHttpClient().Do(request)
}

func (c *merClient) getHttpClient() *http.Client {
	if c.client != nil {
		return c.client
	}

	c.client =  &http.Client{
		Timeout: c.getTimeOut(),
		Transport: &http.Transport{
			MaxIdleConnsPerHost: c.getMaxIdleConnection(), // maximum open connections
			ResponseHeaderTimeout: c.getResponseTimeOut(), // limit headers read time response
			DialContext: (&net.Dialer{
				Timeout: c.getConnectionTimeOut(),
			}).DialContext,
		},
	}
	return c.client
}

func (c *merClient) getMaxIdleConnection() int  {
	if c.builder.maxIdleConnection > 0 {
		return c.builder.maxIdleConnection
	}
	return defaultMaxIdleConnection
}

func (c *merClient) getResponseTimeOut() time.Duration {
	if c.builder.responseTimeOut > 0 {
		return c.builder.responseTimeOut
	}
	return defaultResponseTimeOut
}

func (c *merClient) getConnectionTimeOut() time.Duration {
	if c.builder.connectionTimeOut > 0 {
		return c.builder.connectionTimeOut
	}
	return defaultConnectionTimeOut
}

func (c *merClient) getTimeOut() time.Duration  {
	if c.builder.timeOut > 0 {
		return c.builder.timeOut
	}
	return defaultTimeOut
}

func (c *merClient) getRequestBody(contentType string, body interface{}) ([]byte, error) {
	if body == nil {
		return nil, nil
	}
	switch contentType {
	case strings.ToLower("application/json") :
		return json.Marshal(body)
	case  strings.ToLower("application/xml") :
		return xml.Marshal(body)
	default:
		return json.Marshal(body)
	}
}

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

	return result
}
