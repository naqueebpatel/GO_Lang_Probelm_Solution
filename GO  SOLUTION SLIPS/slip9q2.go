package main

import "fmt"

func main() {
	var i, j, cnt int
	fmt.Println("PRIME NUMBERS BETWEEN 1-100:::")
	for i = 1; i <= 100; i++ {
		cnt = 0
		for j = 1; j <= i; j++ {
			if i%j == 0 {
				cnt += 1
			}
		}
		if cnt == 2 {
			fmt.Println(i)
		}

	}

}
