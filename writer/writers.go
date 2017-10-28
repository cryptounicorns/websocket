package writer

import (
	"io"

	"github.com/gobwas/ws/wsutil"
)

func NewServerBinary(w io.Writer) io.Writer {
	return New(w, wsutil.WriteServerBinary)
}

func NewServerText(w io.Writer) io.Writer {
	return New(w, wsutil.WriteServerText)
}

func NewClientBinary(w io.Writer) io.Writer {
	return New(w, wsutil.WriteClientBinary)
}

func NewClientText(w io.Writer) io.Writer {
	return New(w, wsutil.WriteClientText)
}
