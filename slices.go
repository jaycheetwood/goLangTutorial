package main

import "fmt"

func slicesTutorial() {
	arr1 := [6]int{1, 3, 5, 7, 11, 13}
	slice1 := []int{1, 3, 5, 7, 11, 13}
	slice2 := slice1[1:2]
	var slice3 = make([]int, 2, 3)

	fmt.Println(arr1)
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(slice3)

	slice3 = slice1[1:4]

	fmt.Println(slice3)
	fmt.Println(len(slice3))

	slice1 = append(slice1, 200, 300, 400)
	fmt.Println(slice1)

	slice2 = append(slice2, []int{7, 8, 9}...)
	fmt.Println(slice2)

	copyslice := make([]int, len(slice1))
	copy(copyslice, slice1)
	fmt.Println(copyslice)
}
