package apidoc

import (
	"fmt"

	"github.com/austinjalexander/util/pkg/server"
)

// Doc writes API documentation.
type Doc struct {
	baseDir     string
	handlerChan chan server.Handler
}

// New returns a new, configured Doc.
func New(dir string) Doc {
	return Doc{
		baseDir:     dir,
		handlerChan: make(chan server.Handler),
	}
}

// Write takes a server handler and writes it.
func (d Doc) Write(h server.Handler) server.Handler {
	d.handlerChan <- h
	return h
}

// Flush closes the handler channel and writes the documentation to disk.
func (d Doc) Flush() {
	fmt.Println("HEY")
	close(d.handlerChan)
	for h := range d.handlerChan {
		fmt.Println(h)
	}
}
