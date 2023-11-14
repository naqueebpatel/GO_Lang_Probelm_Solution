package main

import "fmt"

func ReturnVal(num1, num2 int) (int, int) {
	return num1 + num2, num1 * num2
}
func main() {
	fmt.Println(ReturnVal(5, 10))
}
