package main
import (
	"fmt"
	"math"
)

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {
  	return math.Pi * math.Pow(c.radius, 2)
}

func main() {
	r := Rectangle{5, 3}
	q := Circle{5}
	shapes := []Shape{r, q}

	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Printf("Type of this shape is: %T \n", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].area())
	}
}