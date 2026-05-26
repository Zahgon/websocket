//go:build !amd64 && !arm64 && !js

package websocket

func mask(b []byte, key uint32) uint32 { _ = "STUB: not implemented"; return 0 }
