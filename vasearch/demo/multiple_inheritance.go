package main
import "fmt"

// START OMIT
type Person struct{
	age int
}
type Employee struct{}
type Manager struct {
	Person   // embedded type. Manager IS A Person
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
	mgr.age = 35
}
// END OMIT
