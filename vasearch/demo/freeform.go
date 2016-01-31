package main
import (
	"fmt"
	"strings"
)

//go:generate stringer -type Lang


type Lang struct {
	Name string
	Year int
	URL  string
}

func (l Lang) String() string {
	return "name=" + l.Name
}

var urls = []Lang{{"golang", 1999, "http://www.ingenuity.com"}}

func main() {
	s1 := []int{2, 4, 6}
	s1 = append(s1, 7)
	lang1 := Lang{"qiagen", 1900, "www.qiagen.com"}
	s2 := make([]int, 10)
	copy(s2, s1)

	for i :=0; i < len(s1); i ++ {
		fmt.Println(i)
	}

	for k, v := range urls {
		fmt.Println(k, v)
	}

	var str string = "testing again"

	for key, val := range strings.Fields(str) {
		fmt.Println(key, val)
	}

	fmt.Println(s1, lang1, s2)
}



