// Package wsjson provides helpers for reading and writing JSON messages.
package wsjson // import "github.com/coder/websocket/wsjson"

import (
	"context"

	"github.com/coder/websocket"
)

// Read reads a JSON message from c into v.
// It will reuse buffers in between calls to avoid allocations.
func Read(ctx context.Context, c *websocket.Conn, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func read(ctx context.Context, c *websocket.Conn, v any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// Write writes the JSON message v to c.
// It will reuse buffers in between calls to avoid allocations.
func Write(ctx context.Context, c *websocket.Conn, v any) error {
	_ = "STUB: not implemented"
	return nil
}

func write(ctx context.Context, c *websocket.Conn, v any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// json.Marshal cannot reuse buffers between calls as it has to return
// a copy of the byte slice but Encoder does as it directly writes to w.
