package bytewriter

import (
	"bytes"
	"io"
	"testing"
)

func TestWriter(t *testing.T) {
	buf := make([]byte, 8)

	w := New(buf)
	if w == nil {
		t.Fatal("New returned nil")
	}
	if b := w.Written(); b != 0 {
		t.Errorf("Expected written 0; got %d", b)
	}

	n, err := w.Write([]byte("testing"))
	if n != 7 || err != nil {
		t.Errorf("Expected write (7, nil); got (%d, %#v)", n, err)
	}
	if w.Written() != 7 {
		t.Errorf("Expected written 7; got %d", w.Written())
	}

	n, err = w.Write([]byte("t"))
	if n != 1 || err != nil {
		t.Errorf("Expected write (1, nil); got (%d, %#v)", n, err)
	}
	if b := w.Written(); b != 8 {
		t.Errorf("Expected written 8; got %d", b)
	}

	n, err = w.Write([]byte("too much"))
	if n != 0 || err != io.EOF {
		t.Errorf("Expected write (0, EOF); got (%d, %#v)", n, err)
	}
	if b := w.Written(); b != 8 {
		t.Errorf("Expected written 8; got %d", b)
	}

	if !bytes.Equal(buf, []byte("testingt")) {
		t.Errorf("Expected buffer %#v; got %#v", buf, []byte("testingt"))
	}

	w = New(buf)

	n, err = w.Write([]byte("this is a long string"))
	if n != 8 || err != io.EOF {
		t.Errorf("Expected write (8, EOF); got (%d, %#v)", n, err)
	}

	if !bytes.Equal(buf, []byte("this is ")) {
		t.Errorf("Expected buffer %#v; got %#v", buf, []byte("this is "))
	}
	if b := w.Written(); b != 8 {
		t.Errorf("Expected written 8; got %d", b)
	}

	n, err = w.Write([]byte("t"))
	if n != 0 || err != io.EOF {
		t.Errorf("Expected write (0, EOF); got (%d, %#v)", n, err)
	}
	if b := w.Written(); b != 8 {
		t.Errorf("Expected written 8; got %d", b)
	}
}
