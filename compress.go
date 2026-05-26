//go:build !js

package websocket

import (
	"compress/flate"
	"io"
	"sync"
)

// CompressionMode represents the modes available to the permessage-deflate extension.
// See https://tools.ietf.org/html/rfc7692
//
// Works in all modern browsers except Safari which does not implement the permessage-deflate extension.
//
// Compression is only used if the peer supports the mode selected.
type CompressionMode int

const (
	// CompressionDisabled disables the negotiation of the permessage-deflate extension.
	//
	// This is the default. Do not enable compression without benchmarking for your particular use case first.
	CompressionDisabled CompressionMode = iota

	// CompressionContextTakeover compresses each message greater than 128 bytes reusing the 32 KB sliding window from
	// previous messages. i.e compression context across messages is preserved.
	//
	// As most WebSocket protocols are text based and repetitive, this compression mode can be very efficient.
	//
	// The memory overhead is a fixed 32 KB sliding window, a fixed 1.2 MB flate.Writer and a sync.Pool of 40 KB flate.Reader's
	// that are used when reading and then returned.
	//
	// Thus, it uses more memory than CompressionNoContextTakeover but compresses more efficiently.
	//
	// If the peer does not support CompressionContextTakeover then we will fall back to CompressionNoContextTakeover.
	CompressionContextTakeover

	// CompressionNoContextTakeover compresses each message greater than 512 bytes. Each message is compressed with
	// a new 1.2 MB flate.Writer pulled from a sync.Pool. Each message is read with a 40 KB flate.Reader pulled from
	// a sync.Pool.
	//
	// This means less efficient compression as the sliding window from previous messages will not be used but the
	// memory overhead will be lower as there will be no fixed cost for the flate.Writer nor the 32 KB sliding window.
	// Especially if the connections are long lived and seldom written to.
	//
	// Thus, it uses less memory than CompressionContextTakeover but compresses less efficiently.
	//
	// If the peer does not support CompressionNoContextTakeover then we will fall back to CompressionDisabled.
	CompressionNoContextTakeover
)

func (m CompressionMode) opts() *compressionOptions { _ = "STUB: not implemented"; return nil }

type compressionOptions struct {
	clientNoContextTakeover bool
	serverNoContextTakeover bool
}

func (copts *compressionOptions) String() string { _ = "STUB: not implemented"; return "" }

// These bytes are required to get flate.Reader to return.
// They are removed when sending to avoid the overhead as
// WebSocket framing tell's when the message has ended but then
// we need to add them back otherwise flate.Reader keeps
// trying to read more bytes.
const deflateMessageTail = "\x00\x00\xff\xff"

type trimLastFourBytesWriter struct {
	w    io.Writer
	tail []byte
}

func (tw *trimLastFourBytesWriter) reset() { _ = "STUB: not implemented"; return }

func (tw *trimLastFourBytesWriter) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Now we need to write as many extra bytes as we can from the previous tail.

// Shift remaining bytes in tail over.

// If p is less than or equal to 4 bytes,
// all of it is is part of the tail.

// Otherwise, only the last 4 bytes are.

var flateReaderPool sync.Pool

func getFlateReader(r io.Reader, dict []byte) io.Reader {
	_ = "STUB: not implemented"
	return *new(io.Reader)
}

func putFlateReader(fr io.Reader) { _ = "STUB: not implemented"; return }

var flateWriterPool sync.Pool

func getFlateWriter(w io.Writer) *flate.Writer { _ = "STUB: not implemented"; return nil }

func putFlateWriter(w *flate.Writer) { _ = "STUB: not implemented"; return }

type slidingWindow struct {
	buf []byte
}

var (
	swPoolMu sync.RWMutex
	swPool   = map[int]*sync.Pool{}
)

func slidingWindowPool(n int) *sync.Pool { _ = "STUB: not implemented"; return nil }

func (sw *slidingWindow) init(n int) {
	if sw.buf != nil {
		return
	}

	if n == 0 {
		n = 32768
	}

	p := slidingWindowPool(n)
	sw2, ok := p.Get().(*slidingWindow)
	if ok {
		*sw = *sw2
	} else {
		sw.buf = make([]byte, 0, n)
	}
}

func (sw *slidingWindow) close() { _ = "STUB: not implemented"; return }

func (sw *slidingWindow) write(p []byte) { _ = "STUB: not implemented"; return }

// We need to shift spaceNeeded bytes from the end to make room for p at the end.
