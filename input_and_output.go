package main

import "fmt";

func main() {
	var num1, num2 int
	fmt.Println("Enter two numbers suparated by space: ")
	fmt.Scanf("%d %d", &num1, &num2)
	fmt.Println("Sum, ", num1+num2)
}
