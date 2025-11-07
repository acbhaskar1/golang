package main

import "fmt"

func main() {
	// if with a short statement
	if num := 9; num%2 == 0 {
		fmt.Printf("%d is even\n", num)
	} else {
		fmt.Printf("%d is odd\n", num)
	}

}
