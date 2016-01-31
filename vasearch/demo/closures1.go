package main
import "fmt"

func main() {
	f := addToNumber()
	fmt.Printf("num1=%d\n", f(1))
	fmt.Printf("num1=%d\n", f(1))
}

func addToNumber() func(i int) int {
	val := 0

	return func(n int) int {
		val = val + n
		return val
	}
}