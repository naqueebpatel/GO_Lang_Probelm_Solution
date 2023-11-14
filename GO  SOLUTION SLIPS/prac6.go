package main

import "fmt"

var sum, rem int

func Sum(num int) {
	if num > 0 {
		rem = num % 10
		sum = sum + rem
		Sum(num / 10)
	} else {
		fmt.Println("SUM OF DIGITS IS:::", sum)
	}
}

func main() {
	var num int
	fmt.Println("ENTER NUMBER TO FIND SUM?")
	fmt.Scanln(&num)
	Sum(num)
}
