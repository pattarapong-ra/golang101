package main

import (
	"testing"
)

// Test for calculateTotalContributions function
// 1. เปิดให้ดู
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
	//+black box concept ว่า ทำไมเราต้องเตรียมของที่ใช้เทสใหม่

	//action
	actualTotal := calculateTotalMoneyInTeam(team)

	//assertion
	if actualTotal != expectedTotal {
		t.Errorf("Expected total %f but got %f", expectedTotal, actualTotal)
	}
} // ตัวอย่างการเขียน test ของ function ที่เขียนไว้ใน main.go

// Test for checkIfEnoughMoney function
// 2. เขียนให้ดู ไปพร้อมๆกับน้อง
// - น้องมีส่วนร่วมกับการทำ unit test
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

// แบ่งกลุ่มน้อง 15 คน เป็น 3 กลุ่ม
// ให้ในแต่ละกลุ่มเป็น BA รับ requirement 2 คน เป็น dev 3 คน
// บอกโจทย์ให้น้อง BA
