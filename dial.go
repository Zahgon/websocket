//go:build !js

package websocket

import (
	"bufio"
	"context"
	"io"
	"net/http"
	"sync"
)

// DialOptions represents Dial's options.
type DialOptions struct {
	// HTTPClient is used for the connection.
	// Its Transport must return writable bodies for WebSocket handshakes.
	// http.Transport does beginning with Go 1.12.
	HTTPClient *http.Client

	// HTTPHeader specifies the HTTP headers included in the handshake request.
	HTTPHeader http.Header

	// Host optionally overrides the Host HTTP header to send. If empty, the value
	// of URL.Host will be used.
	Host string

	// Subprotocols lists the WebSocket subprotocols to negotiate with the server.
	Subprotocols []string

	// CompressionMode controls the compression mode.
	// Defaults to CompressionDisabled.
	//
	// See docs on CompressionMode for details.
	CompressionMode CompressionMode

	// CompressionThreshold controls the minimum size of a message before compression is applied.
	//
	// Defaults to 512 bytes for CompressionNoContextTakeover and 128 bytes
	// for CompressionContextTakeover.
	CompressionThreshold int

	// OnPingReceived is an optional callback invoked synchronously when a ping frame is received.
	//
	// The payload contains the application data of the ping frame.
	// If the callback returns false, the subsequent pong frame will not be sent.
	// To avoid blocking, any expensive processing should be performed asynchronously using a goroutine.
	OnPingReceived func(ctx context.Context, payload []byte) bool

	// OnPongReceived is an optional callback invoked synchronously when a pong frame is received.
	//
	// The payload contains the application data of the pong frame.
	// To avoid blocking, any expensive processing should be performed asynchronously using a goroutine.
	//
	// Unlike OnPingReceived, this callback does not return a value because a pong frame
	// is a response to a ping and does not trigger any further frame transmission.
	OnPongReceived func(ctx context.Context, payload []byte)
}

func (opts *DialOptions) cloneWithDefaults(ctx context.Context) (context.Context, context.CancelFunc, *DialOptions) {
	_ = "STUB: not implemented"
	return *new(context.Context), *new(context.CancelFunc), nil
}

// Dial performs a WebSocket handshake on url.
//
// The response is the WebSocket handshake response from the server.
// You never need to close resp.Body yourself.
//
// If an error occurs, the returned response may be non nil.
// However, you can only read the first 1024 bytes of the body.
//
// This function requires at least Go 1.12 as it uses a new feature
// in net/http to perform WebSocket handshakes.
// See docs on the HTTPClient option and https://github.com/golang/go/issues/26937#issuecomment-415855861
//
// URLs with http/https schemes will work and are interpreted as ws/wss.
func Dial(ctx context.Context, u string, opts *DialOptions) (*Conn, *http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func dial(ctx context.Context, urls string, opts *DialOptions, rand io.Reader) (_ *Conn, _ *http.Response, err error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// We read a bit of the body for easier debugging.

func handshakeRequest(ctx context.Context, urls string, opts *DialOptions, copts *compressionOptions, secWebSocketKey string) (*http.Response, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func secWebSocketKey(rr io.Reader) (string, error) { _ = "STUB: not implemented"; return "", nil }

func verifyServerResponse(opts *DialOptions, copts *compressionOptions, secWebSocketKey string, resp *http.Response) (*compressionOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func verifySubprotocol(subprotos []string, resp *http.Response) error {
	_ = "STUB: not implemented"
	return nil
}

func verifyServerExtensions(copts *compressionOptions, h http.Header) (*compressionOptions, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// We can't adjust the deflate window, but decoding with a larger window is acceptable.

var bufioReaderPool sync.Pool

func getBufioReader(r io.Reader) *bufio.Reader { _ = "STUB: not implemented"; return nil }

func putBufioReader(br *bufio.Reader) { _ = "STUB: not implemented"; return }

var bufioWriterPool sync.Pool

func getBufioWriter(w io.Writer) *bufio.Writer { _ = "STUB: not implemented"; return nil }

func putBufioWriter(bw *bufio.Writer) { _ = "STUB: not implemented"; return }
