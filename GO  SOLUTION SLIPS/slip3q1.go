package main

import "fmt"

func main() {
	var num, temp, rev, rem int
	fmt.Println("ENTER NUMBER TO CHECK?")
	fmt.Scanln(&num)
	temp = num
	for num > 0 {
		rem = num % 10
		rev = rev*10 + rem
		num = num / 10
	}
	if temp == rev {
		fmt.Println(temp, " IS PALINDROME")
	} else {
		fmt.Println(temp, " IS NOT PALINDROME")
	}
}
