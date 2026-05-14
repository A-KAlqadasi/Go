package main

import "fmt"

func main() {

	fmt.Println("Profit Calculator")
	var revenue float64 
	var expenses float64
	var taxRate float64

	revenue = readInput("Enter revenue: ")
	expenses = readInput("Enter expenses: ")
	taxRate = readInput("Enter tax rate: ")

	// profitBeforeTax := revenue - expenses
	// taxAmount := profitBeforeTax * (taxRate / 100)
	// profitAfterTax := profitBeforeTax - taxAmount
	profitBeforeTax, taxAmount, profitAfterTax := calculateProfits(revenue, expenses, taxRate)
	
	printOutput("Tax: ", taxAmount)
	printOutput("Profit Before Tax: ", profitBeforeTax)
	printOutput("Profit After Tax: ", profitAfterTax)

}

func calculateProfits(revenue float64, expenses float64, taxRate float64) (float64, float64, float64) {
	profitBeforeTax := revenue - expenses
	taxAmount := profitBeforeTax * (taxRate / 100)
	profitAfterTax := profitBeforeTax - taxAmount

	return profitBeforeTax, taxAmount, profitAfterTax
}

func readInput(prompt string) float64 {
	var value float64
	fmt.Print(prompt)
	fmt.Scanln(&value)
	return value
}

func printOutput(message string, value float64) {
	fmt.Printf(message + "%v \n", value);
}