package main

import "fmt"

type Example struct {
	a string
	b int
	c bool
}

func main() {
	myExample := Example{
		a: "Chandler",
		b: 64,
		c: true,
	}
	fmt.Println(myExample)
	//mycontact := newcontact("Alex", 30)
	//fmt.Println(mycontact)
	myContact := Newcontact("Legrand", 22, map[string]string{"Fixe:": "05649446", "Portable:": "846664"})
	fmt.Println(myContact.displayContact())
}
