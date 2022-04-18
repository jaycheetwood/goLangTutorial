package main

import "fmt"

func rangeTutorial() {
	bytes := []int{67, 108, 111, 117, 100, 65, 99, 97, 100, 101, 109, 121}
	for idx, value := range bytes {
		fmt.Print(string(value))
		if len(bytes)-1 == idx {
			fmt.Println()
		}
	}
	team := map[string]player{"John": {3, 10}, "Bob": {14, 11}}
	fmt.Println(team)

	for key, value := range team {
		fmt.Printf("%s -> %d\n", key, value)
	}

	for _, value := range team {
		fmt.Println(value)
	}

	for key := range team {
		fmt.Println("key:", key)
	}
}
