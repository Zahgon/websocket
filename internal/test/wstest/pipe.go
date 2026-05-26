//go:build !js

package wstest

import (
	"bufio"
	"net"
	"net/http"
	"net/http/httptest"

	"github.com/coder/websocket"
)

// Pipe is used to create an in memory connection
// between two websockets analogous to net.Pipe.
func Pipe(dialOpts *websocket.DialOptions, acceptOpts *websocket.AcceptOptions) (clientConn, serverConn *websocket.Conn) {
	_ = "STUB: not implemented"
	return nil, nil
}

type fakeTransport struct {
	h http.HandlerFunc
}

func (t fakeTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

type testHijacker struct {
	*httptest.ResponseRecorder
	serverConn net.Conn
}

var _ http.Hijacker = testHijacker{}

func (hj testHijacker) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	_ = "STUB: not implemented"
	return *new(net.Conn), nil, nil
}
