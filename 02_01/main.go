package main

import "fmt"

func main() {
	//strings are immutable
	s := "hello"
	// s[0] = 'H' // this would cause a compile-time error
	fmt.Println(s)

	// to modify a string, create a new one
	s = "H" + s[1:]
	fmt.Println(s)

	// another way to modify specific characters
	s = "H" + s[1:2] + "L" + s[3:] + "O"
	fmt.Println(s)
}
