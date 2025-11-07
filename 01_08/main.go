package main

import "fmt"

func main() {
	// switch with multiple expressions in cases
	score := 85
	switch {
	case score >= 90:
		fmt.Printf("Grade: A\n")
	case score >= 80:
		fmt.Printf("Grade: B\n")
	case score >= 70:
		fmt.Printf("Grade: C\n")
	default:
		fmt.Printf("Grade: F\n")
	}
	// switch with fallthrough
	day := "Friday"
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
	// switch with type assertion
	var value interface{} = "Hello, Go!"
	switch v := value.(type) {
	case int:
		fmt.Printf("value is an integer: %d\n", v)
	case string:
		fmt.Printf("value is a string: %s\n", v)
	default:
		fmt.Printf("value is of unknown type\n")
	}
}
