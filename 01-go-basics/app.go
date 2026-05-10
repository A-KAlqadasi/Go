package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.0	
	var investmentAmount, expectedReturnRate, years float64
	fmt.Print("Enter investment amount: ")
	fmt.Scan(&investmentAmount)
	fmt.Print("Enter expected return rate (in %): ")
	fmt.Scan(&expectedReturnRate)
	fmt.Print("Enter number of years: ")
	fmt.Scan(&years)
		
	futureValue := investmentAmount * math.Pow(1 + expectedReturnRate / 100, years)
	futureRealValue := futureValue / math.Pow(1 + inflationRate / 100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

}