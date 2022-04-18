package main

import "fmt"

type people struct {
	firstname string
	surname   string
	age       int
}

func (p *people) fullname() string {
	return p.firstname + " " + p.surname
}

func (p *people) canDrive() bool {
	if p.age >= 20 {
		return true
	} else {
		return false
	}
}

func (p *people) updateAge(newAge int) {
	p.age = newAge
}

func methodStructs() {
	people1 := people{"John", "Doe", 40}
	people2 := people{"Mark", "Doe", 19}

	fmt.Printf("%s can drive: %t\n", people1.fullname(), people1.canDrive())
	fmt.Printf("%s can drive: %t\n", people2.fullname(), people2.canDrive())

	people2.updateAge(people2.age + 1) //marks birthday
	fmt.Println(people2.age)

	fmt.Printf("%s can drive: %t\n", people2.fullname(), people2.canDrive())
}
