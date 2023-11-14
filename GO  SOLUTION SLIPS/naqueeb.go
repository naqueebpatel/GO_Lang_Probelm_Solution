package main

import "fmt"

func main() {
	var str string
	var slice = []string{"INDIA", "AMERICA", "USA"}
	fmt.Println("INSERTING VALUE INTO SLICE:::")
	fmt.Println("ENTER STRING TO SLICE")
	fmt.Scanln(&str)
	slice = append(slice, str)
	fmt.Println(slice)

	fmt.Println("COPY STRING B:::")
	b := make([]string, 5)
	copy(b, slice)
	fmt.Println(b)

	var index int
	fmt.Println("ENTER INDEX TO REMOVE ELEMENT")
	fmt.Scanln(&index)
	RemSlice := Remove(slice, index)
	fmt.Println(RemSlice)

	fmt.Println("PRINTING SLICE VALUES:::")
	j := 0
	for range slice {
		fmt.Println(slice[j])
		j++
	}
}
func Remove(s []string, i int) []string {
	return append(s[:i], s[i+1:]...)
}
