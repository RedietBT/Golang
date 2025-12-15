package main

import 
  "fmt"

  type Persons struct {
	Name string
	Age  int
  }

func UpdateAge(p *Persons, newAge int) {
	p.Age = newAge
}

func pointer() {
	maria:= Persons{Name: "Maria", Age: 25}
	fmt.Println("Before update:", maria)
	UpdateAge(&maria, 30)
	fmt.Println("After update:", maria)
}