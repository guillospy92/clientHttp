package examples_test

import (
	"fmt"
	"github.com/guillospy92/clientHttp/gohttp"
	"github.com/guillospy92/clientHttp/gohttp/gohttp_mock"
	"net/http"
	"testing"
)

func TestGetServices(t *testing.T) {
	// start mock server
	gohttp_mock.StartMockServer()

	// initial mock response
	mockResponse := gohttp_mock.Mock{
		Method:     http.MethodGet,
		StatusCode: 200,
		Status:     "OK",
		Url:        "www.example.com",
		Response:   `{"name" : "test", "last_name" : "test"}`,
		Error:      nil,
	}

	// initial mock response
	mockResponse2 := gohttp_mock.Mock{
		Method:     http.MethodGet,
		StatusCode: 200,
		Status:     "OK",
		Url:        "www.example.com",
		Response:   `{"name" : "Guillermo", "last_name" : "Guillermo"}`,
		Error:      nil,
	}

	// add mock in map of the mock server
	gohttp_mock.AddMockServer(mockResponse)
	gohttp_mock.AddMockServer(mockResponse2)

	client := gohttp.NewClient().Build()

	response, err := client.Get(mockResponse.Url, nil)

	if err != nil {
		panic(err)
	}
	fmt.Println(response.GetString(), err, "guillermo rtomo")

	response2, err := client.Get(mockResponse.Url, nil)

	if err != nil {
		panic(err)
	}
	fmt.Println(response2.GetString(), err)

	// add mock in map of the mock server
	gohttp_mock.AddMockServer(mockResponse)

	// clean map mock server
	gohttp_mock.FlushMockServer()

	// stop mockServer
	gohttp_mock.StopMockServer()
}
