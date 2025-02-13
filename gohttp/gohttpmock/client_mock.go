package gohttpmock

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

// HTTPClientMock implements the HttpClient interface
type HTTPClientMock struct{}

// Do HTTPClientMock if the MockUpServer.Enabled looks for the mock corresponding to the list of mocks of the servers
// if the mock is not found, it returns an error saying that a corresponding mock was not found
func (*HTTPClientMock) Do(request *http.Request) (*http.Response, error) {
	requestBody, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	defer func(requestBody io.ReadCloser) {
		err := requestBody.Close()
		if err != nil {
			log.Printf("error closing request body %v", err)
		}
	}(requestBody)

	mock := GetMock()
	if mock != nil {
		response := http.Response{
			StatusCode:    mock.StatusCode,
			Body:          io.NopCloser(strings.NewReader(mock.Response)),
			ContentLength: int64(len(mock.Response)),
		}
		if mock.Error != nil {
			return nil, mock.Error
		}
		return &response, nil
	}
	mockNoFound := fmt.Sprintf("No Mock find method %s, body %s", request.Method, request.URL.String())

	panic(mockNoFound)
}
