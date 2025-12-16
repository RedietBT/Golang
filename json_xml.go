package main
import (
	"fmt"
	"encoding/json"
)

type Person struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Email string `json:"email"`
}

func main() {
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