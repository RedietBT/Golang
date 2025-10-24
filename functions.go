package main

import "fmt"

func multiply(a int, b int) int {
	return a * b
}

func function() {
	result:= multiply(5, 10)
	fmt.Println("The product is:", result)
}