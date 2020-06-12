package main

import (
	"io"
	"bytes"
)

func test(w io.Writer) {
    if w != nil {
        w.Write([]byte("ok"))
    }
}

func main() {
    var buf bytes.Buffer
    test(buf)
}