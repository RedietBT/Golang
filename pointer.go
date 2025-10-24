package main

import 
  "fmt"

  type Person struct {
	Name string
	Age  int
  }

func UpdateAge(p *Person, newAge int) {
	p.Age = newAge
}

func pointer() {
	maria:= Person{Name: "Maria", Age: 25}
	fmt.Println("Before update:", maria)
	UpdateAge(&maria, 30)
	fmt.Println("After update:", maria)
}