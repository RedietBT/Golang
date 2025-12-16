package main
import (
	"fmt"
	"io/ioutil"
)

func reading_and_wrighting_file() {

	data, err := ioutil.ReadFile("example.txt")


	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("File content:", string(data))
}

