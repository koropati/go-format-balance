package main

import (
	"fmt"
	"strconv"
	"strings"
)

func hasZeroAfterComma(number float64) bool {
	// Convert the float number to a string.
	strNumber := strconv.FormatFloat(number, 'g', -1, 64)
	// Check if the string contains a decimal point followed by zero.
	if strings.Contains(strNumber, ".0") {
		return true
	}
	return false
}

func getFirstNumberBeforeComma(number float64) float64 {
	// Convert the float number to a string.
	strNumber := strconv.FormatFloat(number, 'g', -1, 64)

	// Get the index of the decimal point.
	decimalPointIndex := strings.Index(strNumber, ".")

	// If there is no decimal point, then return the number itself.
	if decimalPointIndex == -1 {
		return number
	}

	// Return the substring of the string up to the decimal point.
	result, _ := strconv.ParseFloat(strNumber[:decimalPointIndex], 64)
	return result
}

func formatBalance(balance float64) (formatedBalance float64) {
	if hasZeroAfterComma(balance) {
		return getFirstNumberBeforeComma(balance)
	} else {
		return balance
	}
}

func main() {
	fmt.Println(formatBalance(7.01))
	fmt.Println(formatBalance(7.101))
}
