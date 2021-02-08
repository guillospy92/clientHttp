package gohttp_mock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// HttpClientMock implements the HttpClient interface
type HttpClientMock struct{}

// HttpClientMock if the MockUpServer.Enabled looks for the mock corresponding to the list of mocks of the servers
// if the mock is not found, it returns an error saying that a corresponding mock was not found
func (c *HttpClientMock) Do(request *http.Request) (*http.Response, error) {
	requestBody, err := request.GetBody()
	if err != nil {
		return nil, err
	}
	defer requestBody.Close()

	mock := GetMock()
	if mock != nil {
		response := http.Response{
			StatusCode:    mock.StatusCode,
			Body:          ioutil.NopCloser(strings.NewReader(mock.Response)),
			ContentLength: int64(len(mock.Response)),
		}
		if mock.Error != nil {
			return nil, mock.Error
		}
		return &response, nil
	}
	mockNoFound := fmt.Sprintf("No Mock find method %s, body %s", request.Method, request.URL.String())

	fmt.Println(1111111111)
	panic(mockNoFound)
}
