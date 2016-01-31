package main
import "fmt"

type Object1 struct{}
type Object2 struct{}

type Object3 struct {
	Object1
	Object2
}

func (O Object1) func1() string {
	return "doThis"
}

func (O Object2) func2() string {
	return "doThat"
}

func main() {
	obj3 := new(Object3)
	fmt.Println(obj3.func1())
	fmt.Println(obj3.func2())
}
