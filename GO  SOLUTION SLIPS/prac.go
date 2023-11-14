package main

import "fmt"

func main() {
	var Arr [50]int
	var size int
	fmt.Println("ENTER THE NUMBER OF ELEMENT")
	fmt.Scanln(&size)
	for i := 0; i < size; i++ {
		fmt.Println("ENTER ELEMENT")
		fmt.Scanln(&Arr[i])
	}
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if Arr[i] > Arr[j] {
				temp := Arr[i]
				Arr[i] = Arr[j]
				Arr[j] = temp
			}
		}
	}
	fmt.Println("MAXIMUM ELEMENT IS:::", Arr[size-1])
}
