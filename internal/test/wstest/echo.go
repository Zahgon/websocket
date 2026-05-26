package wstest

import (
	"context"

	"github.com/coder/websocket"
)

// EchoLoop echos every msg received from c until an error
// occurs or the context expires.
// The read limit is set to 1 << 30.
func EchoLoop(ctx context.Context, c *websocket.Conn) error { _ = "STUB: not implemented"; return nil }

// Echo writes a message and ensures the same is sent back on c.
func Echo(ctx context.Context, c *websocket.Conn, max int) error {
	_ = "STUB: not implemented"
	return nil
}

func randMessage(typ websocket.MessageType, n int) []byte { _ = "STUB: not implemented"; return nil }
