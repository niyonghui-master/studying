package main

import (
	"bytes"
	"fmt"
	"io"
)

func test(w io.Writer) {
    if w != nil {
		w.Write([]byte("ok"))
		fmt.Println("nil")
    } else {
		fmt.Println("nil")
	}
}

func main() {
	// var buf io.Writer
	var buf *bytes.Buffer
	buf = new(bytes.Buffer)
    test(buf)
}