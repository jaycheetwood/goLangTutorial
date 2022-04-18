package main

import (
	"fmt"
)

var (
	Name1 = "global1"
	Name2 = "global2"
)

var var1, var2 = "local1", "local2"

//const maxValue //compile error

const population int = 10000

func stringVariables() {
	var name1 string
	var name2 = "CloudAcademy"
	name3 := "Devops 2022"
	name4 := name2 + " " + name3

	fmt.Println(name1)
	name1 = "blah"
	fmt.Println(name1)

	fmt.Println(name2)
	fmt.Println(name3)
	fmt.Println(name4)

	fmt.Printf("%v -- %v\n", name2, name3)

	fmt.Println(Name1, Name2, var1, var2)
}

func intVariables() {
	var count1 int
	var count2 = 100
	count3 := 300
	count4 := count2 + count3

	fmt.Println(count1) //prints 0
	count1 = 10
	fmt.Println(count1) //prints 10

	fmt.Println(count2)
	fmt.Println(count3)
	fmt.Println(count4)

	count3++
	count4--

	fmt.Println(count3)
	fmt.Println(count4)
}

func booleanVariables() {
	var isBig bool //defaults to false
	var isFast, hasFourWheels = false, true

	fmt.Println(isBig)
	fmt.Println(isFast)
	fmt.Println(hasFourWheels)

	fmt.Println(!hasFourWheels)
	fmt.Println(isFast && hasFourWheels)
	fmt.Println(isFast || hasFourWheels)

	if !isFast {
		fmt.Println("It's slow!")
	}

	if hasFourWheels {
		fmt.Println("It's a car!")
	}
}

func constVariables() {
	const name = "CloudAcademy"

	//name = "Blah" //compile error

	if true {
		const color = "red"
		fmt.Println(population)
		fmt.Println(color)
	}
	//fmt.Println(color) //compile error
}

func main() {
	stringVariables()
	intVariables()
	booleanVariables()
	constVariables()
	controlFlow()
	controlFlowFor()
	controlFlowSwitch()
	arrayTutorial()
	slicesTutorial()
	mapTutorial()
	rangeTutorial()
	mathFunction()
	variadicFunctions()
	multipleReturnValues()
	variableFunctions()
	structDataTypes()
	pointerTutorial()
	pointersWhy()
	methodStructs()
	printMethod()
	appleInterfaces()
	goErrors()
	deferTutorial()
	panicRecovery()
	waitGroups()
}
