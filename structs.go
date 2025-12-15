package main
import "fmt"

type Book struct {
	Title string
	Auther  string
	Page int
	IsPublished bool
}

func structs(){
	book1 := Book{
		Title: "The Go Programming Language",
		Auther: "Alan A. A. Donovan",
		Page: 450,
		IsPublished: true,
	}
	fmt.Println("book:", book1)
}