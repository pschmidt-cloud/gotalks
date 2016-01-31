package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c *Circle) scale(f float64)  {
	c.radius = c.radius * f
}

func main() {
	v := Circle{1}
	v.scale(2)
	fmt.Println(v.area())
}
