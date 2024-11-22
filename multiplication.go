package main

import "fmt"

var (
	x int
	y int
)

func main() {
	var resultat int
	fmt.Println("Donner la table mutliplication")
	fmt.Scan(&x)
	fmt.Println("Donner les nombre à multiplier")
	fmt.Scan(&y)
	nombre := make([]int, y )

	for y:=0; y<=12; y++{
		resultat = x*y

		
		fmt.Printf("%v * %v = %v\n", x, y, resultat)
	}

}