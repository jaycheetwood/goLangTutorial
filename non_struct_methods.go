package main

import (
	"fmt"
	"strings"
)

type upstring string

func (msg upstring) up() string {
	s := string(msg)
	return strings.ToUpper(s)
}

func printMethod() {
	message := upstring("cloudacademy")
	fmt.Println(message.up())
}
