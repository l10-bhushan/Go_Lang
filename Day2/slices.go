package main

import "fmt"

func slices() {
	fmt.Println("Slices in go")

	arr := [5]int{1, 2, 3, 4, 5}
	s1 := arr[:3]
	fmt.Println(s1)
	// When we modify the slice the original array also changes
	s1[0] = 100
	fmt.Println(s1, arr)

	// Slices have 3 parameter pointer , capacity , lenght
	fmt.Println(s1, len(s1), cap(s1))
	// Here, you will see the len is 3 but the capacity is 5 because original
	// Array size is 5
	// The capacity also depends on pointer position
	s1 = arr[1:3]
	fmt.Println(s1, len(s1), cap(s1))
	// In the above case cap will return 4 because the pointer for s1
	// starts from 1st position of arr instead of 0th position

	// Dynamic slices, flexible type
	s2 := []string{"hello", "world"}
	// Before appending
	fmt.Println(s2, len(s2), cap(s2))

	// Appending to the slice
	s2 = append(s2, "Leo")
	// After appending

	fmt.Println(s2, len(s2), cap(s2))

	// Making slice using make() function

	// Creates a empty slice with len 10 and cap 10
	// s3 := make([]int, 10)

	for _, value := range s2 {
		fmt.Println(value)
	}

}
