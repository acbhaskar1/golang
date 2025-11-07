package main

import "fmt"

func main() {
	//count of even ended numbers of a multiplying two four digit numbers
	count := 0
	for i := 1000; i <= 9999; i++ {
		for j := 1000; j <= 9999; j++ {
			product := i * j
			if isEvenEnded(product) {
				count++
			}
		}
	}
	fmt.Println("Count of even ended numbers of multiplying two four digit numbers:", count)
}
func isEvenEnded(num int) bool {
	lastDigit := num % 10

	// Handle single-digit numbers
	if num < 10 {
		return true
	}

	// Get first digit efficiently
	firstDigit := num
	for firstDigit >= 10 {
		firstDigit /= 10
	}

	return firstDigit == lastDigit
}
