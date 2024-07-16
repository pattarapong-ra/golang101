package main

import (
	"testing"
)
 // ตัวอย่างการเขียน test ของ function ที่เขียนไว้ใน main.go
func TestCalculateTotalMoneyInTeam(t *testing.T) {
	//arrange
	team := []Team{
		{"Ham", 5.75},
		{"Golf", 3.50},
		{"Parn", 4.50},
		{"Muiewju", 4.25},
		{"Petch", 100.00},
		{"Pack", 5.50},
	}
	expectedTotal := 123.50
	//action
	actualTotal := calculateTotalMoneyInTeam(team)

	//assertion
	if actualTotal != expectedTotal {
		t.Errorf("Expected total %f but got %f", expectedTotal, actualTotal)
	}
}

func TestCheckIfEnoughMoney(t *testing.T) {
	total := 19.50
	goalAmount := 15.00
	expected := false
	actual := checkIfEnoughMoney(total, goalAmount)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
	total = 10.00
	expected = false
	actual = checkIfEnoughMoney(total, goalAmount)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
