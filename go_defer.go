package main

import "fmt"

func doSomething(msg string) {
	fmt.Println(msg)
}

func system() int {
	fmt.Println("System Started...")

	defer doSomething("cleanup")
	defer doSomething("stop")

	fmt.Println("System Finished!")

	return 1
}

func deferTutorial() {
	data := system()
	fmt.Println(data)
}
