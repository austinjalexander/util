package apidoc

import (
	"fmt"
	"github.com/austinjalexander/util/pkg/server"
)

// doc writes API documentation.
type doc struct {
	baseDir     string
	handlerChan chan server.Handler
}

// New returns a new, configured doc.
func New(dir string) doc {
	return doc{
		baseDir:     dir,
		handlerChan: make(chan server.Handler),
	}
}

// Write takes a server handler and writes it.
func (d doc) Write(h server.Handler) server.Handler {
	d.handlerChan <- h
	return h
}

func (d doc) Flush() {
	close(d.handlerChan)
	for h := d.handlerChan {
		fmt.Println(h)
	}
}
