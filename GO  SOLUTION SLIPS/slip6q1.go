package main

import "fmt"

func main() {
	var ch string
	fmt.Println("ENTER CHARACTER?")
	fmt.Scanln(&ch)
	if ch == "A" || ch == "a" || ch == "E" || ch == "e" || ch == "I" || ch == "i" || ch == "O" || ch == "o" || ch == "U" || ch == "u" {
		fmt.Println(ch, "IS VOWEL")
	} else {
		fmt.Println(ch, " IS CONSONANT")
	}
}
