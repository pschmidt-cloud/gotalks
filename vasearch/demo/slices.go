package main
import "fmt"

func main() {
	реки := [3]string{"лена", "амур", "волга"}
	города := []string{"сочи"}

	города = append(города, "москва") // can append to a slice but not to an array
	fmt.Println(реки, города)
}
