package assert

import (
	"testing"
)

// Equal asserts exp == act.
func Equal(t testing.TB, name string, exp, got any) { _ = "STUB: not implemented"; return }

// Success asserts err == nil.
func Success(t testing.TB, err error) { _ = "STUB: not implemented"; return }

// Error asserts err != nil.
func Error(t testing.TB, err error) { _ = "STUB: not implemented"; return }

// Contains asserts the fmt.Sprint(v) contains sub.
func Contains(t testing.TB, v any, sub string) { _ = "STUB: not implemented"; return }

// ErrorIs asserts errors.Is(got, exp)
func ErrorIs(t testing.TB, exp, got error) { _ = "STUB: not implemented"; return }
