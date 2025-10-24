package main

import (
"fmt"
)

func devide (a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("division errror: cannot divide %d by %d", a, b)
	}
	return a / b, nil
}

func Error_handling() {
	retsult, err := devide(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", retsult)
	}
}