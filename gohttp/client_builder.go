package gohttp

import (
	"net/http"
	"time"
)

const (
	// defaultMaxIdleConnection default maximum open requests for a client
	defaultMaxIdleConnection = 5

	// defaultResponseTimeOut maximum waiting time for a request to respond
	defaultResponseTimeOut = 5 * time.Second

	// defaultConnectionTimeOut time limit to establish a connection
	defaultConnectionTimeOut = 1 * time.Second

	// defaultTimeOut the maximum waiting time for a request to respond
	defaultTimeOut = 10 * time.Second
)

// clientBuilder structure that contains the configuration parameters to arm a client, all parameters are optional
type clientBuilder struct {
	Header            http.Header
	userAgent         string
	timeOut           time.Duration
	maxIdleConnection int
	connectionTimeOut time.Duration
	responseTimeOut   time.Duration
}

// ClientBuilderInterface contains the necessary methods to build a client
type ClientBuilderInterface interface {
	SetHeaders(headers http.Header) *clientBuilder
	SetTimeOut(timeOut time.Duration) *clientBuilder
	SetMaxIdleConnection(maxIdleConnection int) *clientBuilder
	SetConnectionTimeOut(timeOut time.Duration) *clientBuilder
	SetResponseTimeOut(responseTimeOut time.Duration) *clientBuilder
	SetUserAgent(userAgent string) *clientBuilder
	Build() ClientInterface
}

// Build exposes a merClient pointer previously to be able to interact with the methods (GET, POST, PUT, DELETE, PATCH)
func (cb *clientBuilder) Build() ClientInterface {
	return &merClient{
		builder: cb,
	}
}

// SetTimeOut set the maximum waiting time for a request to respond
func (cb *clientBuilder) SetTimeOut(timeOut time.Duration) *clientBuilder {
	cb.timeOut = timeOut
	return cb
}

// SetMaxIdleConnection set maximum open requests for a client
func (cb *clientBuilder) SetMaxIdleConnection(maxIdleConnection int) *clientBuilder {
	cb.maxIdleConnection = maxIdleConnection
	return cb
}

// SetConnectionTimeOut time limit to establish a connection
func (cb *clientBuilder) SetConnectionTimeOut(connectionTimeOut time.Duration) *clientBuilder {
	cb.connectionTimeOut = connectionTimeOut
	return cb
}

// SetResponseTimeOut set maximum time to read the response headers
func (cb *clientBuilder) SetResponseTimeOut(responseTimeOut time.Duration) *clientBuilder {
	cb.responseTimeOut = responseTimeOut
	return cb
}

// SetHeaders set the request headers before building the client
func (cb *clientBuilder) SetHeaders(headers http.Header) *clientBuilder {
	cb.Header = headers
	return cb
}

// SetUserAgent set user agent headers
func (cb *clientBuilder) SetUserAgent(userAgent string) *clientBuilder {
	cb.userAgent = userAgent
	return cb
}

// NewClient return new instance pointer of clientBuilder
func NewClient() ClientBuilderInterface {
	return &clientBuilder{}
}
