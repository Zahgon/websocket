package errd

// Wrap wraps err with fmt.Errorf if err is non nil.
// Intended for use with defer and a named error return.
// Inspired by https://github.com/golang/go/issues/32676.
func Wrap(err *error, f string, v ...any) { _ = "STUB: not implemented"; return }
