package main

import (
	"fmt"
	"math"
)

type MyCustomInteger int

func (f MyCustomInteger) Abs() int {
	if f < 0 {
		return int(-f) + 20
	}
	return int(f)
}

func main() {
	f := MyCustomInteger(-2)
	fmt.Println(f.Abs())
	fmt.Println(math.Abs(-2))
}