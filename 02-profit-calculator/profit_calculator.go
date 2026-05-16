package main

import (
	"fmt"
	"os"
)

func main() {

	fmt.Println("Profit Calculator")
	var revenue float64 
	var expenses float64
	var taxRate float64

	revenue = readInput("Enter revenue: ")
	if !isValidInput(revenue) {
		fmt.Println("Invalid input for revenue. Please enter a positive number.")
		return
	}
	expenses = readInput("Enter expenses: ")
	if !isValidInput(expenses) {
		fmt.Println("Invalid input for expenses. Please enter a positive number.")
		return
	}

	taxRate = readInput("Enter tax rate: ")
	if !isValidInput(taxRate) {
		fmt.Println("Invalid input for tax rate. Please enter a positive number.")
		return
	}

	profitBeforeTax, taxAmount, profitAfterTax := calculateProfits(revenue, expenses, taxRate)
	printOutput("Tax: ", taxAmount)
	printOutput("Profit Before Tax: ", profitBeforeTax)
	printOutput("Profit After Tax: ", profitAfterTax)
	
	writeProfitsToFile(profitBeforeTax, taxAmount, profitAfterTax)

}

func writeProfitsToFile(profitBeforeTax float64, taxAmount float64, profitAfterTax float64) {
	content := fmt.Sprintf("Profit Before Tax: %v\nTax Amount: %v\nProfit After Tax: %v\n", profitBeforeTax, taxAmount, profitAfterTax)

	err :=os.WriteFile("profits.txt", []byte(content), 0644)
	if err != nil {
		panic("Failed to write to file")
	}

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

func isValidInput(value float64) bool {
	return value > 0
}

func printOutput(message string, value float64) {
	fmt.Printf(message + "%v \n", value);
}