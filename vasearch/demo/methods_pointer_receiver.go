package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// passes a copy of a Circle
func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

// passes a pointer to a Circle
func (c *Circle) scale(f float64)  {
	c.radius = c.radius * f
}

func main() {
	v := Circle{1}
	v.scale(2)
	fmt.Println(v.area())
}
