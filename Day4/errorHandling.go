package main

import (
	"errors"
	"fmt"
)

func upgradedDivide(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by 0")
	} else {
		return a / b, nil
	}
}

func divide(a int, b int) int {
	return a / b
}

func defferedFunction() {
	fmt.Println("Deffered function")
	// If a panic occurs we will go inside of the deffered function and execute all the cleanup operations
	// recover helps you not to crash a program
	// You should always right your recover function inside a deffered function
	r := recover()
	fmt.Println(r)
}

func ErrorHandling() {
	// The defer keyword will delay the execution of the function, i.e after everything is ran
	// defer keyword statement will run
	// this is like .finally in javascript
	// You defer the functions that cleanup your code, it will not run after a panic is occured
	defer defferedFunction()
	// Panic statement, anything after the panic statement will not run
	// divide(1, 0)
	result, error := upgradedDivide(1, 0)
	fmt.Println(result, error)
	// panic("Shit... this crashed")
	fmt.Println("run")
	// fmt.Println(divide(1, 0))
}
