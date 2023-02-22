package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main29() {
	var c map[string]interface{}
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERROR:", err)
		return
	}
	fmt.Println("Name:", c["name"])
}
