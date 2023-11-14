package main

import "fmt"

func main() {
	var str string
	var v, c, i int
	fmt.Println("ENTER STRING TO COUNT")
	fmt.Scanln(&str)
	for i = 0; i < len(str); i++ {
		if str[i] == 'A' || str[i] == 'a' || str[i] == 'E' || str[i] == 'e' || str[i] == 'I' || str[i] == 'i' || str[i] == 'O' || str[i] == 'o' || str[i] == 'U' || str[i] == 'u' {
			v += 1
		} else {
			c += 1
		}
	}
	fmt.Println("THE STRING IS:::", str)
	fmt.Println("VOWEL COUNT:::", v)
	fmt.Println("CONSONANT COUNT:::", c)
}
