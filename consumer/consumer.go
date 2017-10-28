package consumer

import (
	"io"
	"io/ioutil"

	"github.com/corpix/loggers"
	"github.com/corpix/loggers/logger/prefixwrapper"
)

type Consumer struct {
	reader io.Reader
	log    loggers.Logger

	messages chan []byte
	done     chan struct{}
}

func (c *Consumer) Consume() <-chan []byte {
	go c.consumeLoop()

	return c.messages
}

func (c *Consumer) consumeLoop() {
	var (
		l = prefixwrapper.New(
			"consumerLoop: ",
			c.log,
		)

		buf []byte
		err error
	)

	for {
		buf, err = ioutil.ReadAll(c.reader)
		if err != nil {
			l.Error(err)
			continue
		}

		select {
		case <-c.done:
			return
		default:
			c.messages <- buf
		}
	}
}

func (c *Consumer) Close() error {
	close(c.done)
	close(c.messages)

	return nil
}

func New(r io.Reader, l loggers.Logger) *Consumer {
	return &Consumer{
		reader: r,
		log: prefixwrapper.New(
			"Consumer: ",
			l,
		),
		messages: make(chan []byte),
		done:     make(chan struct{}),
	}
}
