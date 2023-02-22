package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func test() {
	var b bytes.Buffer

	b.Write([]byte("Hello"))

	fmt.Fprintln(&b, " World!")

	io.Copy(os.Stdout, &b)
}
