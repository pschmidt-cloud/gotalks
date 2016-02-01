package main
import "fmt"

type Person struct{}
type Employee struct{}

type Manager struct {
	Person
	Employee
}

func (p Person) name() string {
	return "Hello, my name is ..."
}

func (e Employee) numOfEmployees() int {
	return 10
}

func main() {
	mgr := new(Manager)
	fmt.Println(mgr.name())
	fmt.Println(mgr.numOfEmployees())
}
