package main

import "fmt"

func Unique_Loops() {
    names := []string{"Alice", "Bob", "Charlie"}
	
	for index, name := range names {
		fmt.Printf("Index: %d, Name: %s\n", index, name)
	}
}