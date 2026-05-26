package websocket

// maskGo applies the WebSocket masking algorithm to p
// with the given key.
// See https://tools.ietf.org/html/rfc6455#section-5.3
//
// The returned value is the correctly rotated key to
// to continue to mask/unmask the message.
//
// It is optimized for LittleEndian and expects the key
// to be in little endian.
//
// See https://github.com/golang/go/issues/31586
func maskGo(b []byte, key uint32) uint32 { _ = "STUB: not implemented"; return 0 }

// At some point in the future we can clean these unrolled loops up.
// See https://github.com/golang/go/issues/31586#issuecomment-487436401

// Then we xor until b is less than 128 bytes.

// Then we xor until b is less than 64 bytes.

// Then we xor until b is less than 32 bytes.

// Then we xor until b is less than 16 bytes.

// Then we xor until b is less than 8 bytes.

// Then we xor until b is less than 4 bytes.

// xor remaining bytes.
