package main

import "fmt"

func main() {

	// switch with multiple cases and no expression
	day := "Saturday"
	switch day {
	case "Saturday", "Sunday":
		fmt.Printf("%s is a weekend\n", day)
	default:
		fmt.Printf("%s is a weekday\n", day)
	}
	// switch with initialization statement
	switch x := 15; {
	case x < 10:
		fmt.Printf("x is less than 10: %d\n", x)
	case x >= 10 && x < 20:
		fmt.Printf("x is between 10 and 20: %d\n", x)
	default:
		fmt.Printf("x is 20 or more: %d\n", x)
	}
	// switch with type assertion
	var value interface{} = 42
	switch v := value.(type) {
	case int:
		fmt.Printf("value is an integer: %d\n", v)
	case string:
		fmt.Printf("value is a string: %s\n", v)
	default:
		fmt.Printf("value is of unknown type\n")
	}
}
