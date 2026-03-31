package main

import "fmt"

func datatypes() {
	// both values positive and negative integers
	// int

	// uint only positive integers

	// byte type is simply but int8 as byte = 8bits

	// rune type stores the int32 value

	// bool // true or false

	// string anything in between " "

	// nil type is undefined or null

	// This is how you declare a variable
	var x string = "Hello, World"
	var y uint32 = 10

	fmt.Println(x, y)

	const name string = "Bhushan"
	const PI float32 = 3.14

	// Implicit declaration
	z := int(3)
	fmt.Println(z)
	// The below statement prints the type of variable
	fmt.Printf("%T", z)
	// Typecasting
	w := int16(z)
	fmt.Printf("%T, %T\n", z, w)
}
