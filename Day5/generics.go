package main

import "fmt"

// Instead of writing your unions at the function level, make them as interface and use them
type Number interface {
	int | float32 | uint
}

// function statement would look like

func addMultiple[T Number](a T, b T) T {
	sum := a + b
	return sum
}

// Basic add function that only works with int
func add(a int, b int) int {
	return a + b
}

// Creating a new function using generics
// This is how you define a generic in go
func modifiedAdd[T int | float64 | uint](x T, y T) T {
	return x + y
}

func getValues[K comparable, V any](mp map[K]V) []V {
	values := []V{}

	for _, value := range mp {
		values = append(values, value)
	}

	return values
}

// Creating a Generic slice

// This is the syntax for creatnig a generic slice if you want slice for some datatypes
// You can replace the any with unions of datatype
// for example type IntFloatSlice[T int | float32 | uint ] []T
type GenericSlice[T any] []T

type GenericStruct[T any] struct {
	values T
}

func (g GenericSlice[T]) Print() {
	for _, value := range g {
		fmt.Println(value)
	}
}

func generics() {
	fmt.Println("Generics in go...")
	// By using generics we can give different types to a function
	value := modifiedAdd(10, 20)
	value1 := modifiedAdd(3.14, 4.12)
	value2 := modifiedAdd(uint(5), uint(4))
	fmt.Println(value, value1, value2)
	simpleMap := map[string]int{"a": 1, "b": 2, "c": 3, "d": 4}
	fmt.Println("Value from map are : ", getValues(simpleMap))
	genericSlice := GenericSlice[int]{10, 20, 30}
	genericSlice.Print()
	stringSliceUsingGeneric := GenericSlice[string]{"cat", "dog", "cow", "pig"}
	stringSliceUsingGeneric.Print()
}
