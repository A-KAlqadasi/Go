package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const fileName = "balance.txt"

func writeBalanceToFile(balance float64) {
	var byteBalance = fmt.Sprint(balance)
	err := os.WriteFile(fileName, []byte(byteBalance), 0644)
	if err != nil {
		fmt.Println("Error writing balance to file", err)
	}

}

func readBalanceFromFile() (float64, error) {
	var balance float64
	data, err := os.ReadFile(fileName)
	
	if err != nil {
		err = errors.New("Failed to read balance from file")
		return 0, err
	}

	balanceText := string(data)
	balance, _= strconv.ParseFloat(balanceText, 64)
	return  balance, nil
}

func main() {
	fmt.Println("============================")
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("============================")
	

	showMainMenu()

	fmt.Println("============================")
	fmt.Println("Thank you for using Go Bank! Goodbye!")
	fmt.Println("============================")
}

func showMainMenu() {

	var option int
	for option != 4 {
		
		balance, err := readBalanceFromFile()
		if err != nil {
			fmt.Println("Error: ", err)
			panic("Can't continue, sorry")
		}

		fmt.Println("\nWhat would you like to do?")
		fmt.Println("============================")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit amount")
		fmt.Println("3. Withdraw amount")
		fmt.Println("4. Exit")
		fmt.Println("============================")
		fmt.Printf("Enter your choise: ")
		fmt.Scan(&option);
		
		

		switch option {
		
		case 1:
			checkBalance(balance)
		case 2:
			{
				var depositAmount float64
				depositAmount = readAmount("Enter the amount you want to deposit: ")
				newBalance, isValidationPassed := depositMoney(balance, depositAmount)
				if(!isValidationPassed) {
					continue
				}
				fmt.Printf("\nYour new balance is: $%.2f\n", newBalance)
				writeBalanceToFile(newBalance)
			}
		case 3:
			var withdrawAmount float64
			withdrawAmount = readAmount("Enter the amount you want to withdraw: ")
			balance, isValidationPassed := withdrawMoney(balance, withdrawAmount)
			if(!isValidationPassed) {
				continue
			}
			fmt.Printf("\nYour new balance is: $%.2f\n", balance)
			writeBalanceToFile(balance)
		case 4:
			fmt.Println("\nExiting the application...")
		default:
			fmt.Println("\nInvalid option. Please try again.")
		}

	}
}
	

func readAmount(message string) float64 {
	fmt.Print(message)
	var amount float64
	fmt.Scan(&amount)
	return amount
}

func checkBalance(balance float64) {
	fmt.Printf("Your balance is: $%.2f\n", balance)
}

func depositMoney(balance float64, depositAmount float64) (float64, bool) {
	
	if (!isValidDepositAmount(depositAmount)) {
		return balance, false
	}

	return balance + depositAmount, true
}

func isValidDepositAmount(depositAmount float64) bool {
	if(depositAmount <=0 ) {
		fmt.Printf("\nInvalid deposit amount")
		return false
	}
	return true
}

func isValidWithdrawAmount (balance float64, withdrawAmount float64)  bool {
	if(withdrawAmount <=0 ) {
		fmt.Printf("\nInvalid withdraw amount")
		return false
	}

	if(withdrawAmount > balance) {
		fmt.Printf("\nInsufficient funds")
		return false
	}
	return true;
}

func withdrawMoney(balance float64, withdrawAmount float64) (float64, bool) {
	if (!isValidWithdrawAmount(balance, withdrawAmount)) {
		return balance, false
	}
	return balance - withdrawAmount, true
}
