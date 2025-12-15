package main

import "fmt"

type car struct {
	Brand string
	Model string
	Year  int
}

func (c car) updateYear(newYear int) {
	fmt.Printf("Car: %s %s, Year: %d\n", c.Brand, c.Model, c.Year)
}
func struct_methods_oop() {

	car := car {
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
	}

	fmt.Println("Before update:", car)
	car.updateYear(2022)
	fmt.Println("After update:", car)
}