package main

import "fmt"

func main() {
	var a, b string
	fmt.Println("Enter first string")
	fmt.Scanln(&a)
	fmt.Println("Enter second string")
	fmt.Scanln(&b)
	if a == b {
		fmt.Println("Is Equal String")
	} else {
		fmt.Println("Is Not Equal String")
	}
}
