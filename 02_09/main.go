package main

import (
	"fmt"
)

func main() {
	// range over array
	numbers := [5]int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}

	// range over slice
	slice := []string{"apple", "banana", "cherry"}
	for index, value := range slice {
		fmt.Printf("Index: %d, Value: %s\n", index, value)
	}
	fmt.Println()
	// range over map
	capitals := map[string]string{
		"USA":    "Washington, D.C.",
		"France": "Paris",
		"Japan":  "Tokyo",
	}
	for country, capital := range capitals {
		fmt.Printf("Country: %s, Capital: %s\n", country, capital)
	}
	fmt.Println()

	// range over string
	str := "hello"
	for index, char := range str {
		fmt.Printf("Index: %d, Character: %c\n", index, char)
	}
	fmt.Println()
	// ignoring index or value
	values := []int{100, 200, 300}
	sum := 0
	for _, value := range values {
		sum += value
	}
	fmt.Printf("Sum of values: %d\n", sum)
}
