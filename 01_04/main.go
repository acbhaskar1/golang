package main

import "fmt"

func main() {
	// switch statement
	x := 10
	switch x {
	case 5:
		fmt.Printf("x is five: %d\n", x)
	case 10:
		fmt.Printf("x is ten: %d\n", x)
	case 15:
		fmt.Printf("x is fifteen: %d\n", x)
	default:
		fmt.Printf("x is something else: %d\n", x)
	}

	//naked switch
	num := 7
	switch {
	case num < 0:
		fmt.Printf("num is negative: %d\n", num)
	case num == 0:
		fmt.Printf("num is zero: %d\n", num)
	case num > 0:
		fmt.Printf("num is positive: %d\n", num)
	}
	// switch with multiple cases
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Printf("%s is a weekend\n", day)
	default:
		fmt.Printf("%s is a weekday\n", day)
	}
}
