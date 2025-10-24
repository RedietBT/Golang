package main

import "fmt"

func main() {
	var age, isMember int = 25, 1

	fmt.Println("Eligibile for Discount:", age > 18 && isMember == 1)
	fmt.Println("Underage or Non-Member:", age <= 18 || isMember != 1)
	fmt.Println("Not a Member:", !(isMember == 1))
}
