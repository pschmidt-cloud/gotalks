package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)

	arr := [3]int{10,20,30}
	for _, v := range arr {
		fmt.Println(v)
	}
}