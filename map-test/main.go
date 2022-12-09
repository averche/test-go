package main

import (
	"fmt"
	"strings"
)

// Convert a number to a Roman numeral
func numberToRoman(n int) string {
	// Define the mapping of digits to Roman numerals
	m := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	// Initialize the result
	result := ""

	// Iterate over the digits in descending order
	for d := 1000; d > 0; d /= 10 {
		// Calculate the digit value
		digit := n / d

		// Add the corresponding Roman numeral to the result
		result += strings.Repeat(m[d], digit)

		// Remove the processed digit from the number
		n %= d
	}

	return result
}

// Define the mapping of digits to Roman numerals
var globalMap = map[int]string{
	1:    "I",
	4:    "IV",
	5:    "V",
	9:    "IX",
	10:   "X",
	40:   "XL",
	50:   "L",
	90:   "XC",
	100:  "C",
	400:  "CD",
	500:  "D",
	900:  "CM",
	1000: "M",
}

func numberToRomanGlobalMap(n int) string {
	// Initialize the result
	result := ""

	// Iterate over the digits in descending order
	for d := 1000; d > 0; d /= 10 {
		// Calculate the digit value
		digit := n / d

		// Add the corresponding Roman numeral to the result
		result += strings.Repeat(globalMap[d], digit)

		// Remove the processed digit from the number
		n %= d
	}

	return result
}

func main() {
	r := numberToRoman(1066)

	fmt.Println(r)
}
