package main

import "fmt"

func main() {
	var v1, v2, v3, num int
	fmt.Println("ENTER THE NUMBER OF TERMS")
	fmt.Scanln(&num)
	v1 = 0
	v2 = 1
	fmt.Println(v1)
	fmt.Println(v2)
	for i := 2; i < num; i++ {
		v3 = v1 + v2
		fmt.Println(v3)
		v1 = v2
		v2 = v3
	}
}
