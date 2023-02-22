package main

import (
	_ "go_in_active/chapter2/matchers"
	"go_in_active/chapter2/search"

	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("NPR")
}
