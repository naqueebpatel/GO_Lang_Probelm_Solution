package main

import "fmt"

type Shape interface {
	area() float64
	peri() float64
}
type circle struct {
	rad float64
}

func (c circle) area() float64 {
	return (3.14 * c.rad * c.rad)
}
func (c circle) peri() float64 {
	return (2 * 3.14 * c.rad)
}
func main() {
	var obj Shape
	obj = circle{5}
	fmt.Println("AREA:::", obj.area())
	fmt.Println("PERIMETER:::", obj.peri())
}
