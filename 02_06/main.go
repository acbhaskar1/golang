package main

import (
	"fmt"
)

func main() {
	//length and capacity of slices
	s := []string{"h", "e", "l", "l", "o"}
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	s = append(s, "!")
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	s = append(s, "How", "are", "you", "doing", "today", "?")
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	s = append(s, "I'am fine.")
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
	s[12] = "I"
	s = append(s, "am", "fine", "thank", "you", ".", "Hope ", "You", "are also", "doing", "great", "!")
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
