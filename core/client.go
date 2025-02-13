package core

import "net/http"

// HTTPClient gives the do method for http operations
type HTTPClient interface {
	Do(request *http.Request) (*http.Response, error)
}
