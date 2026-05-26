//go:build !js

package websocket

import (
	"bufio"
	"compress/flate"
	"context"
	"io"
)

// Writer returns a writer bounded by the context that will write
// a WebSocket message of type dataType to the connection.
//
// You must close the writer once you have written the entire message.
//
// Only one writer can be open at a time, multiple calls will block until the previous writer
// is closed.
func (c *Conn) Writer(ctx context.Context, typ MessageType) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// Write writes a message to the connection.
//
// See the Writer method if you want to stream a message.
//
// If compression is disabled or the compression threshold is not met, then it
// will write the message in a single frame.
func (c *Conn) Write(ctx context.Context, typ MessageType, p []byte) error {
	_ = "STUB: not implemented"
	return nil
}

type msgWriter struct {
	c *Conn

	mu      *mu
	writeMu *mu
	closed  bool

	ctx    context.Context
	opcode opcode
	flate  bool

	trimWriter  *trimLastFourBytesWriter
	flateWriter *flate.Writer
}

func newMsgWriter(c *Conn) *msgWriter { _ = "STUB: not implemented"; return nil }

func (mw *msgWriter) ensureFlate() { _ = "STUB: not implemented"; return }

func (mw *msgWriter) flateContextTakeover() bool { _ = "STUB: not implemented"; return false }

func (c *Conn) writer(ctx context.Context, typ MessageType) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

func (c *Conn) write(ctx context.Context, typ MessageType, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mw *msgWriter) reset(ctx context.Context, typ MessageType) error {
	_ = "STUB: not implemented"
	return nil
}

func (mw *msgWriter) putFlateWriter() { _ = "STUB: not implemented"; return }

// writeCompressedFrame compresses and writes p as a single frame.
func (mw *msgWriter) writeCompressedFrame(ctx context.Context, p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Buffer compressed output so we can write as
// a single frame instead of chunked frames.

// Write writes the given bytes to the WebSocket connection.
func (mw *msgWriter) Write(p []byte) (_ int, err error) { _ = "STUB: not implemented"; return 0, nil }

// Only enables flate if the length crosses the
// threshold on the first frame

func (mw *msgWriter) write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// Close flushes the frame to the connection.
func (mw *msgWriter) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (mw *msgWriter) close() { _ = "STUB: not implemented"; return }

func (c *Conn) writeControl(ctx context.Context, opcode opcode, p []byte) error {
	_ = "STUB: not implemented"
	return nil
}

// writeFrame handles all writes to the connection.
func (c *Conn) writeFrame(ctx context.Context, fin bool, flate bool, opcode opcode, p []byte) (_ int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *Conn) writeFramePayload(p []byte) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// If the buffer is full, we need to flush.

// Start of next write in the buffer.

// extractBufioWriterBuf grabs the []byte backing a *bufio.Writer
// and returns it.
func extractBufioWriterBuf(bw *bufio.Writer, w io.Writer) []byte {
	_ = "STUB: not implemented"
	return nil
}

func (c *Conn) writeError(code StatusCode, err error) { _ = "STUB: not implemented"; return }
