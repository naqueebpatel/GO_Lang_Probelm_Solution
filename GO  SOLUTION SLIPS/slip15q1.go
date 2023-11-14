package main

import "fmt"

func main() {
	var Even [10]int
	var Odd [10]int
	var i, size, num int
	var x, y int
	fmt.Println("ENTER NO OF ELEMENTS?")
	fmt.Scanln(&size)
	for i = 0; i < size; i++ {
		fmt.Println("ENTER ELEMENT")
		fmt.Scanln(&num)
		if num%2 == 0 {
			Even[x] = num
			x += 1
		} else {
			Odd[y] = num
			y += 1
		}
	}
	fmt.Println("EVEN ARRAY:::")
	for i = 0; i < x; i++ {
		fmt.Println(Even[i])
	}
	fmt.Println("ODD ARRAY:::")
	for i = 0; i < y; i++ {
		fmt.Println(Odd[i])
	}
}
