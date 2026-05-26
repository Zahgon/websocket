package xsync

// Go allows running a function in another goroutine
// and waiting for its error.
func Go(fn func() error) <-chan error { _ = "STUB: not implemented"; return nil }
