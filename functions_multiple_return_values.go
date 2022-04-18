package main

import (
	"errors"
	"fmt"
)

func divide(num1, num2 int) (int, error) {
	if num2 == 0 || num1 == 0 {
		return 0, errors.New("division by zero is not allowed")
	} else {
		return (num1 / num2), nil
	}
}

func multipleReturnValues() {
	if result, err := divide(10, 2); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
