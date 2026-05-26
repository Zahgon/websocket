//go:build !js

package websocket

import (
	"bufio"
	"context"
	"io"
	"sync"
	"sync/atomic"
)

// MessageType represents the type of a WebSocket message.
// See https://tools.ietf.org/html/rfc6455#section-5.6
type MessageType int

// MessageType constants.
const (
	// MessageText is for UTF-8 encoded text messages like JSON.
	MessageText MessageType = iota + 1
	// MessageBinary is for binary messages like protobufs.
	MessageBinary
)

// Conn represents a WebSocket connection.
// All methods may be called concurrently except for Reader and Read.
//
// You must always read from the connection. Otherwise control
// frames will not be handled. See Reader and CloseRead.
//
// Be sure to call Close on the connection when you
// are finished with it to release associated resources.
//
// On any error from any method, the connection is closed
// with an appropriate reason.
//
// This applies to context expirations as well unfortunately.
// See https://github.com/nhooyr/websocket/issues/242#issuecomment-633182220
type Conn struct {
	noCopy noCopy

	subprotocol    string
	rwc            io.ReadWriteCloser
	client         bool
	copts          *compressionOptions
	flateThreshold int
	br             *bufio.Reader
	bw             *bufio.Writer

	readTimeoutStop  atomic.Pointer[func() bool]
	writeTimeoutStop atomic.Pointer[func() bool]

	// Read state.
	readMu         *mu
	readHeaderBuf  [8]byte
	readControlBuf [maxControlPayload]byte
	msgReader      *msgReader

	// Write state.
	msgWriter      *msgWriter
	writeFrameMu   *mu
	writeBuf       []byte
	writeHeaderBuf [8]byte
	writeHeader    header

	// Close handshake state.
	closeStateMu     sync.RWMutex
	closeReceivedErr error
	closeSentErr     error

	// CloseRead state.
	closeReadMu   sync.Mutex
	closeReadCtx  context.Context
	closeReadDone chan struct{}

	closing atomic.Bool
	closeMu sync.Mutex // Protects following.
	closed  chan struct{}

	pingCounter    atomic.Int64
	activePingsMu  sync.Mutex
	activePings    map[string]chan<- struct{}
	onPingReceived func(context.Context, []byte) bool
	onPongReceived func(context.Context, []byte)
}

type connConfig struct {
	subprotocol    string
	rwc            io.ReadWriteCloser
	client         bool
	copts          *compressionOptions
	flateThreshold int
	onPingReceived func(context.Context, []byte) bool
	onPongReceived func(context.Context, []byte)

	br *bufio.Reader
	bw *bufio.Writer
}

func newConn(cfg connConfig) *Conn { _ = "STUB: not implemented"; return nil }

// Subprotocol returns the negotiated subprotocol.
// An empty string means the default protocol.
func (c *Conn) Subprotocol() string { _ = "STUB: not implemented"; return "" }

func (c *Conn) close() error { _ = "STUB: not implemented"; return nil }

// Have to close after c.closed is closed to ensure any goroutine that wakes up
// from the connection being closed also sees that c.closed is closed and returns
// closeErr.

// With the close of rwc, these become safe to close.

func (c *Conn) setupWriteTimeout(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Conn) clearWriteTimeout() { _ = "STUB: not implemented"; return }

func (c *Conn) setupReadTimeout(ctx context.Context) { _ = "STUB: not implemented"; return }

func (c *Conn) clearReadTimeout() { _ = "STUB: not implemented"; return }

func swapTimeoutStop(p *atomic.Pointer[func() bool], newStop *func() bool) {
	_ = "STUB: not implemented"
	return
}

func (c *Conn) flate() bool { _ = "STUB: not implemented"; return false }

// Ping sends a ping to the peer and waits for a pong.
// Use this to measure latency or ensure the peer is responsive.
// Ping must be called concurrently with Reader as it does
// not read from the connection but instead waits for a Reader call
// to read the pong.
//
// TCP Keepalives should suffice for most use cases.
func (c *Conn) Ping(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (c *Conn) ping(ctx context.Context, p string) error { _ = "STUB: not implemented"; return nil }

type mu struct {
	c  *Conn
	ch chan struct{}
}

func newMu(c *Conn) *mu { _ = "STUB: not implemented"; return nil }

func (m *mu) forceLock() { _ = "STUB: not implemented"; return }

func (m *mu) tryLock() bool { _ = "STUB: not implemented"; return false }

func (m *mu) lock(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// To make sure the connection is certainly alive.
// As it's possible the send on m.ch was selected
// over the receive on closed.

// Make sure to release.

func (m *mu) unlock() { _ = "STUB: not implemented"; return }

type noCopy struct{}

func (*noCopy) Lock() { _ = "STUB: not implemented"; return }
