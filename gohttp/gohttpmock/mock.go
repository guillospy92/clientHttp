package gohttpmock

import (
	"github.com/guillospy92/clientHttp/core"
)

// Mock gives a mock element for tests if the MockUpServer variable is active
type Mock struct {
	Error      error
	Method     string
	URL        string
	Response   string
	Status     string
	StatusCode int
}

// GetResponse answers a data of type core.Response the same as it would return a true answer
func (m *Mock) GetResponse() (core.ResponseInterface, error) {
	return &core.Response{
		StatusCode: m.StatusCode,
		Status:     m.Status,
		Body:       []byte(m.Response),
	}, nil
}
