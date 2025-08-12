package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func old_main() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue, realFutureValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	// futureValue, realFutureValue := calculateFutureValuesInSimpleWay(investmentAmount, expectedReturnRate, years)

	fmt.Println("Expected future value", futureValue)
	fmt.Println("REAL future value", realFutureValue)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

// SHORTER WAY TO CALCULATE FUTURE VALUES
// func calculateFutureValuesInSimpleWay(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
// 	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
// 	rfv = fv / math.Pow(1+inflationRate/100, years)
// 	return
// }
