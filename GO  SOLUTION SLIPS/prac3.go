package main

import "fmt"

func main() {
	var Vowel [10]string
	var Const [10]string
	var str string
	var v, c int
	fmt.Println("ENTER STRING?")
	fmt.Scanln(&str)
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			continue
		} else if str[i] == 'A' || str[i] == 'E' || str[i] == 'I' || str[i] == 'O' || str[i] == 'U' {
			Vowel[v] = string(str[i])
			v += 1
		} else {
			Const[c] = string(str[i])
			c += 1
		}
	}
	fmt.Println("VOWEL ARRAY")
	for i := 0; i < v; i++ {
		fmt.Println(Vowel[i])
	}
	fmt.Println("CONSONANT ARRAY")
	for i := 0; i < c; i++ {
		fmt.Println(Const[i])
	}
}
