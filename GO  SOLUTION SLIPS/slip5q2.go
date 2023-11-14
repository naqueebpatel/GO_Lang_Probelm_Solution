package main

import "fmt"

type Shape interface {
	area() float64
	peri() float64
}
type circle struct {
	radius float64
}
type rect struct {
	length, breadth float64
}

func (c circle) area() float64 {
	return (3.14 * c.radius * c.radius)
}
func (c circle) peri() float64 {
	return (2 * 3.14 * c.radius)
}
func (r rect) area() float64 {
	return (r.length * r.breadth)
}
func (r rect) peri() float64 {
	return (2 * (r.length + r.breadth))
}

func main() {
	var obj Shape
	var r, l, b float64
	fmt.Println("ENTER RADIUS OF CIRCLE")
	fmt.Scanln(&r)
	fmt.Println("ENTER LENGTH OF RECTANGLE")
	fmt.Scanln(&l)
	fmt.Println("ENTER BREADTH OF RECTANGLE")
	fmt.Scanln(&b)
	obj = circle{r}
	fmt.Println("AREA CIRCLE:::", obj.area())
	fmt.Println("PERIMETER CIRCLE:::", obj.peri())
	obj = rect{l, b}
	fmt.Println("AREA RECTANGLE:::", obj.area())
	fmt.Println("PERIMETER RECTANGLE:::", obj.peri())
}
