package main

import (
	"bytes"
	"fmt"
	"os"
)

func main37() {
	var b bytes.Buffer

	b.Write([]byte("Hello "))

	fmt.Fprintf(&b, "World")

	b.WriteTo(os.Stdout)
}
