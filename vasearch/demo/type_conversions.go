package main

import (
	"fmt"
	"math"
	"strconv"
)

// Unlike in C, in Go assignment between items of different type requires an explicit conversion.

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	var str string = strconv.Itoa(33)
	fmt.Println(x, y, z, str)
}