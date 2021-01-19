package gohttp

import (
	"net/http"
	"time"
)

const (
	defaultMaxIdleConnection = 5
	defaultResponseTimeOut = 5 * time.Second
	defaultConnectionTimeOut = 1 * time.Second
	defaultTimeOut = 10 * time.Second
)

type clientBuilder struct {
	Header http.Header
	timeOut time.Duration
	maxIdleConnection int
	connectionTimeOut time.Duration
	responseTimeOut time.Duration
}

type ClientBuilderInterface interface {
	Build() *merClient
	SetHeaders(headers http.Header) *clientBuilder
	SetTimeOut(timeOut time.Duration) *clientBuilder
	SetMaxIdleConnection(maxIdleConnection int) *clientBuilder
	SetConnectionTimeOut(timeOut time.Duration) *clientBuilder
	SetResponseTimeOut(responseTimeOut time.Duration) *clientBuilder
}

func (cb *clientBuilder) Build() MerClientInterface {
	return &merClient{
		builder: cb,
	}
}

func (cb *clientBuilder) SetTimeOut(timeOut time.Duration)  *clientBuilder {
	cb.timeOut = timeOut
	return cb
}
func (cb *clientBuilder) SetMaxIdleConnection(maxIdleConnection int)  *clientBuilder {
	cb.maxIdleConnection = maxIdleConnection
	return cb
}

func (cb *clientBuilder) SetConnectionTimeOut(connectionTimeOut time.Duration)  *clientBuilder {
	cb.connectionTimeOut = connectionTimeOut
	return cb
}

func (cb *clientBuilder) SetResponseTimeOut(responseTimeOut time.Duration)  *clientBuilder {
	cb.responseTimeOut = responseTimeOut
	return cb
}

func (cb *clientBuilder) SetHeaders(headers http.Header)  *clientBuilder {
	cb.Header = headers
	return cb
}

func NewClient() *clientBuilder {
	return &clientBuilder{}
}


