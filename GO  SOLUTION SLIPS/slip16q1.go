package main

import "fmt"

func main() {
	var Arr [10]int
	var i, size, Even, Odd int
	fmt.Println("ENTER NO OF ELEMENTS?")
	fmt.Scanln(&size)
	for i = 0; i < size; i++ {
		fmt.Println("ENTER ELEMENT")
		fmt.Scanln(&Arr[i])
		if Arr[i]%2 == 0 {
			Even += 1
		} else {
			Odd += 1
		}
	}
	fmt.Println("TOTAL EVEN COUNT:::", Even)
	fmt.Println("TOTAL ODD COUNT:::", Odd)
}
