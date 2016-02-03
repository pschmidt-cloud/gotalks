package main
import "fmt"

// START OMIT
func main() {
	stackDemo()        		// go build -gcflags=-m <gofile> to see the analysis output
	oops := heapDemo()  	// referenced outside of the function so put on the  heap
	fmt.Println(oops)
}

func stackDemo()  {
	arr := make([]int, 5)     // not referenced outside the function so put on the stack
	arr[0] = 10
}

func heapDemo() []int {
	return make([]int, 10)
}
// END OMIT
