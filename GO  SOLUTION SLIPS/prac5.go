package main

import "fmt"

func main() {
	var v1, v2, v3 int
	var Arr [10]int
	v1 = 0
	v2 = 1
	Arr[0] = v1
	Arr[1] = v2
	for i := 2; i < 10; i++ {
		v3 = v1 + v2
		Arr[i] = v3
		v1 = v2
		v2 = v3
	}
	fmt.Println("FIBONACCI ARRAY ELEMENT")
	for i := 0; i < 10; i++ {
		fmt.Println(Arr[i])
	}
}
