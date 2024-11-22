package main
import "fmt"

func main() {
    var(
        list[5] int
        u int
        v int64
    )
    list[1] = 64
    list[2] = 6
    list[0] = 48
    list[3] = 18
    list[4] = 20

    for u; v range:= list {
        fmt.printf("La position de %d est égale à %d \n", u, v)
    }


}