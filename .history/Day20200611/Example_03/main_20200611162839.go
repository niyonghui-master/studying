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
    var buf io.Wre
    test(buf)
}