package main

import "fmt"

func main() {
	var revenue float64
	var expenses float64
	var taxRate float64

	revenue = GetUserInput("Enter Revenue : ")

	expenses = GetUserInput("Enter Expenses : ")

	taxRate = GetUserInput("Enter Tax Rate : ")

	Calculate(revenue, expenses, taxRate)

}
func GetUserInput(infoText string) float64 {
	var inputUser float64
	fmt.Print(infoText)
	fmt.Scan(&inputUser)
	return inputUser
}

func Calculate(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	fmt.Println("The Earnings Before Tax (EBT): ", ebt)
	profit = ebt * (1 - taxRate/100)
	fmt.Println("The Earnings After Tax (Profit): ", profit)

	ratio = ebt / profit
	fmt.Println("the Ratio :", ratio)
	return ebt, profit, ratio
}
