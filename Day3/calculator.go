package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var calculatorOptions = [5]string{"Add", "Subtract", "Multiply", "Divide", "Quit"}

func add(nums ...int) int {
	sum := 0
	for _, value := range nums {
		sum += value
	}
	return sum
}

func division(num1 int, num2 int) int {
	return num1 % num2
}

func substract(nums ...int) int {
	result := 0
	for _, value := range nums {
		result -= value
	}
	return result
}

func multiply(nums ...int) int {
	result := 0
	for _, value := range nums {
		result *= value
	}
	return result
}

func calculator() {
	// Building a calculator
	a := 0
	var nums []int = []int{}
	for a < 5 {
		fmt.Println("Welcome to our calculator")
		fmt.Println("Select any of the below options:")
		for i, value := range calculatorOptions {
			fmt.Printf("%d : %s", i+1, value)
			fmt.Println()
		}
		fmt.Println("Enter your option: ")
		fmt.Scanln(&a)
		fmt.Println("You choose option: ", a)
		fmt.Println("Enter your input (space-separated integers): ")

		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input := scanner.Text()
			for _, numStr := range strings.Fields(input) {
				num, _ := strconv.Atoi(numStr)
				nums = append(nums, num)
			}
		}
		fmt.Println("Input is : ", nums)
		switch a {
		case 1:
			fmt.Println(calculatorOptions[a-1])
			fmt.Println("Sum is : ", add(nums...))
		case 2:
			fmt.Println(calculatorOptions[a-1])
			fmt.Println("Result is : ", substract(nums...))
		case 3:
			fmt.Println(calculatorOptions[a-1])
			fmt.Println("Product is : ", multiply(nums...))
		case 4:
			fmt.Println(calculatorOptions[a-1])
			fmt.Println("Result is : ", division(nums[0], nums[1]))
		case 5:
			fmt.Println(calculatorOptions[a-1])
		default:
			fmt.Println("Incorrect option...")
			a = 5
		}
		fmt.Println()
	}
}
