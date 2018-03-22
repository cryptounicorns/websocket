package websocket

import (
	"github.com/gorilla/websocket"
)

type Writer struct {
	*websocket.Conn
}

func (w *Writer) Write(buf []byte) (int, error) {
	return len(buf), w.Conn.WriteMessage(websocket.BinaryMessage, buf)
}
