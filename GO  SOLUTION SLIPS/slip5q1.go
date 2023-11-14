package main

import "fmt"

func main() {
	var num, i int
	fmt.Println("ENTER NUMBER ?")
	fmt.Scanln(&num)
	for i = 1; i <= 10; i++ {
		fmt.Println(num, "x", i, "=", num*i)
	}
}
