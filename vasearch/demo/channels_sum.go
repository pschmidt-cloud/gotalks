package main

import (
	"fmt"
	"runtime"
)

// START OMIT
func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}

func main() {
	fmt.Println("Number of CPUs: ", runtime.NumCPU())

	a := []int{10, 4, 6, 30, 8, 2}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
// END OMIT