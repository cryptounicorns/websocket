package websocket

import (
	"io"
	"sync"

	"github.com/gorilla/websocket"
)

type Reader struct {
	*websocket.Conn
	unread []byte
	lock   *sync.Mutex
}

func (w *Reader) Read(buf []byte) (int, error) {
	w.lock.Lock()
	defer w.lock.Unlock()

	var (
		t   int
		b   []byte
		n   int
		err error
	)

	if len(w.unread) > 0 {
		b = w.unread
	} else {
		for {
			t, b, err = w.Conn.ReadMessage()
			if err != nil {
				return len(b), err
			}
			switch t {
			case websocket.TextMessage:
			case websocket.BinaryMessage:
			default:
				// XXX: Unsupported message
				continue
			}

			break
		}
	}

	n = copy(buf, b)
	if n < len(b) {
		w.unread = b[n:]
	} else {
		w.unread = nil
	}

	if len(w.unread) == 0 {
		err = io.EOF
	}

	return n, err
}

func NewReader(c *websocket.Conn) *Reader {
	return &Reader{
		Conn: c,
		lock: &sync.Mutex{},
	}
}
