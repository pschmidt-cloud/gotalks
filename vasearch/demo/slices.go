package main
import "fmt"

func main() {
	riversOfRussia := [3]string{"лена", "амур", "об"}  // array
	s := []string{}   // slice -- or can be made with make([]string,1)

	s = append(s, "дон") // can append to a slice but not to an array
	fmt.Println(riversOfRussia, s)
}
