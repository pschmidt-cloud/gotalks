package main

// START OMIT
func main() {
	stackDemo()        		// go build -gcflags=-m <gofile> to see the analysis output
	oops := heapDemo()  	// referenced outside of the function so put on the  heap
}

func stackDemo()  {
	arr := make([]int, 10)     // not referenced outside the function so put on the stack
	arr[0] = 10
}

func heapDemo() []int {
	return make([]int, 10)
}
// END OMIT
