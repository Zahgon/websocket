package util

// WriterFunc is used to implement one off io.Writers.
type WriterFunc func(p []byte) (int, error)

func (f WriterFunc) Write(p []byte) (int, error) {
	_ = "STUB: not implemented"

	// ReaderFunc is used to implement one off io.Readers.
	return 0, nil
}

type ReaderFunc func(p []byte) (int, error)

func (f ReaderFunc) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }
