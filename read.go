//go:build !js

package websocket

import (
	"bufio"
	"context"
	"io"
	"strings"
	"sync/atomic"

	"github.com/coder/websocket/internal/util"
)

// Reader reads from the connection until there is a WebSocket
// data message to be read. It will handle ping, pong and close frames as appropriate.
//
// It returns the type of the message and an io.Reader to read it.
// The passed context will also bound the reader.
// Ensure you read to EOF otherwise the connection will hang.
//
// Call CloseRead if you do not expect any data messages from the peer.
//
// Only one Reader may be open at a time.
//
// If you need a separate timeout on the Reader call and the Read itself,
// use time.AfterFunc to cancel the context passed in.
// See https://github.com/nhooyr/websocket/issues/87#issue-451703332
// Most users should not need this.
func (c *Conn) Reader(ctx context.Context) (MessageType, io.Reader, error) {
	_ = "STUB: not implemented"
	return *

	// Read is a convenience method around Reader to read a single message
	// from the connection.
	new(MessageType), *new(io.Reader), nil
}

func (c *Conn) Read(ctx context.Context) (MessageType, []byte, error) {
	_ = "STUB: not implemented"
	return *new(MessageType), nil, nil
}

// CloseRead starts a goroutine to read from the connection until it is closed
// or a data message is received.
//
// Once CloseRead is called you cannot read any messages from the connection.
// The returned context will be cancelled when the connection is closed.
//
// If a data message is received, the connection will be closed with StatusPolicyViolation.
//
// Call CloseRead when you do not expect to read any more messages.
// Since it actively reads from the connection, it will ensure that ping, pong and close
// frames are responded to. This means c.Ping and c.Close will still work as expected.
//
// This function is idempotent.
func (c *Conn) CloseRead(ctx context.Context) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

// SetReadLimit sets the max number of bytes to read for a single message.
// It applies to the Reader and Read methods.
//
// By default, the connection has a message read limit of 32768 bytes.
//
// When the limit is hit, reads return an error wrapping ErrMessageTooBig and
// the connection is closed with StatusMessageTooBig.
//
// Set to -1 to disable.
func (c *Conn) SetReadLimit(n int64) {
	_ = "STUB: not implemented"

	// We read one more byte than the limit in case
	// there is a fin frame that needs to be read.
	return
}

const defaultReadLimit = 32768

func newMsgReader(c *Conn) *msgReader { _ = "STUB: not implemented"; return nil }

func (mr *msgReader) resetFlate() { _ = "STUB: not implemented"; return }

func (mr *msgReader) putFlateReader() { _ = "STUB: not implemented"; return }

func (mr *msgReader) close() { _ = "STUB: not implemented"; return }

func (mr *msgReader) flateContextTakeover() bool { _ = "STUB: not implemented"; return false }

func (c *Conn) readRSV1Illegal(h header) bool {
	_ = "STUB: not implemented"
	// If compression is disabled, rsv1 is illegal.
	return false
}

// rsv1 is only allowed on data frames beginning messages.

func (c *Conn) readLoop(ctx context.Context) (header, error) {
	_ = "STUB: not implemented"
	return *new(header), nil
}

// Pass through CloseErrors when receiving a close frame.

// prepareRead sets the readTimeout context and returns a done function
// to be called after the read is done. It also returns an error if the
// connection is closed. The reference to the error is used to assign
// an error depending on if the connection closed or the context timed
// out during use. Typically, the referenced error is a named return
// variable of the function calling this method.
func (c *Conn) prepareRead(ctx context.Context, err *error) (func(), error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *Conn) readFrameHeader(ctx context.Context) (_ header, err error) {
	_ = "STUB: not implemented"
	return *new(header), nil
}

func (c *Conn) readFramePayload(ctx context.Context, p []byte) (_ int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Conn) handleControl(ctx context.Context, h header) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// opClose

// Only unlock readMu if this connection is being closed becaue
// c.close will try to acquire the readMu lock. We unlock for
// writeClose as well because it may also call c.close.

func (c *Conn) reader(ctx context.Context) (_ MessageType, _ io.Reader, err error) {
	_ = "STUB: not implemented"
	return *new(MessageType), *new(io.Reader), nil
}

type msgReader struct {
	c *Conn

	ctx         context.Context
	flate       bool
	flateReader io.Reader
	flateBufio  *bufio.Reader
	flateTail   strings.Reader
	limitReader *limitReader
	dict        *slidingWindow

	fin           bool
	payloadLength int64
	maskKey       uint32

	// util.ReaderFunc(mr.Read) to avoid continuous allocations.
	readFunc util.ReaderFunc
}

func (mr *msgReader) reset(ctx context.Context, h header) { _ = "STUB: not implemented"; return }

func (mr *msgReader) setFrame(h header) { _ = "STUB: not implemented"; return }

func (mr *msgReader) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (mr *msgReader) read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

type limitReader struct {
	c     *Conn
	r     io.Reader
	limit atomic.Int64
	n     int64
}

func newLimitReader(c *Conn, r io.Reader, limit int64) *limitReader {
	_ = "STUB: not implemented"
	return nil
}

func (lr *limitReader) reset(r io.Reader) { _ = "STUB: not implemented"; return }

func (lr *limitReader) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
