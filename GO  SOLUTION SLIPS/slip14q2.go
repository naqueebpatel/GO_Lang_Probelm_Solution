package main

import "fmt"

type Book struct {
	bid, price    int
	title, author string
}

func (b *Book) Accept() {
	fmt.Println("ENTER BOOK ID")
	fmt.Scanln(&b.bid)
	fmt.Println("ENTER BOOK TITLE")
	fmt.Scanln(&b.title)
	fmt.Println("ENTER BOOK AUTHOR")
	fmt.Scanln(&b.author)
	fmt.Println("ENTER BOOK PRICE")
	fmt.Scanln(&b.price)
}
func (b *Book) Display() {
	fmt.Println("BOOK ID:::", b.bid)
	fmt.Println("BOOK TITLE:::", b.title)
	fmt.Println("BOOK AUTHOR:::", b.author)
	fmt.Println("BOOK PRICE:::", b.price)
}

func main() {
	var size, i int
	obj := make([]Book, 100)
	fmt.Println("ENTER NO OF DETAILS")
	fmt.Scanln(&size)
	for i = 0; i < size; i++ {
		obj[i].Accept()
	}
	for i = 0; i < size; i++ {
		obj[i].Display()
	}
}
