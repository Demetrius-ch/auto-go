package main

import(
	"fmt"
	"github.com/Demetrius-ch/myprogramm/functions"
)

func CalculateValue(intChan chan int) {
	randomNumber := functions.GenerateRandomNumber(50)
	intChan <- randomNumber
}
func main() {
	foo:= make(chan int)
	defer close(foo)

	go CalculateValue(foo)
	num := <- foo
	
	fmt.Println(num)
}

