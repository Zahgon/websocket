//go:build !js

package websocket

import (
	"bufio"
)

// opcode represents a WebSocket opcode.
type opcode int

// https://tools.ietf.org/html/rfc6455#section-11.8.
const (
	opContinuation opcode = iota
	opText
	opBinary
	// 3 - 7 are reserved for further non-control frames.
	_
	_
	_
	_
	_
	opClose
	opPing
	opPong
	// 11-16 are reserved for further control frames.
)

// header represents a WebSocket frame header.
// See https://tools.ietf.org/html/rfc6455#section-5.2.
type header struct {
	fin    bool
	rsv1   bool
	rsv2   bool
	rsv3   bool
	opcode opcode

	payloadLength int64

	masked  bool
	maskKey uint32
}

// readFrameHeader reads a header from the reader.
// See https://tools.ietf.org/html/rfc6455#section-5.2.
func readFrameHeader(r *bufio.Reader, readBuf []byte) (h header, err error) {
	_ = "STUB: not implemented"
	return *new(header), nil
}

// maxControlPayload is the maximum length of a control frame payload.
// See https://tools.ietf.org/html/rfc6455#section-5.5.
const maxControlPayload = 125

// writeFrameHeader writes the bytes of the header to w.
// See https://tools.ietf.org/html/rfc6455#section-5.2
func writeFrameHeader(h header, w *bufio.Writer, buf []byte) (err error) {
	_ = "STUB: not implemented"
	return nil
}
