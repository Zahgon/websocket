package main

import (
	"net/http"

	"golang.org/x/time/rate"

	"github.com/coder/websocket"
)

// echoServer is the WebSocket echo server implementation.
// It ensures the client speaks the echo subprotocol and
// only allows one message every 100ms with a 10 message burst.
type echoServer struct {
	// logf controls where logs are sent.
	logf func(f string, v ...any)
}

func (s echoServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}

// echo reads from the WebSocket connection and then writes
// the received message back to it.
// The entire function has 10s to complete.
func echo(c *websocket.Conn, l *rate.Limiter) error { _ = "STUB: not implemented"; return nil }
