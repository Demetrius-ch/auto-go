package main

import "fmt"

func main() {
	var n int
	fmt.Print("Combien de nombres voulez-vous entrer ?")
	fmt.Scanln(&n)
	
	nombres := make([]int, n)
	
	fmt.Println("Entrer les nombres: ")
	for i:=0; i<n; i++ {
		fmt.Scan(&nombres[i])
	}
	min, max := nombres[0], nombres[0] 
	
	for _, num := range nombres {
		if num < min {
			min = num
		}
		if num > max {
			max= num
		}
	}
	fmt.Println("Le plus petit nombre est : ", min)
	fmt.Println("Le plus grand nombre est : ", max)

}