package writer

import (
	"io"
)

type WriteWith func(w io.Writer, buf []byte) error

type Writer struct {
	writer    io.Writer
	writeWith WriteWith
}

func (w *Writer) Write(m []byte) (int, error) {
	var (
		err error
	)

	err = w.writeWith(
		w.writer,
		m,
	)
	if err != nil {
		return 0, err
	}

	return len(m), nil
}

func New(w io.Writer, ww WriteWith) *Writer {
	return &Writer{
		writer:    w,
		writeWith: ww,
	}
}
