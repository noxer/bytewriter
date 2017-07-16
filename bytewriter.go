package bytewriter

import "io"

type Writer struct {
	p []byte
	s int
}

// New creates a new io.Writer that writes all data in to p, it will never reallocate p.
// When p is full the writer returns io.EOF. The writer does not support concurrent use.
func New(p []byte) *Writer {
	return &Writer{p: p, s: len(p)}
}

// Write writes bytes into the byte slice, returning io.EOF when it is full. It will
// never expand the slice.
func (w *Writer) Write(p []byte) (n int, err error) {
	// Check for a full slice
	if len(w.p) == 0 {
		return 0, io.EOF
	}

	// Copy the data
	copy(w.p, p)

	if len(w.p) < len(p) {
		// Completely filled w.p and more bytes waiting
		n = len(w.p)
		err = io.EOF
		w.p = nil
		return
	}
	if len(w.p) == len(p) {
		// Completely filled w.p but no more bytes waiting
		n = len(p)
		w.p = nil
		return
	}

	// Fitting write
	w.p = w.p[len(p):]
	n = len(p)
	return
}

// Written returns the number of bytes that have been written to the slice.
func (w *Writer) Written() int {
	return w.s - len(w.p)
}
