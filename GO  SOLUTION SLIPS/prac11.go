package main

import "fmt"

func swap(x *int, y *int) {
	temp := *x
	*x = *y
	*y = temp
	fmt.Println("VALUE OF X AFTER SWAP", *x)
	fmt.Println("VALUE OF Y AFTER SWAP", *y)
}
func main() {
	x := 10
	y := 20
	fmt.Println("VALUE OF X BEFORE SWAP", x)
	fmt.Println("VALUE OF Y BEFORE SWAP", y)
	swap(&x, &y)
}
