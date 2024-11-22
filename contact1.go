package main

import "fmt"

type contact1 struct {
	nom string
	age int
	numbers map[string]string

}
func Newcontact(nom string, age int, numbers map[string]string) contact1{
	c1 := contact1{
		nom: nom,
		age: age,
		numbers: numbers,
	}
	return c1
}

func (c1 contact1) displayContact() string {
	display := fmt.Sprintf("Contact : %v (%v ans)\n", c1.nom, c1.age)

	display+= "-----------------------\n"
	for u, v:= range c1.numbers{
		display= fmt.Sprintf("%v %v\n", u,v)
	}
	return display

}