package main

import "fmt"

func main() {
	var Arr [10]int
	var ptr *int
	var size int
	fmt.Println("ENTER NO OF ELEMENTS")
	fmt.Scanln(&size)
	for i := 0; i < size; i++ {
		fmt.Println("ENTER ELEMENTS?")
		fmt.Scanln(&Arr[i])
	}
	for i := size - 1; i >= 0; i-- {
		ptr = &Arr[i]
		fmt.Println(*ptr)
	}
}
