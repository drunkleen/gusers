package utmpx

// #include <utmpx.h>
// #include <string.h>
import "C"

// Close the utmpx database.
//
// This must be called after a loop of calling next() to ensure the database is
// closed.
func closeutmpx() {
	C.endutxent()
}

// Reset the utmpx database.
//
// This resets the internal pointer to the first entry so that a loop of next()
// will start from the beginning of the database.
func resetutmpx() {
	C.setutxent()
}

// toStr takes a C string and its maximum length, and returns a Go string.
//
// If the C string is nil or its maximum length is 0, an empty string is
// returned. Otherwise, the C string is converted to a Go string, but only up
// to the length returned by C.strnlen. If C.strnlen returns 0, an empty string
// is returned.
func toStr(b *C.char, max C.size_t) string {
	if b == nil || max == 0 {
		return ""
	}
	l := C.strnlen(b, max)
	if l == 0 {
		return ""
	}
	return C.GoStringN(b, C.int(l))
}
