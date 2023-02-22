package main

import "fmt"

type apple struct {
	user  user
	email string
}

func (u *apple) notify() {
	fmt.Printf("Sending apple email to %s<%s>\n", u.user.name, u.email)
}
