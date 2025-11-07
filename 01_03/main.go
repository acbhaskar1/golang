package main

import (
	"fmt"
)

func main() {

	// conditonal statement
	x := 10

	if x > 5 {
		fmt.Printf("x is big: %d\n", x)
	}

	// if-else statement
	if x%2 == 0 {
		fmt.Printf("x is even: %d\n", x)
	} else {
		fmt.Printf("x is odd: %d\n", x)
	}

	// if-else-if ladder
	if x < 0 {
		fmt.Printf("x is negative: %d\n", x)
	} else if x == 0 {
		fmt.Printf("x is zero: %d\n", x)
	} else {
		fmt.Printf("x is positive: %d\n", x)
	}

}
