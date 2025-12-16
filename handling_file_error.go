package main

import (
	"fmt"
	"os"
)

func handling_file_error() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	fmt.Println("File created and written successfully")
}