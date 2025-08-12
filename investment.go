/*
package main

import (
	"fmt"
	"math"
)


-------- 1. BASIC LEVEL OF GO CODING --------
func main() {
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	// var futureValue = float64(investmentAmount) * math.Pow((1+expectedReturnRate/100), float64(years))
	var futureValue = investmentAmount * math.Pow((1+expectedReturnRate/100), years)

	fmt.Println(futureValue)
}


-------- 2. USING OF LITTLE SHORT-HAND SYNTAX WHILE CODING --------
func main() {
	var investmentAmount, years float64 = 1000, 10
	expectedReturnRate := 5.5

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	const inflationRate = 2.5
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Print(futureValue)
	fmt.Print(futureRealValue)
}

-------- 3. INTERACTIVE USER INPUT FOR INVESTMENT CALCULATION --------
func main() {
	var investmentAmount, years float64
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	const inflationRate = 2.5
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Expected future value", futureValue)
	fmt.Println("REAL future value", futureRealValue)
}
*/

package main
