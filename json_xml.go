package main
import (
	"fmt"
	"encoding/json"
	"encoding/xml"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

type BookXML struct {
	Title  string `xml:"title"`
	Author string `xml:"author"`
	Year   int    `xml:"year"`
}

func jsons() {
	data := `{"name": "Alice", "age": 30, "email": "alice@example.com"}`

	var person Person
	err := json.Unmarshal([]byte(data), &person)

	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	fmt.Println("JSON DATA.")
	fmt.Printf("Name: %s, Age: %d, Email: %s\n", person.Name, person.Age, person.Email)
}

func xmls() {
	xmldata := `<Book>
	<title>The Go Programming Language</title>
	<author>Alan A. A. Donovan</author>
	<year>2015</year>
</Book>`

	var book BookXML
	err := xml.Unmarshal([]byte(xmldata), &book)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Parsed XML DATA.")
	fmt.Printf("Title: %s, Author: %s, Year: %d\n", book.Title, book.Author, book.Year)
	encodedData, err := xml.MarshalIndent(book, "", " ")

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Re-encoded XML DATA.")
	fmt.Println(string(encodedData))
}