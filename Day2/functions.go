package main

import "fmt"

func myFunc() {
	fmt.Println("This is a function in go")
}

func add(num1 int, num2 int) int {
	return num1 + num2
}

func mutipleReturnValue() (int, string) {
	return 10, "Messi"
}

func callFunc(callable func(int) int) int {
	return callable(10)
}

func doubleNumber(num1 int) int {
	return num1 * 2
}

// Function returning a function
func greet(str1 string) func(str2 string) string {
	return func(str2 string) string {
		return str1 + str2
	}
}

// Function with multiple parameters
func sum(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func functions() {
	fmt.Println("Functions in Go")
	// This or we can write a anonymous function
	v1 := callFunc(doubleNumber)
	fmt.Println(v1)

	// Below is the anonymous function
	v2 := callFunc(func(x int) int {
		return x * 10
	})
	fmt.Println(v2)
	greeting := greet("Hello, ")
	greetingLeo := greeting("Leo")
	fmt.Println(greetingLeo)
	fmt.Println(sum(10, 20, 30, 40, 50))
}
