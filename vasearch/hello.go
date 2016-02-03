package main

import (
	"fmt"
)

func main() {
	var hello string = "hello"
	world := "world"

    fmt.Printf("%v %v \n", hello, world)
	fmt.Printf("%T %T", hello, world)
}
