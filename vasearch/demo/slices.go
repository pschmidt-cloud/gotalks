package main
import "fmt"

func main() {
	easternRivers := [3]string{"лена", "амур", "об"}  // array
	westernRiver := []string{}   // slice -- or can be made with make([]string,1)

	westernRiver = append(westernRiver, "дон") // can append to a slice but not to an array
	fmt.Println(easternRivers, westernRiver)
}
