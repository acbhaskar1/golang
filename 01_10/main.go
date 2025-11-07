package main

import "fmt"

func main() {
	//for with a short statement
	for i := 1; i <= 5; i++ {
		fmt.Printf("Iteration %d\n", i)
	}

	// while loop using for
	count := 1
	for count <= 5 {
		fmt.Printf("Count %d\n", count)
		count++
	}
	// infinite loop with break
	sum := 0
	for {
		sum += 2
		if sum >= 10 {
			break
		}
		fmt.Printf("Sum is %d\n", sum)
	}

	// range loop over a slice
	numbers := []int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index %d: Value %d\n", index, value)
	}
}
