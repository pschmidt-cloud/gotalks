package main
import (
	"fmt"
	"runtime"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	nCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(nCPU)
	fmt.Println("Number of CPUs: ", nCPU)

	a := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	fmt.Println(a[:len(a)/2])
	fmt.Println(a[len(a)/2:])
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}