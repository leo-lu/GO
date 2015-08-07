package main

import (
	"fmt"
)

type Notifyer interface {
	Notify()
	test()
}

type user struct {
	name  string
	email string
}

func (u *user) Notify() {
	fmt.Println("user notify.", u.name)
}

func (u user) test() {
	fmt.Println("test")
}

func main() {
	d := &user{"xiao", "xiao@.xxx.com"}
	test(d)
}

func test(n Notifyer) {
	v, rst := n.(*user)
	if rst {
		v.Notify()
	}

	switch v := n.(type) {
	case *user:
		fmt.Println(v.name)
		v.Notify()
	default:
		fmt.Println("none...")
	}
}
