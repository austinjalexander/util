// compare-files -filepath1=./compare-files/testdata/f1 -filepath2=./compare-files/testdata/f2
// compare-files -filepath1=./compare-files/testdata/f1 -filepath2=./compare-files/testdata/f1-copy

// go test -v ./...
package main

import (
	"bytes"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	filepath1 = flag.String("filepath1", "", "The filepath of the first file.")
	filepath2 = flag.String("filepath2", "", "The filepath of the second file.")
)

func init() {
	flag.Parse()
}

func main() {
	if *filepath1 == "" || *filepath2 == "" {
		log.Fatal("error: need to include both a filepath1 and a filepath2")
	}

	isEqual, b1, b2, err := compareFiles(*filepath1, *filepath2)
	if err != nil {
		log.Fatal(err)
	}
	if isEqual {
		fmt.Println("\n\tf1 and f2 have the same contents!\n")
	} else {
		fmt.Println("\n\tf1 is different from f2!\n")
	}
	fmt.Printf("f1: %+v\n\n", b1)
	fmt.Printf("f2: %+v\n\n", b2)
	fmt.Println()
}

func compareFiles(filepath1, filepath2 string) (bool, []byte, []byte, error) {
	f1, err := os.Open(filepath1)
	if err != nil {
		return false, nil, nil, fmt.Errorf("error: opening filepath1 %q: %s", filepath1, err)
	}

	f2, err := os.Open(filepath2)
	if err != nil {
		return false, nil, nil, fmt.Errorf("error: opening filepath2 %q: %s", filepath2, err)
	}

	h1 := sha256.New()
	n1, err := io.Copy(h1, f1)
	if err != nil {
		return false, nil, nil, fmt.Errorf("error: copying f1: %s", err)
	}

	h2 := sha256.New()
	n2, err := io.Copy(h2, f2)
	if err != nil {
		return false, nil, nil, fmt.Errorf("error: copying f2: %s", err)
	}

	if n1 != n2 {
		return false, nil, nil, fmt.Errorf("error: number of bytes written for f1: %d, number of bytes written for f2: %d", n1, n2)
	}

	b1 := h1.Sum(nil)
	b2 := h2.Sum(nil)

	if !bytes.Equal(b1, b2) {
		return false, b1, b2, nil
	}
	return true, b1, b2, nil
}
