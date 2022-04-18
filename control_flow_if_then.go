package main

import "fmt"

func controlFlow() {
	var population int = 5500

	if population < 5000 {
		fmt.Println("small")
	} else if population >= 5000 && population < 7000 {
		fmt.Println("medium")
	} else {
		fmt.Println("large")
	}

	if toggle := true; toggle {
		fmt.Println("Its on!")
	}
}
