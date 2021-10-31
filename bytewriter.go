package bytewriter

import "errors"

// ErrSliceFull is returned when the target byte slice can hold no more data.
var ErrSliceFull = errors.New("bytewriter: slice is full")

// Writer allows writing into a byte slice without reallocating it.
type Writer struct {
	p []byte
	s int
}

// New creates a new io.Writer that writes all data in to p, it will never reallocate p.
// When p is full the writer returns bytewriter.ErrSliceFull as the error. The writer does not support concurrent use.
func New(p []byte) *Writer {
	return &Writer{p: p, s: len(p)}
}

// Write writes bytes into the byte slice, returning bytewriter.ErrSliceFull as an error when it is full. It will
// never expand the slice.
func (w *Writer) Write(p []byte) (n int, err error) {
	// Check for a full slice
	if len(w.p) == 0 {
		return 0, ErrSliceFull
	}

	// Copy the data
	copy(w.p, p)

	if len(w.p) < len(p) {
		// Completely filled w.p and more bytes waiting
		n = len(w.p)
		err = ErrSliceFull
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

// WriteByte writes a single byte into the writer.
func (w *Writer) WriteByte(c byte) error {
	// Check for a full slice
	if len(w.p) == 0 {
		return ErrSliceFull
	}

	w.p[0] = c

	if len(w.p) == 1 {
		w.p = nil
	} else {
		w.p = w.p[1:]
	}

	return nil
}

// Written returns the number of bytes that have been written to the slice.
func (w *Writer) Written() int {
	return w.s - len(w.p)
}
