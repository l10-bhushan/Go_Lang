package main

import "fmt"

func loops() {

	fmt.Println("Loops in Go")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	a := 0
	for a < 5 {
		fmt.Println("Value of a : ", a)
		a++
	}

	str := "hello"
	// If we do this we will get the int value of 'h' instead of 'h'
	fmt.Println(str[0])
	// To fix this we need to convert it to string first
	fmt.Println(string(str[0]))

	for _, char := range str {
		fmt.Printf("%c\n", char)
	}
}
