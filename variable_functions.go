package main

import "fmt"

func extend(val string) func() string {
	i := 0
	return func() string {
		i++
		return val[:i]
	}
}

func variableFunctions() {
	ca := "cloudacademy"

	word := extend(ca)

	for i := 0; i < len(ca); i++ {
		fmt.Println(word())
	}
}
