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

var travelsDatabase []travels = []travels{}

func checkIfEnteredDestinationAvailable(code string) bool {
	for _, value := range mockDestinations {
		if strings.EqualFold(code, value.destinationCode) {
			return true
		}
	}
	return false
}

func (p passanger) getUserDetails() {
	fmt.Println()
	fmt.Println("Passenger details :")
	fmt.Printf("Name: %s\nAge: %d\nAddress: %s\n", p.name, p.age, p.address)
	fmt.Println()
}

func readUserDetails(p passanger) passanger {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter the passanger details: ")
	fmt.Print("Enter name: ")
	name, _ := reader.ReadString('\n')
	p.name = strings.TrimSpace(name)
	fmt.Print("Enter age: ")
	str, _ := reader.ReadString('\n')
	age, _ := strconv.Atoi(strings.TrimSpace(str))
	p.age = age
	fmt.Print("Enter address: ")
	address, _ := reader.ReadString('\n')
	p.address = strings.TrimSpace(address)
	return p
}

func BookingSystem() {
	fmt.Println("Welcome to leo travels")
	fmt.Println("Please see below the list of destinations")
	fmt.Println()
	for _, value := range mockDestinations {
		fmt.Printf("Destination Code:%s\nName : %s\nState : %s\nFare : %s\n", value.destinationCode, value.name, value.continent, value.fare)
		fmt.Println()
	}
	var option string = "No"
	fmt.Print("Please enter the destination code :  ")
	code, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	code = strings.TrimSpace(code)
	exists := checkIfEnteredDestinationAvailable(code)
	if exists == false {
		fmt.Println("Destination not available... ")
		os.Exit(0)
	}
	fmt.Print("To proceed with your booking, choose Y/N : ")
	fmt.Scanln(&option)
	if strings.ToLower(option) == "y" && exists {
		var p passanger = passanger{}
		p = readUserDetails(p)
		p.getUserDetails()
		fmt.Println("To proceed, please select y/n : ")
		paymentConfirmation, _ := bufio.NewReader(os.Stdin).ReadString('\n')
		if strings.ToLower(paymentConfirmation) == "y" {
			t := travels{
				id:                 "1234456",
				passangerDetails:   p,
				destinationDetails: mockDestinations[0],
			}
			travelsDatabase = append(travelsDatabase, t)
			fmt.Println(travelsDatabase)
			os.Exit(0)
		} else {
			fmt.Println("No payment method selected...")
			os.Exit(0)
		}
		os.Exit(0)
	} else {
		fmt.Println("Destination you entered is not available currently...")
		os.Exit(0)
	}

}
