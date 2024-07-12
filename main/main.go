package main

import (
	"fmt"
)

// Define a struct to represent a member's money in the team
type Team struct {
	Name   string
	Amount float64
}

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

func main() {
	// Create an array of team name
	team := []Team{
		{"Ham", 5.75},
		{"Golf", 3.50},
		{"Parn", 4.50},
		{"Muiewju", 4.25},
		{"Petch", 100.00},
		{"Pack", 5.50},
	}
	// Define the goal amount to buy the item
	var coffeePrice float64 = 15.00
	// Calculate the total contributions
	total := calculateTotalMoneyInTeam(team)
	// Check if the total contributions are enough to meet the goal amount
	enoughMoney := checkIfEnoughMoney(total, coffeePrice)
	// Print the result
	fmt.Println("Is the total amount enough?", enoughMoney)

	isEnough, change := buyCoffee(total, coffeePrice)
	if !isEnough {
		fmt.Println("Not enough money.")
		return
	} else {
		fmt.Printf("Change: $%.2f\n", change)
	}
}

func buyCoffee(total float64, price float64) (bool, float64) {
	var isEnough bool
	var change float64
	if total < price {
		isEnough, change = false, 0
	} else {
		isEnough, change = true, total-price
	}
	return isEnough, change
}
