package main

import "fmt"

func main() {
	var yr int
	fmt.Println("ENTER YEAR TO CHECK?")
	fmt.Scanln(&yr)
	if yr%4 == 0 {
		fmt.Println(yr, " IS A LEAP YEAR.")
	} else {
		fmt.Println(yr, " IS NOT A LEAP YEAR.")
	}
}
