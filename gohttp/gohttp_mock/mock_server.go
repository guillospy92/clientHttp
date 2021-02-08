package gohttp_mock

// MockUpServer global variable that lets you know in what state the mockery system is
var (
	MockUpServer = mockServer{
		mocks: []*Mock{},
	}
)

// mockserver provides the necessary elements to mock httpclient
// saves all the mock elements in its mocks attribute every time the getMock method is called this removes
// the first element of the array to avoid duplicate responses
// If we have a use case where more than one request is made, we must add the necessary mocks to supply them,
// since in the mock stack every time a mock is called, it removes the first element of the stack.
// If a mock is not found in the mock pile, it will launch a panic
type mockServer struct {
	// indicates that the mock system is active
	Enabled bool
	// pre-populated mock slices from AddMockServer method
	mocks []*Mock
}

// StartMockServer it tells the http client core that the mock is activated therefore all the requests made in the Do will go through the mock_client
func StartMockServer() {
	MockUpServer.Enabled = true
}

// StopMockServer it tells the http client core that the mock is stopped, therefore all requests made in the Do will go through the real client
func StopMockServer() {
	MockUpServer.Enabled = false
}

func AddMockServer(mock Mock) {
	MockUpServer.mocks = append(MockUpServer.mocks, &mock)
}

// gets the first item from the mock stack and removes it in turn
func GetMock() *Mock {
	mock := MockUpServer.mocks[0]
	MockUpServer.mocks = MockUpServer.mocks[1:]
	return mock
}
