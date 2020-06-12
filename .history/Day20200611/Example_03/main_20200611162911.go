package main

import (
	"fmt"
	"io"
)

func test(w io.Writer) {
    if w != nil {
        w.Write([]byte("ok"))
    } else {
		fmt.Println("nil")
	}
}

func main() {
    var buf io.Writer
    test(buf)
}