package websocket

import (
	"io"
	"net/http"
)

type Handler interface {
	ServeWebsocket(w io.WriteCloser, r *http.Request)
}

type HandlerFunc func(w io.WriteCloser, r *http.Request)

func (hf HandlerFunc) ServeWebsocket(w io.WriteCloser, r *http.Request) {
	hf(w, r)
}
