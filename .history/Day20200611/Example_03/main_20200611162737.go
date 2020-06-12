package main

import (
	""
)

func test(w io.Writer) {
    if w != nil {
        w.Write([]byte("ok"))
    }
}

func main() {
    var buf *bytes.Buffer
    test(buf)
}