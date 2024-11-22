package main

import "fmt"

type Animal interface{
	noise() string
	nombre() int
}

type chat struct{
	nom string
	breed string
}
type serpent struct{
	nom string 
	breed string
	verminous bool
}
type perroquet struct {
	nom string
	breed string
}

func info(a Animal) {
	fmt.Println(" Le bruit de cet animal est ",a.noise(),"il possède",a.nombre(), "pattes")
}
func (c *chat) noise() string {
	return "Miaou"
}
func (c *chat) nombre() int {
	return 4
}
func (s serpent) noise() string {
	
	return "ssssssss"
}
func (s serpent) nombre() int{
	return 10 
}
func main() {
	Chat := chat {
		nom: "Minou",
		breed: "Siamois",
	}
	info(&Chat)
	Serpent := serpent {
		nom: "kpo",
		breed: "Python regius",
		verminous: true,
	}
	info(&Serpent)


}