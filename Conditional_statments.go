package main

import "fmt"

func Conditional_statments() {
var day int = 3
switch day {
case 1:
	fmt.Println("Monday")
case 2:
	fmt.Println("Tuesday")
case 3:
	fmt.Println("Wednesday")
default:
	fmt.Println("Invalid day")
}	
}