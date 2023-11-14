package main

import "fmt"

func main() {
	var num, i int
	fmt.Println("Enter Number TO Find Factors::::")
	fmt.Scanln(&num)
	fmt.Println(("Factors Are:::"))
	for i = 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)
		}
	}

}
