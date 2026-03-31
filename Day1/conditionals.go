package main

import "fmt"

func conditionals() {

	fmt.Println("Conditonals in Go")
	a := 10
	switch a {
	case 10:
		fmt.Println("Hurrayyyyyy")
	case 2:
		fmt.Println("Badluck....")
	case 3:
		fmt.Println("Shitttttt")
	default:
		fmt.Println("Fuckeeeeeedddddddd")
	}

	if a > 10 {
		fmt.Println("Less than 20")
	} else if a <= 10 {
		fmt.Println("In range")
	} else {
		fmt.Println("Na")
	}
}
