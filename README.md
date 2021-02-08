# Go HTTP Client
una libreria simple para un cliente http en golang

## Installation

```bash
# Go Modules
require github.com/guillospy92/clientHttp
```

```
import "github.com/guillospy92/clientHttp"
```

## Usage
### import package
```
"github.com/guillospy92/clientHttp/gohttp"
```
### calls
gohttp.NewClient() return *clientBuilder

#### call get
```
client := gohttp.NewClient().Build()
response, err := client.Get(url string, headers http.Header)
response.GetStatusCode() return int
response.GetStatus() return string
response.GetByte() return []byte
response.GetString() return string of response
response.UnMarshal(data interface{}) error

```

#### call post
```
client := gohttp.NewClient().Build()
response, err := client.Post(url string, headers http.Header, body interface{})
response.GetStatusCode() return int
response.GetStatus() return string
response.GetByte() return []byte
response.GetString() return string of response
response.UnMarshal(data interface{}) error
```

#### call put
```
client := gohttp.NewClient().Build()
response, err := client.Put(url string, headers http.Header)
response.GetStatusCode() return int
response.GetStatus() return string
response.GetByte() return []byte
response.GetString() return string of response
response.UnMarshal(data interface{}) error
```

#### call Delete
```
client := gohttp.NewClient().Build()
response, err := client.Delete(url string, headers http.Header, body interface{})
response.GetStatusCode() return int
response.GetStatus() return string
response.GetByte() return []byte
response.GetString() return string of response
response.UnMarshal(data interface{}) error
```

#### call Path
```
client := gohttp.NewClient().Build()
response, err := client.Patch(url string, headers http.Header, body interface{})
response.GetStatusCode() return int
response.GetStatus() return string
response.GetByte() return []byte
response.GetString() return string of response
response.UnMarshal(data interface{}) error
```

## Custom parameter Builder
```
SetHeaders(headers http.Header) *clientBuilder
	SetTimeOut(timeOut time.Duration) *clientBuilder
	SetMaxIdleConnection(maxIdleConnection int) *clientBuilder
	SetConnectionTimeOut(timeOut time.Duration) *clientBuilder
	SetResponseTimeOut(responseTimeOut time.Duration) *clientBuilder
	SetUserAgent(userAgent string) *clientBuilder
```

## Testing

```
// start mock server
gohttp_mock.StartMockServer()

add mock server

create mock
mock := gohttp_mock.Mock{
    Method:     http.MethodPost,
    StatusCode: 200,
    Status:     "OK",
    Url:        "www.example.com",
    Response:   `{"name" : "test", "last_name" : "test"}`,
    Error:      nil,
},

// add mock
gohttp_mock.AddMockServer(mock)

// stop mock server
gohttp_mock.StopMockServer()

// execute client
client := gohttp.NewClient().Build()
response, err := client.Get(url string, headers http.Header)
the var response will return the same mock indicate
```


