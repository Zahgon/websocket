package bpool

import (
	"bytes"
	"sync"
)

var bpool = sync.Pool{
	New: func() any {
		return &bytes.Buffer{}
	},
}

// Get returns a buffer from the pool or creates a new one if
// the pool is empty.
func Get() *bytes.Buffer { _ = "STUB: not implemented"; return nil }

// Put returns a buffer into the pool.
func Put(b *bytes.Buffer) { _ = "STUB: not implemented"; return }
