package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type travels struct {
	id                 string
	passangerDetails   passanger
	destinationDetails destination
}

type passanger struct {
	name    string
	id      string
	age     int
	address string
}

type destination struct {
	name            string
	id              string
	continent       string
	fare            string
	destinationCode string
}

var mockDestinations = []destination{{
	name:            "Malaysia",
	id:              "1234",
	continent:       "Asia",
	fare:            "70,000INR",
	destinationCode: "MYR",
}, {
	name:            "China",
	id:              "5678",
	continent:       "Asia",
	fare:            "80,000INR",
	destinationCode: "RMB",
}, {
	name:            "Singapore",
	id:              "9101",
	continent:       "Asia",
	fare:            "90,000INR",
	destinationCode: "SGD",
}}

func checkIfEnteredDestinationAvailable(code string) bool {
	for _, value := range mockDestinations {
		if strings.EqualFold(code, value.destinationCode) {
			return true
		}
	}
	return false
}

func BookingSystem() {
	fmt.Println("Welcome to leo travels")
	fmt.Println("Please see below the list of destinations")
	fmt.Println()
	for _, value := range mockDestinations {
		fmt.Printf("Name : %s\nState : %s\nFare : %s\n", value.name, value.continent, value.fare)
		fmt.Println()
	}
	var option string = "No"
	fmt.Print("Please enter the destination code :  ")
	code, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	code = strings.TrimSpace(code)
	exists := checkIfEnteredDestinationAvailable(code)
	fmt.Print("To proceed with your booking, choose Y/N : ")
	fmt.Scanln(&option)
	if strings.ToLower(option) == "y" && exists {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter the passanger details: ")
		fmt.Print("Enter name: ")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		fmt.Print("Enter age: ")
		str, _ := reader.ReadString('\n')
		age, _ := strconv.Atoi(strings.TrimSpace(str))
		fmt.Print("Enter address: ")
		address, _ := reader.ReadString('\n')
		address = strings.TrimSpace(address)
		fmt.Println(age)
		os.Exit(0)
	} else {
		fmt.Println("Destination you entered is not available currently")
		os.Exit(0)
	}

}
