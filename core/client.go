package core

import "net/http"

// HttpClient gives the do method for http operations
type HttpClient interface {
	Do(request *http.Request) (*http.Response, error)
}
