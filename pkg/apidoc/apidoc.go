package apidoc

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"

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
		handlerChan: make(chan server.Handler, 1),
	}
}

// Write takes a server handler and writes it.
func (d Doc) Write(h server.Handler) server.Handler {
	d.handlerChan <- h
	return h
}

// Flush closes the handler channel and writes the documentation to disk.
func (d Doc) Flush() error {
	close(d.handlerChan)

	for h := range d.handlerChan {
		fmt.Println(h)

		dir := filepath.Join(d.baseDir, h.Path)
		err := os.MkdirAll(dir, 0666)
		if err != nil {
			return err
		}

		f, err := os.Create(filepath.Join(dir, filepath.Base(h.Path)+".md"))
		if err != nil {
			return err
		}
		defer func() {
			err := f.Close()
			if err != nil {
				log.Println(err)
			}
		}()

		w := bufio.NewWriter(f)

		for _, method := range h.Methods {
			_, err := w.WriteString(fmt.Sprintf("# %s `%s`\n\n", method, h.Path))
			if err != nil {
				return err
			}
		}
		w.Flush()
	}
	return nil
}
