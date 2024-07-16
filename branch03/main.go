package main

import (
	"fmt"
)

// Define a struct to represent a member's money in the team
type Team struct {
	Name   string
	Amount float64
}

func main() {
	// Create an array of team name
	team := []Team{
		{"Neptune", 22},
		{"FahApi", 54},
		{"JajahAmon", 34},
		{"BamChaya", 46},
		{"Obey", 34},
	}
	// Price of each size of coffee
	var tallCoffee float64 = 150.00
	var grandeCoffee float64 = 190.00
	var ventiCoffee float64 = 220.00

	fmt.Println("Team Members:", team)
	fmt.Println("Tall Coffee Price: ", tallCoffee)
	fmt.Println("Grande Coffee Price: ", grandeCoffee)
	fmt.Println("Venti Coffee Price: ", ventiCoffee)
	// code here
}

// buyCoffee function
// func buyCoffee() () {
// }

// for loop
// Function to calculate the total money in the team
func calculateTotalMoneyInTeam(team []Team) float64 {
	var total float64 = 0
	// Loop through each member in the array
	for _, member := range team {
		total += member.Amount
	}
	return total
}

// if check
// Function to check if the total money meets the goal amount
func checkIfEnoughMoney(total float64, goalAmount float64) bool {
	// Print the total amount collected
	fmt.Printf("Total Amount Collected: $%.2f\n", total)
	// Check if the total amount is enough to meet the goal
	if total >= goalAmount {
		fmt.Println("Congratulations! You have enough money to buy the item.")
		return true
	} else {
		fmt.Println("Sorry, you do not have enough money to buy the item.")
		return false
	}
}
