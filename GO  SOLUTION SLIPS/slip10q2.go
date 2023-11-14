package main

import "fmt"

func main() {
	var i, j, num, c, s int
	fmt.Println("ENTER NO OF ROWS TO PRINT PASCALS TRIANGLE?")
	fmt.Scanln(&num)
	for i = 0; i < num; i++ {
		for s = 1; s < num-i; s++ {
			fmt.Printf(" ")
		}
		for j = 0; j < 0; j++ {
			if i == 0 || j == 0 {
				c = 1
			} else {
				c = (c * (i - j + 1)) / j
			}
			fmt.Printf("%d", c)
		}
		fmt.Println()
	}
}
