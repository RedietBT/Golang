package main
import "fmt"

func Loops() {
	numbers := []int{10, 20, 30, 40, 50}

	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
