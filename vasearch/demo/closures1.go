package main
import "fmt"


func main() {
	num1, num2 := adder(), adder()

	fmt.Printf("num1=%d, num2=%d \n", num1(1), num2(2))
	fmt.Printf("num1=%d, num2=%d \n", num1(5), num2(5))
}

func adder() func(i int) int {
	fmt.Println("init")
	sum := 0

	return func(n int) int {
		fmt.Printf("before adding, sum=[%d]\n", sum)
		sum = sum + n
		return sum
	}
}