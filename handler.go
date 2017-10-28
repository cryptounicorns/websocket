package websocket

import (
	"io"
	"net/http"
)

type Handler interface {
	ServeWebsocket(w io.Writer, r *http.Request)
}
