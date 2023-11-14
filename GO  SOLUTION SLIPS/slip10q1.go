package main

import "fmt"

var sum, rem int

func SumDigi(num int) int {
	if num > 0 {
		rem = num % 10
		sum = sum + rem
		num = num / 10
		SumDigi(num)
	}
	return sum
}
func main() {
	var num int
	fmt.Println("ENTER NUMBER TO FIND?")
	fmt.Scanln(&num)
	fmt.Println("SUM OF DIGITS:::", SumDigi(num))
}
