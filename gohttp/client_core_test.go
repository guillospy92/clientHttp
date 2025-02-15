package gohttp

import (
	"encoding/xml"
	"net/http"
	"testing"
)

func TestRequestHeader(t *testing.T) {
	t.Parallel()
	clientBuilder := clientBuilder{}
	commonHeader := make(http.Header)
	commonHeader.Set("Content-Type", "application/json")
	commonHeader.Set("User-Agent", "cool-http-client")
	clientBuilder.SetHeaders(commonHeader)
	client := merClient{
		builder: &clientBuilder,
		client:  &http.Client{},
	}
	customHeader := make(http.Header)
	customHeader.Set("x-Request-Id", "Abc-123")
	headers := client.getRequestHeaders(customHeader)
	if len(headers) != 3 {
		t.Errorf("error we expect 2 headers")
	}

	if headers.Get("Content-Type") != "application/json" {
		t.Errorf("key content type not equal that %v", headers.Get("Content-Type"))
	}

	if headers.Get("User-Agent") != "cool-http-client" {
		t.Errorf("key user agent not equal that %v", headers.Get("User-Agent"))
	}

	if headers.Get("x-Request-Id") != "Abc-123" {
		t.Errorf("key x request Id type not equal that %v", headers.Get("x-Request-Id"))
	}
}

func TestRequestBody(t *testing.T) {
	client := merClient{}

	// request body application-json
	t.Run("Json", func(t *testing.T) {
		t.Parallel()
		body := struct {
			Name     string `json:"name"`
			LastName string `json:"last_name"`
		}{
			Name:     "test",
			LastName: "test",
		}

		resp, err := client.getRequestBody("application/json", body)

		if err != nil {
			t.Errorf("Error Marshal json %v", err)
		}

		if string(resp) != `{"name":"test","last_name":"test"}` {
			t.Errorf("Error json not equal to response %v", string(resp))
		}
	})

	// request body application-xml
	t.Run("xml", func(t *testing.T) {
		t.Parallel()
		body := struct {
			XMLName  xml.Name `xml:"rss"`
			Name     string   `xml:"name"`
			LastName string   `xml:"last_name"`
		}{
			Name:     "test",
			LastName: "test",
		}

		resp, err := client.getRequestBody("application/xml", &body)

		if err != nil {
			t.Errorf("Error Marshal xml %v", err)
		}

		if string(resp) != `<rss><name>test</name><last_name>test</last_name></rss>` {
			t.Errorf("Erros xml not equal to response %v", string(resp))
		}
	})

	// request body application-xml
	t.Run("Default", func(t *testing.T) {
		t.Parallel()
		body := []string{"body 1", "body 2"}
		resp, err := client.getRequestBody("", body)

		if err != nil {
			t.Errorf("Error Marshal json %v", err)
		}

		if string(resp) != `["body 1","body 2"]` {
			t.Errorf("Error json not equal to response %v", string(resp))
		}
	})

	// request body application-xml
	t.Run("nil", func(t *testing.T) {
		t.Parallel()
		resp, err := client.getRequestBody("", nil)

		if err != nil {
			t.Errorf("Error Marshal json %v", err)
		}

		if resp != nil {
			t.Errorf("Error json not equal to response %v", resp)
		}
	})

	t.Run("builder", func(t *testing.T) {
		t.Parallel()
		builder := NewClient().
			SetMaxIdleConnection(5).
			SetConnectionTimeOut(4).
			SetResponseTimeOut(3).
			SetTimeOut(4).SetUserAgent("user-Agent")
		client = merClient{
			builder: builder,
		}
		if client.getMaxIdleConnection() != builder.maxIdleConnection {
			t.Errorf("Error MaxIdleConnection ")
		}

		if client.getResponseTimeOut() != builder.responseTimeOut {
			t.Errorf("Error ResponseTimeOut")
		}

		if client.getConnectionTimeOut() != builder.timeOut {
			t.Errorf("Error timeOut")
		}
	})
}
