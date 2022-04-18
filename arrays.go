package main

import "fmt"

func arrayTutorial() {
	var arr1 [4]int                    //init with 0s
	arr2 := [...]int{3, 1, 5, 10, 100} //compiler infers size of array
	fmt.Println(arr1)
	fmt.Println(arr2)

	fmt.Println(len(arr1)) //4
	fmt.Println(len(arr2)) //5

	for _, value := range arr2 {
		fmt.Println(value)
	}

	arr1[0] = 10 //setting
	// arr1[4] = 1 //compile error - out of bounds/index doesnt exist

	fmt.Println(arr1)
	fmt.Println(arr1[0]) //getting

	//Multidimensional
	arr3 := [2][2]string{
		{"a1", "a2"},
		{"b1", "b2"},
	}

	fmt.Println(arr3)
}
