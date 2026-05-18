package main

import (
	"fmt"

	"go-bank.com/fileops"

	"github.com/Pallinder/go-randomdata"
)

const fileName = "balance.txt"

func main() {
	fmt.Println("============================")
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("You can reach us 24/7 ", randomdata.PhoneNumber())
	fmt.Println("============================")
	

	bankStartScreen()

	fmt.Println("============================")
	fmt.Println("Thank you for using Go Bank! Goodbye!")
	fmt.Println("============================")
}



func bankStartScreen() {

	var option int
	for option != 4 {
		
		balance, err := fileops.ReadValueFromFile(fileName)
		if err != nil {
			fmt.Println("Error: ", err)
			panic("Can't continue, sorry")
		}

		showOptions()
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
				fileops.WriteValueToFile(newBalance, fileName)
			}
		case 3:
			var withdrawAmount float64
			withdrawAmount = readAmount("Enter the amount you want to withdraw: ")
			balance, isValidationPassed := withdrawMoney(balance, withdrawAmount)
			if(!isValidationPassed) {
				continue
			}
			fmt.Printf("\nYour new balance is: $%.2f\n", balance)
			fileops.WriteValueToFile(balance, fileName)
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
