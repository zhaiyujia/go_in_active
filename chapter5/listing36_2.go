package main

import "fmt"

type admin struct {
	name  string
	email string
}

func (u *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", u.name, u.email)
}
