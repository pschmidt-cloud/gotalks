package main
import (
	"fmt"
	"math"
)

// START OMIT
type Shape interface {
	area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	q := Circle{1}
	shapes := []Shape{q}
	for n, _ := range shapes {
		fmt.Println("Area of this shape is: ", shapes[n].area())
	}
}
// END OMIT

type Rectangle struct {
	height float64
	width float64
}

func (r Rectangle)  area() float64 {
	return r.height * r.width
}



