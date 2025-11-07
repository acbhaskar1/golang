package main

import "fmt"

func main() {

	// conditional statement with logical operators

	x := 15
	// Compounded if statements Option A

	if x > 10 && x < 20 {
		fmt.Printf("x is between 10 and 20: %d\n", x)
	}

	// if-else with logical operators
	if x < 10 || x > 20 {
		fmt.Printf("x is either less than 10 or greater than 20: %d\n", x)
	} else {
		fmt.Printf("x is between 10 and 20: %d\n", x)
	}

	// nested if statements Option B

	if x > 10 {
		if x < 20 {
			fmt.Printf("x is between 10 and 20: %d\n", x)
		}
	}
}
