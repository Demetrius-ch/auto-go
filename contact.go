package main

import "fmt"

type contact struct {
	name string
	age int
	phone map[string]string
}
func newcontact(name string, age int) contact {
	c:=contact{
		name:  name,
		age:   age,
		phone: map[string]string{},
	}
	return c
}
func main() {
	mynewcontact := newcontact("Chandler", 22)
	fmt.Println(mynewcontact)
}