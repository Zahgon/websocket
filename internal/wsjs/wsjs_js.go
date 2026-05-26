//go:build js
// +build js

// Package wsjs implements typed access to the browser javascript WebSocket API.
//
// https://developer.mozilla.org/en-US/docs/Web/API/WebSocket
package wsjs

import (
	"syscall/js"
)

func handleJSError(err *error, onErr func()) { _ = "STUB: not implemented"; return }

// New is a wrapper around the javascript WebSocket constructor.
func New(url string, protocols []string) (c WebSocket, err error) {
	_ = "STUB: not implemented"
	return *new(WebSocket), nil
}

// WebSocket is a wrapper around a javascript WebSocket object.
type WebSocket struct {
	v js.Value
}

func (c WebSocket) setBinaryType(typ string) { _ = "STUB: not implemented"; return }

func (c WebSocket) addEventListener(eventType string, fn func(e js.Value)) func() {
	_ = "STUB: not implemented"
	return nil
}

// CloseEvent is the type passed to a WebSocket close handler.
type CloseEvent struct {
	Code     uint16
	Reason   string
	WasClean bool
}

// OnClose registers a function to be called when the WebSocket is closed.
func (c WebSocket) OnClose(fn func(CloseEvent)) (remove func()) {
	_ = "STUB: not implemented"
	return nil
}

// OnError registers a function to be called when there is an error
// with the WebSocket.
func (c WebSocket) OnError(fn func(e js.Value)) (remove func()) {
	_ = "STUB: not implemented"
	return nil
}

// MessageEvent is the type passed to a message handler.
type MessageEvent struct {
	// string or []byte.
	Data any

	// There are more fields to the interface but we don't use them.
	// See https://developer.mozilla.org/en-US/docs/Web/API/MessageEvent
}

// OnMessage registers a function to be called when the WebSocket receives a message.
func (c WebSocket) OnMessage(fn func(m MessageEvent)) (remove func()) {
	_ = "STUB: not implemented"
	return nil
}

// Subprotocol returns the WebSocket subprotocol in use.
func (c WebSocket) Subprotocol() string { _ = "STUB: not implemented"; return "" }

// OnOpen registers a function to be called when the WebSocket is opened.
func (c WebSocket) OnOpen(fn func(e js.Value)) (remove func()) {
	_ = "STUB: not implemented"
	return nil
}

// Close closes the WebSocket with the given code and reason.
func (c WebSocket) Close(code int, reason string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// SendText sends the given string as a text message
// on the WebSocket.
func (c WebSocket) SendText(v string) (err error) { _ = "STUB: not implemented"; return nil }

// SendBytes sends the given message as a binary message
// on the WebSocket.
func (c WebSocket) SendBytes(v []byte) (err error) { _ = "STUB: not implemented"; return nil }

func extractArrayBuffer(arrayBuffer js.Value) []byte { _ = "STUB: not implemented"; return nil }

func uint8Array(src []byte) js.Value { _ = "STUB: not implemented"; return *new(js.Value) }
