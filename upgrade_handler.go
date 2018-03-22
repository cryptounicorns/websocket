package websocket

import (
	"io"
	"net/http"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"
	"github.com/gobwas/ws"
)

var (
	DefaultHeaders = http.Header{}
)

// HTTPUpgradeHandler is a net/http.Handler which is responsible
// for incomming HTTP request upgrade process and handling the
// upgraded request with some Handler which works only with
// websockets.
type HTTPUpgradeHandler struct {
	Handler
	log loggers.Logger
}

func (h *HTTPUpgradeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	r.Body.Close()

	var (
		c   io.WriteCloser
		err error
	)

	c, _, _, err = ws.UpgradeHTTP(
		r,
		w,
		DefaultHeaders,
	)
	if err != nil {
		h.log.Error(err)
		return
	}

	h.ServeWebsocket(c, r)
}

func NewHTTPUpgradeHandler(h Handler, l loggers.Logger) *HTTPUpgradeHandler {
	return &HTTPUpgradeHandler{
		Handler: h,
		log: prefixwrapper.New(
			"HTTPUpgradeHandler: ",
			l,
		),
	}
}
