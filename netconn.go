package websocket

import (
	"context"
	"io"
	"net"
	"sync/atomic"
	"time"
)

// NetConn converts a *websocket.Conn into a net.Conn.
//
// It's for tunneling arbitrary protocols over WebSockets.
// Few users of the library will need this but it's tricky to implement
// correctly and so provided in the library.
// See https://github.com/nhooyr/websocket/issues/100.
//
// Every Write to the net.Conn will correspond to a message write of
// the given type on *websocket.Conn.
//
// The passed ctx bounds the lifetime of the net.Conn. If cancelled,
// all reads and writes on the net.Conn will be cancelled.
//
// If a message is read that is not of the correct type, the connection
// will be closed with StatusUnsupportedData and an error will be returned.
//
// Close will close the *websocket.Conn with StatusNormalClosure.
//
// When a deadline is hit and there is an active read or write goroutine, the
// connection will be closed. This is different from most net.Conn implementations
// where only the reading/writing goroutines are interrupted but the connection
// is kept alive.
//
// The Addr methods will return the real addresses for connections obtained
// from websocket.Accept. But for connections obtained from websocket.Dial, a mock net.Addr
// will be returned that gives "websocket" for Network() and "websocket/unknown-addr" for
// String(). This is because websocket.Dial only exposes a io.ReadWriteCloser instead of the
// full net.Conn to us.
//
// When running as WASM, the Addr methods will always return the mock address described above.
//
// A received StatusNormalClosure or StatusGoingAway close frame will be translated to
// io.EOF when reading.
//
// Furthermore, the ReadLimit is set to -1 to disable it.
func NetConn(ctx context.Context, c *Conn, msgType MessageType) net.Conn {
	_ = "STUB: not implemented"
	return *new(net.Conn)
}

// If the lock cannot be acquired, then there is an
// active write goroutine and so we should cancel the context.

// Prevents future writes from writing until the deadline is reset.

// If the lock cannot be acquired, then there is an
// active read goroutine and so we should cancel the context.

// Prevents future reads from reading until the deadline is reset.

type netConn struct {
	c       *Conn
	msgType MessageType

	writeTimer   *time.Timer
	writeMu      *mu
	writeExpired atomic.Int64
	writeCtx     context.Context
	writeCancel  context.CancelFunc

	readTimer   *time.Timer
	readMu      *mu
	readExpired atomic.Int64
	readCtx     context.Context
	readCancel  context.CancelFunc
	readEOFed   bool
	reader      io.Reader
}

var _ net.Conn = &netConn{}

func (nc *netConn) Close() error { _ = "STUB: not implemented"; return nil }

func (nc *netConn) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (nc *netConn) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (nc *netConn) read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type websocketAddr struct{}

func (a websocketAddr) Network() string { _ = "STUB: not implemented"; return "" }

func (a websocketAddr) String() string { _ = "STUB: not implemented"; return "" }

func (nc *netConn) SetDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (nc *netConn) SetWriteDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (nc *netConn) SetReadDeadline(t time.Time) error { _ = "STUB: not implemented"; return nil }
