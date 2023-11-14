package main

import "fmt"

func main() {
	var num int
	fmt.Println("ENTER NO TO CHECK")
	fmt.Scanln(&num)
	if num%2 == 0 {
		fmt.Printf("%d IS EVEN", num)
	} else {
		fmt.Printf("%d IS ODD", num)
	}
}
