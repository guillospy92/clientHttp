package core

import (
	"encoding/json"
	"net/http"
)

// ResponseInterface interface iterate with library
type ResponseInterface interface {
	GetHeader() http.Header
	GetStatusCode() int
	GetStatus() string
	GetByte() []byte
	GetString() string
	UnMarshal(data any) error
}

// Response abstraction of data from * http.Response
type Response struct {
	Headers    http.Header
	Status     string
	Body       []byte
	StatusCode int
}

// GetHeader get header of response
func (r *Response) GetHeader() http.Header {
	return r.Headers
}

// GetStatusCode get status code of response
func (r *Response) GetStatusCode() int {
	return r.StatusCode
}

// GetStatus get status of response
func (r *Response) GetStatus() string {
	return r.Status
}

// GetByte get response in byte
func (r *Response) GetByte() []byte {
	return r.Body
}

// GetString return response in string
func (r *Response) GetString() string {
	return string(r.Body)
}

// UnMarshal convert data type any in json struct
func (r *Response) UnMarshal(data any) error {
	return json.Unmarshal(r.Body, data)
}
