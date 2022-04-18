package main

import "fmt"

func notString(msg string) {
	msg = fmt.Sprintf("not%s", msg)
}

func notStringPtr(msg *string) {
	*msg = fmt.Sprintf("not%s", *msg)
}

func pointersWhy() {
	message := "cloudacademy"

	notString(message)
	fmt.Println(message)

	notStringPtr(&message)
	fmt.Println(message)
}
