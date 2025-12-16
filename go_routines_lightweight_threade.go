package main
import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("Task: %s - Step %d\n", name, i)
		time.Sleep(300 * time.Millisecond)
	}
}

func go_routines() {
	go task("A")
	go task("B")

	time.Sleep(2 * time.Second)
	fmt.Println("All tasks completed.")
}