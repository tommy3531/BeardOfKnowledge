package controllers

import (
	"fmt"
)

func GetCurrentSenators() {
	fmt.Println("Get ALL Senators")
}

func GetCurrentHouseMembers() {
	// 102-115
	fmt.Println("Get Current House Memebers: (102-115)")
}

func GetNewMembers() {
	fmt.Println("Display New Members of Congress")
}

func GetCurrentMemberByState() {
	fmt.Println("Get Member by State: Need Two Digit State")
}

func GetMembersLeavingOffice() {
	fmt.Println("Get Member Leaving Office: ")
}

func GetMemberExpenses() {
	// leg_id, year, quarter
	fmt.Println("Get Member Expenses by year and quarter ")
}

func GetMemberExpensesByCategory() {
	// leg_id
	// Category -> travel, personnel, rent-utilities, other-services, supplies, franked-mail, printing, equipment, total
	fmt.Println("Get Expenses By Category")
}

