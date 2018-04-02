package apidoc

import (
	"fmt"

	"github.com/austinjalexander/util/pkg/server"
)

// Doc writes API documentation.
type Doc struct {
	baseDir string
}

// New returns a new, configured doc.
func New(dir string) Doc {
	return Doc{baseDir: dir}
}

// Write takes a server handler and writes it.
func (d Doc) Write(h server.Handler) server.Handler {
	fmt.Printf("%+v", h)
	return h
}
