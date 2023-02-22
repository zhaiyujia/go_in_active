package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Contact struct {
	Name    string `json:"name"`
	Title   string `json:"title"`
	Contact struct {
		Home string `json:"home"`
		Call string `json:"call"`
	} `json:"contact"`
}

var JSON = `{
    "name": "GO",
    "title": "ammer",
    "contact": {
        "home": "415",
        "call": "555"
    }
}`

func main27() {
	var c Contact
	err := json.Unmarshal([]byte(JSON), &c)
	if err != nil {
		log.Println("ERRPR:", err)
		return
	}
	fmt.Println(c.Contact)
}
