package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var taxHate float64

	revenue = getUserInfo("Enter the expected revenue: ")

	expenses = getUserInfo("Enter the expected expenses: ")

	taxHate = getUserInfo("Enter the expected taxHate: ")

	fmt.Println("------------------------------------------------------------")

	profit, profitWithoutTax, profitRatio := calculateValues(revenue, expenses, taxHate)

	fmt.Println("this is the true profit: ", profit)
	fmt.Println("this is the profit without taxes: ", profitWithoutTax)
	fmt.Println("this is the ratio true profit and profit without taxes: ", profitRatio)
}

func getUserInfo(label string) (value float64) {
	fmt.Print(label)
	fmt.Scan(&value)

	return value
}

func calculateValues(revenue, expenses, taxHate float64) (profit, profitWithoutTax, profitRatio float64) {
	profit = (revenue - expenses) * (1 - taxHate/100)
	profitWithoutTax = revenue - expenses
	profitRatio = profit / profitWithoutTax

	return profit, profitWithoutTax, profitRatio
}
