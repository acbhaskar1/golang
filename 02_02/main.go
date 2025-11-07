package main

import (
	"fmt"
	"strings"
)

func main() {
	//slices are mutable
	s := []string{"h", "e", "l", "l", "o"}
	fmt.Println(strings.Join(s, ""))
	s[0] = "H"
	fmt.Println(strings.Join(s, ""))
	s[2] = "L"
	s[4] = "O"
	fmt.Println(strings.Join(s, ""))
}
