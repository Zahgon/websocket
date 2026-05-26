//go:build !js

package websocket

import (
	"context"
	"net/http"
)

// AcceptOptions represents Accept's options.
type AcceptOptions struct {
	// Subprotocols lists the WebSocket subprotocols that Accept will negotiate with the client.
	// The empty subprotocol will always be negotiated as per RFC 6455. If you would like to
	// reject it, close the connection when c.Subprotocol() == "".
	Subprotocols []string

	// InsecureSkipVerify is used to disable Accept's origin verification behaviour.
	//
	// You probably want to use OriginPatterns instead.
	InsecureSkipVerify bool

	// OriginPatterns lists the host patterns for authorized origins.
	// The request host is always authorized.
	// Use this to enable cross origin WebSockets.
	//
	// i.e javascript running on example.com wants to access a WebSocket server at chat.example.com.
	// In such a case, example.com is the origin and chat.example.com is the request host.
	// One would set this field to []string{"example.com"} to authorize example.com to connect.
	//
	// Each pattern is matched case insensitively with path.Match (see
	// https://golang.org/pkg/path/#Match). By default, it is matched
	// against the request origin host. If the pattern contains a URI
	// scheme ("://"), it will be matched against "scheme://host".
	//
	// Please ensure you understand the ramifications of enabling this.
	// If used incorrectly your WebSocket server will be open to CSRF attacks.
	//
	// Do not use * as a pattern to allow any origin, prefer to use InsecureSkipVerify instead
	// to bring attention to the danger of such a setting.
	OriginPatterns []string

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

func (opts *AcceptOptions) cloneWithDefaults() *AcceptOptions {
	_ = "STUB: not implemented"
	return nil
}

// Accept accepts a WebSocket handshake from a client and upgrades the
// connection to a WebSocket.
//
// Accept will not allow cross origin requests by default.
// See the InsecureSkipVerify and OriginPatterns options to allow cross origin requests.
//
// Accept will write a response to w on all errors.
//
// Note that using the http.Request Context after Accept returns may lead to
// unexpected behavior (see http.Hijacker).
func Accept(w http.ResponseWriter, r *http.Request, opts *AcceptOptions) (*Conn, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func accept(w http.ResponseWriter, r *http.Request, opts *AcceptOptions) (_ *Conn, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// See https://github.com/nhooyr/websocket/issues/166

// https://github.com/golang/go/issues/32314

func verifyClientRequest(w http.ResponseWriter, r *http.Request) (errCode int, _ error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// The RFC states to remove any leading or trailing whitespace.

func authenticateOrigin(r *http.Request, originHosts []string) error {
	_ = "STUB: not implemented"
	return nil
}

func match(pattern, s string) (bool, error) { _ = "STUB: not implemented"; return false, nil }

func selectSubprotocol(r *http.Request, subprotocols []string) string {
	_ = "STUB: not implemented"
	return ""
}

func selectDeflate(extensions []websocketExtension, mode CompressionMode) (*compressionOptions, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// We used to implement x-webkit-deflate-frame too for Safari but Safari has bugs...
// See https://github.com/nhooyr/websocket/issues/218

func acceptDeflate(ext websocketExtension, mode CompressionMode) (*compressionOptions, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// We can't adjust the deflate window, but decoding with a larger window is acceptable.

func headerContainsTokenIgnoreCase(h http.Header, key, token string) bool {
	_ = "STUB: not implemented"
	return false
}

type websocketExtension struct {
	name   string
	params []string
}

func websocketExtensions(h http.Header) []websocketExtension { _ = "STUB: not implemented"; return nil }

func headerTokens(h http.Header, key string) []string { _ = "STUB: not implemented"; return nil }

var keyGUID = []byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11")

func secWebSocketAccept(secWebSocketKey string) string { _ = "STUB: not implemented"; return "" }
