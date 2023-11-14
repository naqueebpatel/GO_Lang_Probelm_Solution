package main

import "fmt"

func main() {
	var n int
	var Arr [10]int
	fmt.Println("ENTER THE SIZE OF ARRAY")
	fmt.Scanln(&n)
	fmt.Println("ENTER ARRAY ELEMENT")
	for i := 0; i < n; i++ {
		fmt.Scanln(&Arr[i])
	}
	fmt.Println("REVERSED ARRAY IS:::")
	for i := n - 1; i >= 0; i-- {
		fmt.Println(Arr[i])
	}
}
