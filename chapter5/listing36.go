package main

type notifier interface {
	notify()
}

func sendNotification(n notifier) {
	n.notify()
}

func main1() {
	u := user{"bill", "bill@emal.com"}
	sendNotification(&u)

	a := admin{"zhai", "zhai@173.com"}
	sendNotification(&a)

	p := apple{user{"apple", "apple@1.com"}, "apple@163.com"}
	sendNotification(&p)

	p.user.notify()
}
