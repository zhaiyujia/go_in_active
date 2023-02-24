package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	r, err := http.Get("http://www.baidu.com")
	if err != nil {
		log.Fatalln(err)
	}

	file, err := os.Create("baidu.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	dest := io.MultiWriter(os.Stdout, file)

	io.Copy(dest, r.Body)

	if err := r.Body.Close(); err != nil {
		log.Fatalln(err)
	}
}
