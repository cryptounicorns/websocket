package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/logrus"
	"github.com/corpix/loggers/logger/prefixwrapper"
	"github.com/gorilla/mux"
	logrusLogger "github.com/sirupsen/logrus"

	"github.com/cryptounicorns/websocket"
)

type Message struct {
	Text   string
	Number int
}

func NewMessage(text string, num int) *Message {
	return &Message{
		Text:   text,
		Number: num,
	}
}

func Mount(r *mux.Router, l loggers.Logger) {
	var (
		handler = func(w io.WriteCloser, r *http.Request) {
			var (
				m   []byte
				n   = 0
				err error
			)
			defer w.Close()

			for {
				m, err = json.Marshal(NewMessage("hello", n))
				if err != nil {
					l.Error(err)
					return
				}
				n++

				_, err = w.Write(m)
				if err != nil {
					// XXX: Because gorilla's error is private we can't
					// type assert. But we can apply this "hack" to obtain type name
					// still better than strings.HasPrefix and others in my opinion.
					switch fmt.Sprintf("%T", err) {
					case "*net.OpError":
					case "*websocket.netError":
					default:
						l.Error(err)
					}

					l.Debugf(
						"Terminating loop for %#v",
						w,
					)
					return
				}
			}
		}
	)
	r.Handle(
		"/",
		websocket.NewHTTPUpgradeHandler(
			websocket.HandlerFunc(handler),
			prefixwrapper.New("GET /: ", l),
		),
	)
}

func main() {
	var (
		r    = mux.NewRouter()
		l    = logrus.New(logrusLogger.New())
		addr = "127.0.0.1:3333"
	)
	Mount(r, l)

	l.Print("Listening ", addr)
	err := http.ListenAndServe(addr, r)
	if err != nil {
		l.Fatal(err)
	}

	select {}
}
