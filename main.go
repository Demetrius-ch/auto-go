package main
import (
	"fmt"
	"github.com/Demetrius-ch/myprogramm/footypes"
)

func main() {
	fmt.Println("test")

	var fooVar footypes.Foo
	fooVar.TypeInt = 18
	fmt.Println(fooVar)
}
