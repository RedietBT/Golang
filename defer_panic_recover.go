package main

import "fmt"

func defer_panic_recover() {
	defer func() {
		if r:= recover(); r != nil {
			fmt.Println("Reecover from panic", r)
		}
	}()

	fmt.Println("Before panic")
	panic("crirtcal error occured!")
	fmt.Println("This will not execute")

}