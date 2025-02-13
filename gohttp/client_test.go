package gohttp

import (
	"errors"

	"github.com/guillospy92/clientHttp/core"
	"github.com/guillospy92/clientHttp/gohttp/gohttpmock"

	"net/http"
	"testing"
)

// test method post
func TestPostPutDeletePatchServicePost(t *testing.T) {
	tests := []struct {
		args struct {
			url         string
			header      http.Header
			requestBody string
		}
		name         string
		wantResponse string
		requestBody  string
		mock         gohttpmock.Mock
		wantError    bool
	}{
		{
			name:         "Response Success",
			wantError:    false,
			wantResponse: `{"name" : "test", "last_name" : "test"}`,
			mock: gohttpmock.Mock{
				Method:     http.MethodPost,
				StatusCode: 200,
				Status:     "OK",
				URL:        "www.example.com",
				Response:   `{"name" : "test", "last_name" : "test"}`,
				Error:      nil,
			},
			args: struct {
				url         string
				header      http.Header
				requestBody string
			}{
				url:         "www.example.com",
				header:      make(http.Header),
				requestBody: `{"param" : "param", "id" : "id"}`,
			},
		},
		{
			name:         "Response Error",
			wantError:    true,
			wantResponse: `{"name" : "test", "last_name" : "test"}`,
			mock: gohttpmock.Mock{
				Method:     http.MethodPost,
				StatusCode: 200,
				Status:     "OK",
				URL:        "www.example.com",
				Response:   `{"name" : "test", "last_name" : "test"}`,
				Error:      errors.New("error response"),
			},
			args: struct {
				url         string
				header      http.Header
				requestBody string
			}{
				url:         "www.example.com",
				header:      make(http.Header),
				requestBody: `{"param" : "param", "id" : "id"}`,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			// start mock server
			gohttpmock.StartMockServer()

			client := NewClient().Build()

			methods := [5]string{"Get", "Post", "Put", "Delete", "Patch"}
			for _, method := range methods {
				// add mock in map of the mock server
				gohttpmock.AddMockServer(tt.mock)
				var response core.ResponseInterface
				var err error
				switch method {
				case "Get":
					response, err = client.Get(tt.args.url, tt.args.header)
				case "Post":
					response, err = client.Post(tt.args.url, tt.args.header, tt.requestBody)
				case "Put":
					response, err = client.Put(tt.args.url, tt.args.header, tt.requestBody)
				case "Delete":
					response, err = client.Delete(tt.args.url, tt.args.header, tt.requestBody)
				case "Patch":
					response, err = client.Patch(tt.args.url, tt.args.header, tt.requestBody)
				}

				// if in test errorWant is false the client no should return err
				if !tt.wantError && err != nil {
					t.Errorf("test name %v want response %v", tt.name, tt.wantResponse)
				}

				// test response if not nil
				if response != nil && response.GetString() != tt.wantResponse {
					t.Errorf("test name %v want response %v", tt.name, tt.wantResponse)
				}

			}
			gohttpmock.StopMockServer()
		})
	}
}
