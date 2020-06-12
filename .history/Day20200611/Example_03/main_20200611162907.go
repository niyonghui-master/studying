package main

import (
	"io"
)

func test(w io.Writer) {
    if w != nil {
        w.Write([]byte("ok"))
    } else {
		fmt.Pr
	}
}

func main() {
    var buf io.Writer
    test(buf)
}