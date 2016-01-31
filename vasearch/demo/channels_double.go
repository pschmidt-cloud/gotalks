package main
import (
	"fmt"
)

func main() {
	ch := make(chan uint64)
	go doubleNumbers(ch, 2, 4, 6)
	for v := range (ch) {
		fmt.Println(v)
	}

	/*sample output
	/usr/local/go/bin/go run /Users/schmidtp/go/src/test/demo/freeform.go
		0 2
		1 4
		4
		8
		2 6
		12
	 */
}


func doubleNumbers(ch chan uint64, nums ...uint64) uint64 {
	var sum uint64

	for k, v := range (nums) {
		fmt.Println(k, v)
		ch <- v*2
	}

	close(ch)
	return sum
}

