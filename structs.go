package main

import "fmt"

type Book struct{
	title, author string
	pages int
}

func main(){

	book:= Book{"Go101","Tapir",256}

	fmt.Println(book)

	book = Book{}

	book = Book{author: "tapir"}

	var book2 Book

	book2.title = "Go101"
	book2.pages = 300

	fmt.Println(book2.pages)
}