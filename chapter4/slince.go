package main

import "fmt"

func testLi() {
	source := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}

	newLice := source[2:3:4]

	for index, data := range newLice {
		fmt.Printf("下标:%d, 数据：%s", index, data)
		fmt.Println()
	}
}
