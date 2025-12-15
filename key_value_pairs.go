package main

import "fmt";

func key_value_pairs() {
	fruits := map[string]string{
		"apple":  "Red",
		"banana": "Yellow",
		"grape":  "Purple",
	}

	fmt.Println("colors of fruits:", fruits["banana"])

	fruits["banana"] = "Green"
	fmt.Println("Updated color of banana:", fruits["banana"])
}