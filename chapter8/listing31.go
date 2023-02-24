package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main31() {
	c := make(map[string]interface{})
	c["name"] = "go"
	c["title"] = "ammer"
	c["contact"] = map[string]interface{}{
		"home": "115",
		"call": "555",
	}

	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Println(string(data))

}
