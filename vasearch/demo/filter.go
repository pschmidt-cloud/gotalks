package main
import "fmt"

func main() {
	nums := []int{10, 50, 100}
	anon := func (i int) bool {
		return i == 10
	}
	newNums := filterSliceDemo(nums, anon)
	fmt.Println(newNums)
}

func filterSliceDemo(a[] int, fn func(i int) bool) []int {
	p := []int{}

	for _, v := range (a) {
		if fn(v) {
			p = append(p, v)
		}
	}
	return p
}