package main

import "fmt"

func main() {

	// switch statement

	x := 25

	switch {
	case x < 10:
		fmt.Printf("x is less than 10: %d\n", x)
	case x >= 10 && x < 20:
		fmt.Printf("x is between 10 and 20: %d\n", x)
	case x >= 20 && x < 30:
		fmt.Printf("x is between 20 and 30: %d\n", x)
	default:
		fmt.Printf("x is 30 or more: %d\n", x)
	}
	// switch with fallthrough

	day := "Monday"

	switch day {
	case "Monday":
		fmt.Printf("%s is the first day of the work week\n", day)
	case "Friday":
		fmt.Printf("%s is the last day of the work week\n", day)
		fallthrough
	case "Saturday", "Sunday":
		fmt.Printf("%s is a weekend\n", day)
	default:
		fmt.Printf("%s is a weekday\n", day)
	}
}
