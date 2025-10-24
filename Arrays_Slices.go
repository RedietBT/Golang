package main

import (
	"fmt"
	
)

func main() {

	arr := [3]int{1, 2, 3}
	fmt.Println("Array:", arr)

	slice := arr[:]
	slice = append(slice, 4)

	fmt.Println("Slice:", slice)
	fmt.Println("Array after appending to slice:", arr)
}