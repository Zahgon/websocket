//go:build !js

package websocket

// StatusCode represents a WebSocket status code.
// https://tools.ietf.org/html/rfc6455#section-7.4
type StatusCode int

// https://www.iana.org/assignments/websocket/websocket.xhtml#close-code-number
//
// These are only the status codes defined by the protocol.
//
// You can define custom codes in the 3000-4999 range.
// The 3000-3999 range is reserved for use by libraries, frameworks and applications.
// The 4000-4999 range is reserved for private use.
const (
	StatusNormalClosure   StatusCode = 1000
	StatusGoingAway       StatusCode = 1001
	StatusProtocolError   StatusCode = 1002
	StatusUnsupportedData StatusCode = 1003

	// 1004 is reserved and so unexported.
	statusReserved StatusCode = 1004

	// StatusNoStatusRcvd cannot be sent in a close message.
	// It is reserved for when a close message is received without
	// a status code.
	StatusNoStatusRcvd StatusCode = 1005

	// StatusAbnormalClosure is exported for use only with Wasm.
	// In non Wasm Go, the returned error will indicate whether the
	// connection was closed abnormally.
	StatusAbnormalClosure StatusCode = 1006

	StatusInvalidFramePayloadData StatusCode = 1007
	StatusPolicyViolation         StatusCode = 1008
	StatusMessageTooBig           StatusCode = 1009
	StatusMandatoryExtension      StatusCode = 1010
	StatusInternalError           StatusCode = 1011
	StatusServiceRestart          StatusCode = 1012
	StatusTryAgainLater           StatusCode = 1013
	StatusBadGateway              StatusCode = 1014

	// StatusTLSHandshake is only exported for use with Wasm.
	// In non Wasm Go, the returned error will indicate whether there was
	// a TLS handshake failure.
	StatusTLSHandshake StatusCode = 1015
)

// CloseError is returned when the connection is closed with a status and reason.
//
// Use Go 1.13's errors.As to check for this error.
// Also see the CloseStatus helper.
type CloseError struct {
	Code   StatusCode
	Reason string
}

func (ce CloseError) Error() string { _ = "STUB: not implemented"; return "" }

// CloseStatus is a convenience wrapper around Go 1.13's errors.As to grab
// the status code from a CloseError.
//
// -1 will be returned if the passed error is nil or not a CloseError.
func CloseStatus(err error) StatusCode { _ = "STUB: not implemented"; return *new(StatusCode) }

// Close performs the WebSocket close handshake with the given status code and reason.
//
// It will write a WebSocket close frame with a timeout of 5s and then wait 5s for
// the peer to send a close frame.
// All data messages received from the peer during the close handshake will be discarded.
//
// The connection can only be closed once. Additional calls to Close
// are no-ops.
//
// The maximum length of reason must be 125 bytes. Avoid sending a dynamic reason.
//
// Close will unblock all goroutines interacting with the connection once
// complete.
func (c *Conn) Close(code StatusCode, reason string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// CloseNow closes the WebSocket connection without attempting a close handshake.
// Use when you do not want the overhead of the close handshake.
func (c *Conn) CloseNow() (err error) { _ = "STUB: not implemented"; return nil }

func (c *Conn) closeHandshake(code StatusCode, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) writeClose(code StatusCode, reason string) error {
	_ = "STUB: not implemented"
	return nil
}

// If the connection closed as we're writing we ignore the error as we might
// have written the close frame, the peer responded and then someone else read it
// and closed the connection.

func (c *Conn) waitCloseHandshake() error { _ = "STUB: not implemented"; return nil }

func (c *Conn) waitGoroutines() error { _ = "STUB: not implemented"; return nil }

func parseClosePayload(p []byte) (CloseError, error) {
	_ = "STUB: not implemented"
	return *new(CloseError), nil
}

// See http://www.iana.org/assignments/websocket/websocket.xhtml#close-code-number
// and https://tools.ietf.org/html/rfc6455#section-7.4.1
func validWireCloseCode(code StatusCode) bool { _ = "STUB: not implemented"; return false }

func (ce CloseError) bytes() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

const maxCloseReason = maxControlPayload - 2

func (ce CloseError) bytesErr() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (c *Conn) casClosing() bool { _ = "STUB: not implemented"; return false }

func (c *Conn) isClosed() bool { _ = "STUB: not implemented"; return false }
