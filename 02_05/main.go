package main

import "fmt"

func main() {
	//using a rune slice to modify a string
	s := "hello"
	fmt.Println(s)
	rs := []rune(s)
	rs[0] = 'H'
	rs[2] = 'L'
	rs[4] = 'O'
	fmt.Printf("Rune slice: %v, (type %T) \n", rs, rs)
	s = string(rs)
	fmt.Println(s)
}
